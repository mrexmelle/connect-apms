package template

import (
	"github.com/mrexmelle/connect-apms/internal/config"
	"github.com/mrexmelle/connect-apms/internal/mapper"
)

type Service struct {
	Config             *config.Config
	TemplateRepository *Repository
}

func NewService(
	cfg *config.Config,
	tr *Repository,
) *Service {
	return &Service{
		Config:             cfg,
		TemplateRepository: tr,
	}
}

func (s *Service) RetrieveByCode(code string) SingleResponseDto {
	result, err := s.TemplateRepository.FindByCode(code)

	return SingleResponseDto{
		Template: result,
		Status:   mapper.ToStatus(err),
	}
}

func (s *Service) RetrieveAll() MultipleResponseDto {
	result, err := s.TemplateRepository.FindAll()

	return MultipleResponseDto{
		Template: result,
		Status:   mapper.ToStatus(err),
	}
}
