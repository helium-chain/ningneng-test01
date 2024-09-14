package zap

import (
	"example.com/ningneng/pkg/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger(isSys bool) (logger *zap.Logger, err error) {
	var (
		mode     = global.Viper.GetString("application.mode")
		rootPath = global.RootPath
	)

	if mode == "test" || mode == "release" {
		fileName := rootPath + global.Viper.GetString("logger.path")
		if isSys {
			fileName = fileName + "/runtime.log"
		} else {
			fileName = fileName + "/logger.log"
		}
		lumberJackLogger := &lumberjack.Logger{
			Filename:   fileName,
			MaxSize:    global.Viper.GetInt("logger.max_size"),
			MaxAge:     global.Viper.GetInt("logger.max_age"),
			MaxBackups: global.Viper.GetInt("logger.max_backups"),
			LocalTime:  true,
			Compress:   global.Viper.GetBool("logger.compress"),
		}

		writer := zapcore.AddSync(lumberJackLogger)
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

		zapCore := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), writer, zap.InfoLevel)
		logger = zap.New(zapCore, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
	} else {
		logger, err = zap.NewDevelopment()
	}

	return
}
