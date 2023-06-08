package initialize

import (
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	return GormMysql()
}

// func RegisterTables() {
// 	db := global.DB
// 	site1 := &model.SiteVideoUrls{
// 		ID:            1,
// 		SiteName:      "桃红",
// 		SiteKey:       "taohsj",
// 		DirectPlayUrl: "https://thplay.taoh2550.com",
// 		CFPlayUrl:     "https://thplay.3cej.com",
// 		CDNPlayUrl:    "https://aws-news-taohplay.3cej.com",
// 		VideoCover:    "https://thimg.3cej.com",
// 		SiteId:        4,
// 	}
// 	site2 := &model.SiteVideoUrls{
// 		ID:            2,
// 		SiteName:      "汤姆叔叔",
// 		SiteKey:       "uncletom",
// 		DirectPlayUrl: "https://tmplay3.tomtv079.com",
// 		CFPlayUrl:     "https://tmplay3.3cej.com",
// 		CDNPlayUrl:    "https://aws-tmplay.3cej.com",
// 		VideoCover:    "https://tmpic3.3cej.com",
// 		SiteId:        1,
// 	}
// 	site3 := &model.SiteVideoUrls{
// 		ID:            3,
// 		SiteName:      "骚虎",
// 		SiteKey:       "saohuold",
// 		DirectPlayUrl: "https://shpplay.2850saohu.com",
// 		CFPlayUrl:     "https://shpplay.3cej.com",
// 		CDNPlayUrl:    "https://aws-news-shpplay.3cej.cc",
// 		VideoCover:    "https://shpimg.3cej.com",
// 		SiteId:        3,
// 	}
// 	site4 := &model.SiteVideoUrls{
// 		ID:            4,
// 		SiteName:      "四虎",
// 		SiteKey:       "dh365",
// 		DirectPlayUrl: "https://4hu.saohu687.com",
// 		CFPlayUrl:     "https://sihplay.3cej.com",
// 		CDNPlayUrl:    "https://aws-news-sihplay.3cej.cc",
// 		VideoCover:    "https://shpimg.3cej.com",
// 		SiteId:        17,
// 	}
// 	err := db.AutoMigrate(
// 		&model.SiteVideoUrls{},
// 	)
// 	if err != nil {
// 		global.LOG.Error("register tables failed", zap.Error(err))
// 		os.Exit(0)
// 	}
// 	var count int64
// 	db.Model(&model.SiteVideoUrls{}).Count(&count)
// 	if count == 0 {
// 		db.CreateInBatches([]model.SiteVideoUrls{
// 			*site1,
// 			*site2,
// 			*site3,
// 			*site4,
// 		}, 6)
// 	}
// 	global.LOG.Info("register tables success")
// }
