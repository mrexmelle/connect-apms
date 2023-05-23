package template

type SingleResponseDto struct {
	Template Entity `json:"template"`
	Status   string `json:"status"`
}

type MultipleResponseDto struct {
	Template []Entity `json:"templates"`
	Status   string   `json:"status"`
}
