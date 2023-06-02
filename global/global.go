package global

import (
	"telebotV2/config"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	VP     *viper.Viper
	CONFIG config.Config
	LOG    *zap.Logger
	Writer *kafka.Writer
)
