package services

import (
	"telebotV2/global"
	"telebotV2/model"
)

func GetSiteVideoUrls(sideId int) (sitePlayUrls model.SiteVideoUrls, err error) {
	err = global.DB.Table("site_video_urls").Where("site_id = ?", sideId).First(&sitePlayUrls).Error
	return
}
