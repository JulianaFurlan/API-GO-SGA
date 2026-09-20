package service

import	"api-gin/internal/models"
import "api-gin/internal/repository"


type SalaService struct {
	repo         repository.SalaRepository
	turmaRepo    repository.TurmaRepository
	alocacaoRepo repository.AlocacaoRepository
}

func NovaSalaService(repo repository.SalaRepository, turmaRepo repository.TurmaRepository, alocacaoRepo repository.AlocacaoRepository) *SalaService {
	return &SalaService{repo: repo, turmaRepo: turmaRepo, alocacaoRepo: alocacaoRepo}
}

func (s *SalaService) Criar(sala *models.Sala) error {
	return s.repo.Criar(sala)
}

func (s *SalaService) Listar() []models.Sala {
	return s.repo.ListarTodas()
}

type ItemAgenda struct {
	TurmaID       string `json:"turma_id"`
	TurmaNome     string `json:"turma_nome"`
	DiaSemana     string `json:"dia_semana"`
	HorarioInicio string `json:"horario_inicio"`
	HorarioFim    string `json:"horario_fim"`
}

func (s *SalaService) Agenda(salaID string) ([]ItemAgenda, error) {
	if _, err := s.repo.BuscarPorId(salaID); err != nil {
		return nil, err
	}

	alocacoes := s.alocacaoRepo.ListarPorSala(salaID)

	var agenda []ItemAgenda
	for _, alocacao := range alocacoes {
		turma, err := s.turmaRepo.BuscarPorId(alocacao.TurmaID)
		if err != nil {
			continue
		}
		agenda = append(agenda, ItemAgenda{
			TurmaID:       turma.ID,
			TurmaNome:     turma.Nome,
			DiaSemana:     alocacao.DiaSemana,
			HorarioInicio: alocacao.HorarioInicio,
			HorarioFim:    alocacao.HorarioFim,
		})
	}
	return agenda, nil
}
