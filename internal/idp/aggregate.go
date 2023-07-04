package idp

type OrganizationMemberAggregate struct {
	Ehid           string `json:"ehid"`
	EmployeeId     string `json:"employeeId"`
	Name           string `json:"name"`
	EmailAddress   string `json:"emailAddress"`
	TitleName      string `json:"titleName"`
	EmploymentType string `json:"employmentType"`
	IsLead         bool   `json:"isLead"`
}
