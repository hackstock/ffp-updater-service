package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

var env = struct {
	Port          int    `envconfig:"FFPUPDATER_PORT" required:"true"`
	Environment   string `envconfig:"FFPUPDATER_ENV" default:"development"`
	SyncFrequency int    `envconfig:"FFPUPDATER_SYNC_FREQUENCY" default:"1"`
	DatabaseURI   string `envconfig:"FFPUPDATER_DATABASE_URI" required:"true"`
	SMSApiConfig  struct {
		Username string `envconfig:"USERNAME" required:"true"`
		Password string `envconfig:"PASSWORD" required:"true"`
		SenderID string `envconfig:"SENDERID" required:"true"`
	} `envconfig:"SMS_API"`
}{}

func init() {
	err := envconfig.Process("", &env)
	if err != nil {
		panic(fmt.Errorf("failed loading env vars : %v", err))
	}
}

func initLogger(target string) (*zap.Logger, error) {
	if target == "production" {
		return zap.NewProduction()
	}

	return zap.NewDevelopment()
}

func main() {
	logger, err := initLogger(env.Environment)
	if err != nil {
		panic(fmt.Errorf("failed initializing logger : %v", err))
	}

	logger.Info("env vars and logger initialized successfully")
}
