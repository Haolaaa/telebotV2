package canal

import (
	"telebotV2/global"
	"telebotV2/utils"

	"github.com/go-mysql-org/go-mysql/canal"
	"go.uber.org/zap"
)

type EventHandler struct {
	canal.DummyEventHandler
}

func (h *EventHandler) OnRow(event *canal.RowsEvent) error {
	_, rows := GetParseValue(event)

	if _, ok := utils.InArray(event.Table.Name, []string{"video_release"}); ok {
		TableEventDispatcher(event, rows)
	}

	return nil
}

func (*EventHandler) String() string {
	return "EventHandler"
}

func getCanalConfig() *canal.Config {
	cfg := canal.NewDefaultConfig()

	cfg.Addr = global.CONFIG.Canal.Addr
	cfg.User = global.CONFIG.Canal.User
	cfg.Password = global.CONFIG.Canal.Password
	cfg.Flavor = global.CONFIG.Canal.Flavor
	cfg.Charset = global.CONFIG.Canal.Charset
	cfg.ServerID = global.CONFIG.Canal.ServerID
	cfg.ParseTime = true

	cfg.Dump.TableDB = global.CONFIG.Canal.Dump.TableDB
	cfg.Dump.DiscardErr = global.CONFIG.Canal.Dump.DiscardErr
	cfg.Dump.SkipMasterData = global.CONFIG.Canal.Dump.SkipMasterData
	cfg.Dump.Tables = global.CONFIG.Canal.Dump.Tables

	return cfg
}

func RunCanal(isPos bool) error {
	cfg := getCanalConfig()
	c, err := canal.NewCanal(cfg)
	if err != nil {
		global.LOG.Error("create canal error: ", zap.Error(err))
	}

	c.SetEventHandler(&EventHandler{})

	if !isPos {
		return c.Run()
	}

	masterPos, err := c.GetMasterPos()
	if err != nil {
		global.LOG.Error("get master pos error: ", zap.Error(err))
	}

	global.LOG.Info("master pos: ", zap.Any("masterPos", masterPos))

	return c.RunFrom(masterPos)
}
