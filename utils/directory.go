package utils

import (
	"errors"
	"os"
	"telebotV2/global"

	"go.uber.org/zap"
)

func PathExits(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("file exists")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExits(v)
		if err != nil {
			return err
		}

		if !exist {
			global.LOG.Error("create directory "+v, zap.Any(" error", err))
		}
	}
	return err
}
