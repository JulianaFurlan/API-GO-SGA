package repository

import "api-gin/internal/models"
import "api-gin/internal/apperrors"
import "sync"

type TurmaRepository interface {
	Criar (turma *models.Turma) error
	ListarTodas() []models.Turma
	BuscarPorId(id string) (*models.Turma, error)
}

type turmaRepositoryMemoria struct {
	turmas map[string]models.Turma
	mutex sync.RWMutex
}

func NovaTurmaRepositoryMemoria() TurmaRepository {
	return &turmaRepositoryMemoria{
		turmas: make(map[string]models.Turma),
	}
}

func (r *turmaRepositoryMemoria) Criar(turma *models.Turma) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, existe := r.turmas[turma.ID]; existe {
		return apperrors.ErrTurmaJaExiste
	}
	r.turmas[turma.ID] = *turma
	return nil
}

func (r *turmaRepositoryMemoria) ListarTodas() []models.Turma {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var turmas []models.Turma
	for _, turma := range r.turmas {
		turmas = append(turmas, turma)
	}
	return turmas
}

func (r *turmaRepositoryMemoria) BuscarPorId(id string) (*models.Turma, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if turma, existe := r.turmas[id]; existe {
		return &turma, nil
	}
	return nil, apperrors.ErrTurmaNaoEncontrada
}
