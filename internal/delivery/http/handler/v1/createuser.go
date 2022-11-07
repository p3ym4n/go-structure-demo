package v1

import (
	"go-structure-demo/internal/controller"
	"go-structure-demo/internal/param"
	"go-structure-demo/internal/repository/postgresrepo"
	"go-structure-demo/internal/validator"
	"net/http"
)

func CreateUser(postgresRepo *postgresrepo.PostgresRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestDTO := new(param.CreateUserRequest)

		err := requestDTO.BindFromChi(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = validator.CreateUserRequest(r.Context(), requestDTO, postgresRepo)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		responseDTO := controller.NewUserController(postgresRepo).CreateUser(r.Context(), requestDTO)
		if responseDTO.Error != nil {
			w.WriteHeader(responseDTO.StatusCode)
			_, _ = w.Write([]byte("the error is" + responseDTO.Error.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(responseDTO.ToJson())
	}
}
