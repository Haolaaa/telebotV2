package services

import (
	"context"
	"encoding/json"
	"telebotV2/global"
	"telebotV2/model"
	"telebotV2/utils"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

func AllVideosTask() {
	videos, err := GetVideos()
	if err != nil {
		global.LOG.Error("get videos error: ", zap.Error(err))
		return
	}

	releasedVideos, err := filterReleasedVideo(videos)
	if err != nil {
		global.LOG.Error("filter released videos error: ", zap.Error(err))
		return
	}

	totalLength := len(releasedVideos)

	for _, releasedVideo := range releasedVideos {
		releasedVideo.Total = totalLength

		messageBytes, err := json.Marshal(releasedVideo)
		if err != nil {
			global.LOG.Error("marshal released video error: ", zap.Error(err))
			return
		}

		err = global.Writer.WriteMessages(
			context.Background(),
			kafka.Message{
				Topic: "video_read_all",
				Key:   []byte("video_read_all"),
				Value: messageBytes,
			},
		)
	}
}

func filterReleasedVideo(videos []model.Video) (filteredVideos []model.VideoReleaseMessage, err error) {
	var releasedVideos []model.VideoRelease

	err = global.DB.Table("video_release").Find(&releasedVideos).Error
	if err != nil {
		global.LOG.Error("get released videos error: ", zap.Error(err))
		return nil, err
	}

	for _, video := range videos {
		for _, releasedVideo := range releasedVideos {
			if releasedVideo.Status == 1 && video.ID == uint(releasedVideo.VideoId) {
				publishedSite, err := GetSiteVideoUrls(int(releasedVideo.SiteID))
				if err != nil {
					global.LOG.Error("get published site error: ", zap.Error(err))
					return nil, err
				}

				playUrl := utils.FormatM3u8Suffix(video.PlayUrl, publishedSite.SiteKey)

				filteredVideos = append(filteredVideos, model.VideoReleaseMessage{
					PublishedSiteName: publishedSite.SiteName,
					VideoId:           int(video.ID),
					Title:             video.Title,
					PlayUrl:           playUrl,
					DirectPlayUrl:     publishedSite.DirectPlayUrl + playUrl,
					CFPlayUrl:         publishedSite.CFPlayUrl + playUrl,
					CDNPlayUrl:        publishedSite.CDNPlayUrl + playUrl,
					CreatedAt:         video.CreatedAt,
				})
			}
		}
	}

	return
}
