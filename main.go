package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

var env = struct {
	Port          int    `envconfig:"PORT" required:"true"`
	Environment   string `envconfig:"ENV" default:"development"`
	SyncFrequency int    `envconfig:"SYNC_FREQUENCY" default:"1"`
	DatabaseURI   string `envconfig:"DATABASE_URI" required:"true"`
}{}

func init() {
	err := envconfig.Process("FFPUPDATER", &env)
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
