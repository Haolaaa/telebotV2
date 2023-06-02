package main

import (
	"telebotV2/core"
	"telebotV2/global"
	"telebotV2/initialize"

	"go.uber.org/zap"
)

func main() {
	global.VP = core.Viper()
	global.LOG = core.Zap()
	zap.ReplaceGlobals(global.LOG)

	global.DB = initialize.Gorm()
	if global.DB != nil {
		db, _ := global.DB.DB()
		initialize.RegisterTables()
		defer db.Close()
	}

	core.RunServer()

	select {}
}
