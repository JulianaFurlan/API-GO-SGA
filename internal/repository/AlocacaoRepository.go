package repository

import "api-gin/internal/models"
import "api-gin/internal/apperrors"
import "sync"

type AlocacaoRepository interface {
	Criar(alocacao *models.Alocacao) error
	BuscarPorTurma(turmaID string) (*models.Alocacao, error)
	ListarPorSala(salaID string) []models.Alocacao
}


type alocacaoRepositoryMemoria struct {
	alocacoes map[string]models.Alocacao
	mutex sync.RWMutex
}

func NovaAlocacaoRepositoryMemoria() AlocacaoRepository {
	return &alocacaoRepositoryMemoria{
		alocacoes: make(map[string]models.Alocacao),
	}
}

func (r *alocacaoRepositoryMemoria) Criar(alocacao *models.Alocacao) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, existe := r.alocacoes[alocacao.ID]; existe {
		return apperrors.ErrAlocacaoJaExiste
	}
	r.alocacoes[alocacao.ID] = *alocacao
	return nil
}

func (r *alocacaoRepositoryMemoria) ListarPorSala(salaID string) []models.Alocacao {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var alocacoes []models.Alocacao
	for _, alocacao := range r.alocacoes {
		if alocacao.SalaID == salaID {
			alocacoes = append(alocacoes, alocacao)
		}
	}
	return alocacoes
}

func (r *alocacaoRepositoryMemoria) BuscarPorTurma(id string) (*models.Alocacao, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, alocacao := range r.alocacoes {
		if alocacao.TurmaID == id {
			return &alocacao, nil
		}
	}
	return nil, apperrors.ErrAlocacaoNaoEncontrada
}
