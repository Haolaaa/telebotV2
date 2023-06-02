package services

import (
	"telebotV2/global"
	"telebotV2/model"
	"time"

	"go.uber.org/zap"
)

func GetVideo(videoId int) (video model.Video, err error) {
	err = global.DB.Table("video").Where("id = ?", videoId).First(&video).Error
	return
}

func GetVideos() (videos []model.Video, err error) {
	now := time.Now()
	twoDaysAgo := now.Add(-24 * time.Hour * 2)

	err = global.DB.Table("video").Where("created_at BETWEEN ? AND ?", twoDaysAgo, now).Find(&videos).Error
	if err != nil {
		global.LOG.Error("get videos error: ", zap.Error(err))
	}

	return
}
