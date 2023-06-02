package utils

import (
	"telebotV2/global"

	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

func InArray(val string, slice []string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}

	return -1, false
}

func MapStructureRow(model interface{}, row map[string]interface{}) (interface{}, bool) {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &model,
	})
	if err != nil {
		global.LOG.Error("mapstructure new decoder error", zap.Error(err))
		return nil, false
	}

	err = decoder.Decode(row)
	if err != nil {
		global.LOG.Error("mapstructure decode error", zap.Error(err))
		return nil, false
	}

	return model, true
}
