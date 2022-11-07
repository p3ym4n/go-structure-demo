package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/goflink/rider-workforce-common/log"
)

func Recoverer(logger log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					// http.ErrAbortHandler should not be logged according to the documentation
					if err != http.ErrAbortHandler {
						logger.ErrorWithContext(r.Context(), "panic recovered", map[string]interface{}{
							log.KeyError: err,
							"stacktrace": string(debug.Stack()),
						})
						w.WriteHeader(http.StatusInternalServerError)
					} else {
						w.WriteHeader(http.StatusGone)
					}

				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
