package main

import (
	"context"
	"go-structure-demo/internal/config"
	"go-structure-demo/internal/delivery/http/httpserver"
	"go-structure-demo/internal/delivery/pubsub/subscriber"
	"go-structure-demo/internal/log"
	"go-structure-demo/internal/pubsub"
	"go-structure-demo/internal/repository/postgresrepo"
	"go-structure-demo/internal/repository/redisrepo"
	"os"
	"os/signal"
)

func main() {
	ctx := context.Background()
	cfg := config.Read()

	logger, loggerCloser := log.NewZapFromEnv(cfg.AppName)
	defer loggerCloser()

	redisRepo, redisRepoClose := redisrepo.New(cfg)
	defer redisRepoClose()

	postgresRepo, postgresRepoCloser := postgresrepo.New(cfg)
	defer postgresRepoCloser()

	pubsubClientA, err := pubsub.New(logger, ctx, cfg.PubSub.ProjectA)
	if err != nil {
		logger.Error("initializing pubsub", err)
	}
	pubsubClientB, err := pubsub.New(logger, ctx, cfg.PubSub.ProjectB)
	if err != nil {
		logger.Error("initializing pubsub", err)
	}

	httpServer := httpserver.New(cfg, logger, redisRepo, postgresRepo, pubsubClientA, pubsubClientB)
	go httpServer.Start()

	pubsubClientCloser := subscriber.Subscribe(ctx, cfg, logger, redisRepo, postgresRepo, pubsubClientA, pubsubClientB)
	defer pubsubClientCloser()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	logger.Info("shutting signal received")
	c, cancel := context.WithTimeout(context.Background(), cfg.HTTP.GracefulShutdown)
	defer cancel()
	httpServer.Shutdown(c)
	<-c.Done()
	logger.Info("bye bye")
}
