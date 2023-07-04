package template

type Field struct {
	Type     string `json:"type" bson:"type"`
	Key      string `json:"key" bson:"key"`
	Label    string `json:"label" bson:"label"`
	Required bool   `json:"required" bson:"required"`
}

type Entity struct {
	Code        string   `json:"code" bson:"code" validate:"required"`
	Description string   `json:"description" bson:"description" validate:"required"`
	Reviewers   []string `json:"reviewers" bson:"reviewers" validate:"required"`
	Fields      []Field  `json:"fields" bson:"fields" validate:"required"`
}
