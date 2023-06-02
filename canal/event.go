package canal

import (
	"context"
	"encoding/json"
	"telebotV2/global"
	"telebotV2/model"
	"telebotV2/services"
	"telebotV2/utils"

	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

const (
	StatusPublished = 1
)

func GetParseValue(event *canal.RowsEvent) (*canal.RowsEvent, map[string]interface{}) {
	var rows = map[string]interface{}{}

	for colIndex, currCol := range event.Table.Columns {
		colValue := event.Rows[len(event.Rows)-1][colIndex]

		rows[currCol.Name] = colValue
	}

	return event, rows
}

func TableEventDispatcher(event *canal.RowsEvent, row map[string]interface{}) {
	if global.Writer == nil {
		global.LOG.Error("kafka writer is nil")
		return
	}

	modelStruct := model.GetModelStruct(event.Table.Name)
	rowModel, ok := utils.MapStructureRow(modelStruct, row)
	if !ok {
		return
	}
	switch event.Action {
	case canal.InsertAction, canal.UpdateAction:
		if event.Table.Name == "video_release" {
			rowModel := rowModel.(model.VideoRelease)
			if rowModel.Status == StatusPublished {
				video, err := services.GetVideo(rowModel.VideoId)
				if err != nil {
					global.LOG.Error("get video error: ", zap.Error(err))
					return
				}

				siteVideoUrls, err := services.GetSiteVideoUrls(int(rowModel.SiteID))
				if err != nil {
					global.LOG.Error("get site video urls error: ", zap.Error(err))
					return
				}

				playUrl := utils.FormatM3u8Suffix(video.PlayUrl, siteVideoUrls.SiteKey)

				message := model.VideoReleaseMessage{
					PublishedSiteName: siteVideoUrls.SiteName,
					VideoId:           int(video.ID),
					Title:             video.Title,
					PlayUrl:           playUrl,
					DirectPlayUrl:     siteVideoUrls.DirectPlayUrl + playUrl,
					CFPlayUrl:         siteVideoUrls.CFPlayUrl + playUrl,
					CDNPlayUrl:        siteVideoUrls.CDNPlayUrl + playUrl,
					CreatedAt:         video.CreatedAt,
				}

				messageBytes, err := json.Marshal(message)
				if err != nil {
					global.LOG.Error("marshal message error: ", zap.Error(err))
					return
				}

				err = global.Writer.WriteMessages(
					context.Background(),
					kafka.Message{
						Topic: "video_release",
						Key:   []byte("video_release"),
						Value: messageBytes,
					},
				)
				if err != nil {
					global.LOG.Error("write message error: ", zap.Error(err))
					return
				}
			}
		}
	}
}
