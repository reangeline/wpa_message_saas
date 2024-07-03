package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/reangeline/micro_saas/internal/presentation/controller"
)

func InitializeWhatsAppRoutes(controller *controller.WhatsAppController, r chi.Router) {

	r.Route("/webhook", func(r chi.Router) {
		r.Post("/", controller.CreateMessageWhatsApp)
		r.Get("/", controller.VerifyTokenWhatsApp)
	})

}
