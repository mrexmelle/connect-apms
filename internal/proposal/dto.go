package proposal

type SingleResponseDto struct {
	Proposal Entity `json:"template"`
	Status   string `json:"status"`
}

type MultipleResponseDto struct {
	Proposal []Entity `json:"templates"`
	Status   string   `json:"status"`
}
