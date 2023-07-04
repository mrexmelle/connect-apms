package proposal

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mrexmelle/connect-apms/internal/config"
)

type Controller struct {
	Config          *config.Config
	ProposalService *Service
}

func NewController(
	cfg *config.Config,
	ps *Service,
) *Controller {
	return &Controller{
		Config:          cfg,
		ProposalService: ps,
	}
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var requestBody Entity
	json.NewDecoder(r.Body).Decode(&requestBody)

	response := c.ProposalService.Create(requestBody)
	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}

func (c *Controller) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	response := c.ProposalService.RetrieveById(id)
	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}
