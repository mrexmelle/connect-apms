package template

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mrexmelle/connect-apms/internal/config"
)

type Controller struct {
	Config          *config.Config
	TemplateService *Service
}

func NewController(cfg *config.Config, svc *Service) *Controller {
	return &Controller{
		Config:          cfg,
		TemplateService: svc,
	}
}

func (c *Controller) GetByCode(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")

	response := c.TemplateService.RetrieveByCode(code)
	if response.Status != "OK" {
		http.Error(w, "GET failure: "+response.Status, http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}

func (c *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	response := c.TemplateService.RetrieveAll()
	if response.Status != "OK" {
		http.Error(w, "GET failure: "+response.Status, http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}
