package proposal

import "go.mongodb.org/mongo-driver/bson/primitive"

type Entity struct {
	Id           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Author       string             `json:"author" bson:"author"`
	TemplateCode string             `json:"templateCode" bson:"template_code"`
	Reviewers    [][]string         `json:"reviewers" bson:"reviewers"`
	Fields       map[string]string  `json:"fields" bson:"fields"`
}
