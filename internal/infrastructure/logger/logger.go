package logger

import (
	"github.com/Kingpant/golang-clean-architecture-template/internal/infrastructure/config"
	"go.uber.org/zap"
)

func InitLogger(appEnv config.AppEnvType) *zap.SugaredLogger {
	switch appEnv {
	case config.AppEnvLocal:
		logger, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}

		return logger.Sugar()
	default:
		logger, err := zap.NewProduction()
		if err != nil {
			panic(err)
		}
		return logger.Sugar()
	}
}
