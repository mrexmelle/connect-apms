package template

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
		CollectionName: "templates",
	}
}

func (r *Repository) CreateWithDb(req Entity) (Entity, error) {
	_, err := r.Config.Db.Collection(r.CollectionName).InsertOne(
		context.Background(),
		req,
	)

	return req, err
}

func (r *Repository) FindByCode(code string) (Entity, error) {
	result := r.Config.Db.Collection(r.CollectionName).FindOne(
		context.Background(),
		bson.M{"code": code},
	)

	var response = Entity{}
	err := result.Decode(&response)
	return response, err
}

func (r *Repository) FindAll() ([]Entity, error) {
	result, err := r.Config.Db.Collection(r.CollectionName).Find(
		context.Background(),
		bson.M{},
	)
	if err != nil {
		return []Entity{}, err
	}

	var response []Entity
	if err = result.All(context.Background(), &response); err != nil {
		return []Entity{}, err
	}

	return response, nil
}
