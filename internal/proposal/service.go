package proposal

import (
	"strings"
	"time"

	"github.com/mrexmelle/connect-apms/internal/config"
	"github.com/mrexmelle/connect-apms/internal/event"
	"github.com/mrexmelle/connect-apms/internal/mapper"
	"github.com/mrexmelle/connect-apms/internal/reviewer"
	"github.com/mrexmelle/connect-apms/internal/template"
)

type Service struct {
	Config             *config.Config
	EventRepository    *event.Repository
	ProposalRepository *Repository
	TemplateRepository *template.Repository
	ReviewerService    *reviewer.Service
}

func NewService(
	cfg *config.Config,
	er *event.Repository,
	pr *Repository,
	tr *template.Repository,
	rs *reviewer.Service,
) *Service {
	return &Service{
		Config:             cfg,
		EventRepository:    er,
		ProposalRepository: pr,
		TemplateRepository: tr,
		ReviewerService:    rs,
	}
}

func (s *Service) RetrieveById(id string) SingleResponseDto {
	result, err := s.ProposalRepository.FindById(id)

	return SingleResponseDto{
		Proposal: result,
		Status:   mapper.ToStatus(err),
	}
}

func (s *Service) RetrieveByAuthor(ehid string) MultipleResponseDto {
	result, err := s.ProposalRepository.FindByAuthor(ehid)

	return MultipleResponseDto{
		Proposal: result,
		Status:   mapper.ToStatus(err),
	}
}

func (s *Service) Create(entity Entity) SingleResponseDto {
	var now = time.Now()

	template, err := s.TemplateRepository.FindByCode(entity.TemplateCode)
	if err != nil {
		return SingleResponseDto{
			Proposal: Entity{},
			Status:   mapper.ToStatus(err),
		}
	}
	reviewerRules := template.Reviewers
	entity.Reviewers = make([][]string, 0)
	for _, r := range reviewerRules {
		if strings.HasPrefix(r, "L:") {
			lead := s.RetrieveLead(entity.Author, r[2:])
			if lead != "" {
				entity.Reviewers = append(entity.Reviewers, []string{lead})
			}
		} else if strings.HasPrefix(r, "U:") {
			user := string(r[2:])
			if user != entity.Author {
				entity.Reviewers = append(entity.Reviewers, []string{user})
			}
		} else if strings.HasPrefix(r, "O:") {
			org := string(r[2:])
			members := s.RetrieveOrganizationMembers(org)
			if len(members) > 0 {
				entity.Reviewers = append(entity.Reviewers, members)
			}
		}
	}

	result, err := s.ProposalRepository.Create(entity)

	go s.EventRepository.Create(event.Entity{
		ProposalId: result.Id.Hex(),
		Time:       now,
		Status:     "created",
		Actor:      entity.Author,
		Note:       "",
	})

	return SingleResponseDto{
		Proposal: result,
		Status:   mapper.ToStatus(err),
	}
}

func (s *Service) RetrieveLead(author string, clause string) string {
	lead := ""
	switch clause {
	case "(OWN_TEAM)":
		{
			profile := s.ReviewerService.RetrieveOwnTeamLeadByEhid(author)
			lead = profile.Ehid
		}
	case "(OWN_PROFILE_1)":
		{
			profile := s.ReviewerService.RetrieveOwnProfileLeadLevelOneByEhid(author)
			lead = profile.Ehid
		}
	case "(OWN_PROFILE_2)":
		{
			profile := s.ReviewerService.RetrieveOwnProfileLeadLevelTwoByEhid(author)
			lead = profile.Ehid
		}
	default:
		{
			if len(clause) > 0 {
				lead = s.ReviewerService.RetrieveOrganizationLeadById(clause)
			}
		}
	}
	if lead != author {
		return lead
	} else {
		return ""
	}
}

func (s *Service) RetrieveOrganizationMembers(clause string) []string {
	result, err := s.Config.IdpClient.GetOrganizationMembers(clause)
	if err != nil {
		return []string{}
	}
	members := make([]string, 0)
	for _, m := range result.Members {
		members = append(members, m.Ehid)
	}
	return members
}
