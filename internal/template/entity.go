package template

type Field struct {
	Type     string `json:"type"`
	Key      string `json:"key"`
	Label    string `json:"label"`
	Required bool   `json:"required"`
}

type Entity struct {
	Code        string   `json:"code"`
	Description string   `json:"description"`
	Reviewers   []string `json:"reviewers"`
	Fields      []Field  `json:"fields"`
}
