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
				downUrl := utils.FormatUrl(video.DownUrl)
				coverUrl := utils.FormatUrl(video.Cover)

				message := model.VideoReleaseMessage{
					PublishedSiteName: siteVideoUrls.SiteName,
					VideoId:           int(video.ID),
					Title:             video.Title,
					CreatedAt:         video.CreatedAt,
				}

				if siteVideoUrls.CDNPlayUrl != "" {
					message.CDNPlayUrl = siteVideoUrls.CDNPlayUrl + playUrl
				}
				if siteVideoUrls.CFPlayUrl != "" {
					message.CFPlayUrl = siteVideoUrls.CFPlayUrl + playUrl
				}
				if siteVideoUrls.DirectPlayUrl != "" {
					message.DirectPlayUrl = siteVideoUrls.DirectPlayUrl + playUrl
				}
				if downUrl != "" {
					message.DownUrl = siteVideoUrls.DownloadUrl + downUrl
				}
				if coverUrl != "" {
					message.CoverUrl = siteVideoUrls.VideoCover + coverUrl
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
