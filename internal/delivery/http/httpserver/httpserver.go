package httpserver

import (
	"context"
	"fmt"
	"github.com/goflink/rider-workforce-common/log"
	"github.com/goflink/rider-workforce-common/pubsub"
	"go-structure-demo/internal/config"
	"go-structure-demo/internal/delivery/http/handler/private"
	v1 "go-structure-demo/internal/delivery/http/handler/v1"
	"go-structure-demo/internal/repository/postgresrepo"
	"go-structure-demo/internal/repository/redisrepo"
	"net/http"
	"os"
)

func New(
	cfg *config.Config,
	logger log.Logger,
	redisRepo *redisrepo.RedisRepo,
	postgresRepo *postgresrepo.PostgresRepo,
	pubsubClientA *pubsub.GCPClient,
	pubsubClientB *pubsub.GCPClient,
) *Server {
	router := newRouter(cfg, logger)

	router.Get("/health", private.Health(redisRepo, postgresRepo, pubsubClientA, pubsubClientB))
	router.Post("/v1/user", v1.CreateUser(postgresRepo))

	return &Server{
		logger: logger,
		srv: http.Server{
			Handler:      router,
			ReadTimeout:  cfg.HTTP.ReadTimeout,
			WriteTimeout: cfg.HTTP.WriteTimeout,
			IdleTimeout:  cfg.HTTP.IdleTimeout,
			Addr:         fmt.Sprintf("%s:%v", os.Getenv("HOST_IP"), cfg.HTTP.Port),
		},
	}
}

type Server struct {
	logger log.Logger
	srv    http.Server
}

func (s *Server) Start() {
	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.logger.Error("server start", err)
	}
}

func (s *Server) Shutdown(ctx context.Context) {
	err := s.srv.Shutdown(ctx)
	if err != nil {
		s.logger.Error("server shutdown", err)
	}
}
