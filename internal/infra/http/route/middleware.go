package route

import (
	"fmt"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	httpSwagger "github.com/swaggo/http-swagger"
)

func InitializeMiddlewares(r chi.Router) {

	r.Use(middleware.Heartbeat("/health"))

	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	swaggerUrl := fmt.Sprintf("http://0.0.0.0:%s/docs/doc.json", "8080")
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(swaggerUrl)))

}
