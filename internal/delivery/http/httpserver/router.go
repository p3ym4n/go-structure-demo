package httpserver

import (
	"github.com/go-chi/chi/v5"
	"github.com/goflink/rider-workforce-common/log"
	"go-structure-demo/internal/config"
	"go-structure-demo/internal/delivery/http/middleware"
	"net/http"
)

func newRouter(cfg *config.Config, logger log.Logger) *chi.Mux {
	router := chi.NewRouter()

	// add the essential middlewares Like
	router.With(middleware.Recoverer(logger))
	// 	requests tracer
	// 	requests logger
	// 	timeout
	// 	CORS

	router.NotFound(func(writer http.ResponseWriter, request *http.Request) {
		// fancy response
	})

	router.MethodNotAllowed(func(writer http.ResponseWriter, request *http.Request) {
		// fancy response
	})

	return router
}
