package route

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/reangeline/micro_saas/internal/presentation/controller"
)

func InitializeUserRoutes(controller *controller.UserController, r chi.Router) {

	r.Route("/users", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Create User"))

		})

		r.Post("/", controller.CreateUserRest)
	})

}
