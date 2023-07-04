package idp

type SuperiorResponseDto struct {
	Superiors []ProfileEntity `json:"superiors"`
	Status    string          `json:"status"`
}

type OrganizationSingleResponseDto struct {
	Organization OrganizationEntity `json:"organization"`
	Status       string             `json:"status"`
}

type OrganizationMemberResponseDto struct {
	Members []OrganizationMemberAggregate `json:"members"`
	Status  string                        `json:"status"`
}
