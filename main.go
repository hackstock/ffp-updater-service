package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hackstock/ffp-updater-service/pkg"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

var env = struct {
	Port          int           `envconfig:"PORT" required:"true"`
	Environment   string        `envconfig:"ENVIRONMENT" default:"development"`
	APIHost       string        `envconfig:"API_HOST"`
	SyncFrequency time.Duration `envconfig:"SYNC_FREQUENCY" default:"30m"`
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
	defer logger.Sync()

	logger.Info("env vars and logger initialized successfully")

	client := &http.Client{Timeout: 30 * time.Second}
	process := pkg.NewRewardsProcess(env.APIHost, client, logger)

	ticker := time.NewTicker(env.SyncFrequency)
	go func() {
		for now := range ticker.C {
			logger.Info("starting to process unprocessed flight records")
			err := process.Run()
			if err != nil {
				logger.Warn("failed processing unprocessed flight records",
					zap.Error(err))
			}
			elapsed := time.Since(now)
			logger.Info("finished processing unprocessed rows in",
				zap.Duration("duration", elapsed))
		}
	}()

	listener, err := net.Listen("tcp4", fmt.Sprintf(":%d", env.Port))
	if err != nil {
		logger.Fatal("failed binding to port",
			zap.Int("port", env.Port),
			zap.Error(err),
		)
	}

	defer listener.Close()

	url := fmt.Sprintf("http://%s/", listener.Addr())
	logger.Info("listening on:",
		zap.String("url", url),
	)

	mux := http.NewServeMux()
	mux.HandleFunc("/status", pkg.StatusHandler())

	srv := http.Server{
		Handler: mux,
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	idleConnsClosed := make(chan struct{})
	go func() {
		defer close(idleConnsClosed)
		recv := <-sigs
		logger.Info("received signal, shutting down", zap.String("signal", recv.String()))

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			logger.Warn("error shutting down server", zap.Error(err))
		}
	}()

	if err := srv.Serve(listener); err != nil {
		if err != http.ErrServerClosed {
			logger.Fatal("http.Serve returned an error",
				zap.Error(err),
			)
		}
	}

	<-idleConnsClosed
}
