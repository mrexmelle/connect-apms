package opts

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/mrexmelle/connect-apms/internal/config"
	"github.com/mrexmelle/connect-apms/internal/event"
	"github.com/mrexmelle/connect-apms/internal/proposal"
	"github.com/mrexmelle/connect-apms/internal/reviewer"
	"github.com/mrexmelle/connect-apms/internal/template"
	"go.uber.org/dig"
)

func NewConfig() *config.Config {
	cfg, err := config.New(
		"application", "yaml",
		[]string{
			"/etc/conf",
			"./config",
		},
	)
	if err != nil {
		panic(err)
	}
	return &cfg
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func Serve(cmd *cobra.Command, args []string) {
	container := dig.New()
	container.Provide(NewConfig)

	container.Provide(event.NewRepository)
	container.Provide(template.NewRepository)
	container.Provide(proposal.NewRepository)

	container.Provide(template.NewService)
	container.Provide(proposal.NewService)
	container.Provide(reviewer.NewService)

	container.Provide(template.NewController)
	container.Provide(proposal.NewController)

	process := func(
		proposalController *proposal.Controller,
		templateController *template.Controller,
		config *config.Config,
	) {
		r := chi.NewRouter()

		r.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://localhost:3000"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}))

		r.Route("/templates", func(r chi.Router) {
			r.Get("/", templateController.GetAll)
			r.Get("/{code}", templateController.GetByCode)
		})

		r.Route("/proposals", func(r chi.Router) {
			r.Post("/", proposalController.Create)
			r.Get("/{id}", proposalController.GetById)
		})

		err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r)

		if err != nil {
			panic(err)
		}
	}

	if err := container.Invoke(process); err != nil {
		panic(err)
	}
}

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start Connect APMS server",
	Run:   Serve,
}
