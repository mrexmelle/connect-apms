package proposal

import (
	"context"
	"fmt"

	"github.com/mrexmelle/connect-apms/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository struct {
	Config         *config.Config
	CollectionName string
}

func NewRepository(cfg *config.Config) *Repository {
	return &Repository{
		Config:         cfg,
		CollectionName: "proposals",
	}
}

func (r *Repository) Create(req Entity) (Entity, error) {
	result, err := r.Config.Db.Collection(r.CollectionName).InsertOne(
		context.Background(),
		req,
	)
	if err != nil {
		return Entity{}, err
	}
	req.Id = result.InsertedID.(primitive.ObjectID)

	return req, err
}

func (r *Repository) FindById(id string) (Entity, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("error di sini: %v - %s", err, objectId.Hex())
		return Entity{}, err
	}
	result := r.Config.Db.Collection(r.CollectionName).FindOne(
		context.Background(),
		bson.M{"_id": objectId},
	)
	if err != nil {
		fmt.Printf("error di bawahnya: %v", err)
		return Entity{}, err
	}

	var response = Entity{}
	err = result.Decode(&response)
	return response, err
}

func (r *Repository) FindByAuthor(author string) ([]Entity, error) {
	result, err := r.Config.Db.Collection(r.CollectionName).Find(
		context.Background(),
		bson.M{"author": author},
	)
	if err != nil {
		return []Entity{}, err
	}

	var response = []Entity{}
	err = result.Decode(&response)
	return response, err
}
