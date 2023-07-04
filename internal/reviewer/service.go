package reviewer

import (
	"github.com/mrexmelle/connect-apms/internal/config"
	"github.com/mrexmelle/connect-apms/internal/idp"
)

type Service struct {
	Config *config.Config
}

func NewService(
	cfg *config.Config,
) *Service {
	return &Service{
		Config: cfg,
	}
}

func (s *Service) RetrieveOwnTeamLeadByEhid(ehid string) idp.ProfileEntity {
	result, err := s.Config.IdpClient.GetSuperiors(ehid)

	if err != nil || len(result.Superiors) < 1 {
		return idp.ProfileEntity{}
	}

	return result.Superiors[len(result.Superiors)-1]
}

func (s *Service) RetrieveOwnProfileLeadLevelOneByEhid(ehid string) idp.ProfileEntity {
	result, err := s.Config.IdpClient.GetSuperiors(ehid)

	if err != nil || len(result.Superiors) < 1 {
		return idp.ProfileEntity{}
	}

	if result.Superiors[len(result.Superiors)-1].Ehid == ehid {
		if len(result.Superiors) > 1 {
			return result.Superiors[len(result.Superiors)-2]
		} else {
			return idp.ProfileEntity{}
		}
	} else {
		return result.Superiors[len(result.Superiors)-1]
	}
}

func (s *Service) RetrieveOwnProfileLeadLevelTwoByEhid(ehid string) idp.ProfileEntity {
	result, err := s.Config.IdpClient.GetSuperiors(ehid)

	if err != nil || len(result.Superiors) < 2 {
		return idp.ProfileEntity{}
	}

	if result.Superiors[len(result.Superiors)-1].Ehid == ehid {
		if len(result.Superiors) > 2 {
			return result.Superiors[len(result.Superiors)-3]
		} else {
			return idp.ProfileEntity{}
		}
	} else {
		return result.Superiors[len(result.Superiors)-2]
	}
}

func (s *Service) RetrieveOrganizationLeadById(id string) string {
	result, err := s.Config.IdpClient.GetOrganization(id)

	if err != nil {
		return ""
	}

	return result.Organization.LeadEhid
}

func (s *Service) RetrieveOrganizationMembersById(id string) []idp.OrganizationMemberAggregate {
	result, err := s.Config.IdpClient.GetOrganizationMembers(id)

	if err != nil {
		return []idp.OrganizationMemberAggregate{}
	}

	return result.Members
}
