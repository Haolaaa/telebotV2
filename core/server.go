package core

import (
	"telebotV2/canal"
	"telebotV2/global"
	"telebotV2/kafka"
	"telebotV2/services"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func RunServer() {
	kafka := kafka.NewKafkaService()
	global.Writer = kafka.Writer

	err := canal.RunCanal(true)
	if err != nil {
		global.LOG.Error("run canal error: ", zap.Error(err))
	}

	cron := cron.New(cron.WithSeconds())
	_, err = cron.AddFunc("0 0 */4 * * *", services.AllVideosTask)
	if err != nil {
		global.LOG.Error("add cron task error: ", zap.Error(err))
	}
	zap.L().Log(zap.InfoLevel, "add cron task success")
	cron.Start()

	global.LOG.Info("server run success")
}
