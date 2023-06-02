package core

import (
	"telebotV2/canal"
	"telebotV2/global"
	"telebotV2/kafka"

	"go.uber.org/zap"
)

func RunServer() {
	kafka := kafka.NewKafkaService()
	global.Writer = kafka.Writer

	err := canal.RunCanal(true)
	if err != nil {
		global.LOG.Error("run canal error: ", zap.Error(err))
	}

	global.LOG.Info("server run success")
}
