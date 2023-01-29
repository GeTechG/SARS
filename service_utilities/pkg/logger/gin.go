package logger

import (
	"errors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"time"
)

func ConnectToGin(router *gin.Engine) error {
	logger := GetLogger()
	if logger == nil {
		return errors.New("logger not initialize")
	}
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, false))
	return nil
}
