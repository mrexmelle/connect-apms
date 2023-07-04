package event

import (
	"context"

	"github.com/mrexmelle/connect-apms/internal/config"
	"go.mongodb.org/mongo-driver/bson"
)

type Repository struct {
	Config         *config.Config
	CollectionName string
}

func NewRepository(cfg *config.Config) *Repository {
	return &Repository{
		Config:         cfg,
		CollectionName: "events",
	}
}

func (r *Repository) Create(req Entity) (Entity, error) {
	_, err := r.Config.Db.Collection(r.CollectionName).InsertOne(
		context.Background(),
		req,
	)

	return req, err
}

func (r *Repository) FindByProposalId(proposalId string) ([]Entity, error) {
	result, err := r.Config.Db.Collection(r.CollectionName).Find(
		context.Background(),
		bson.M{"proposal_id": proposalId},
	)

	if err != nil {
		return []Entity{}, err
	}

	var response = []Entity{}
	err = result.Decode(&response)
	return response, err
}
