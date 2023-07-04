package idp

type ProfileEntity struct {
	Ehid         string `json:"ehid"`
	EmployeeId   string `json:"employeeId"`
	Name         string `json:"name"`
	EmailAddress string `json:"emailAddress"`
	Dob          string `json:"dob"`
}

type OrganizationEntity struct {
	Id           string `json:"id"`
	Hierarchy    string `json:"hierarchy"`
	Name         string `json:"name"`
	LeadEhid     string `json:"leadEhid"`
	EmailAddress string `json:"emailAddress"`
}
