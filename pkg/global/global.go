package global

import (
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	// RootPath 项目根目录
	RootPath string

	Viper *viper.Viper

	Logger *zap.Logger

	Location *time.Location
)

func init() {
	Viper = viper.New()
}
