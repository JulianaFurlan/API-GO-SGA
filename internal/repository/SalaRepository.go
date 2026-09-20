package repository

import "api-gin/internal/models"
import "api-gin/internal/apperrors"
import "sync"

type SalaRepository interface {
	Criar (sala *models.Sala) error
	ListarTodas() []models.Sala
	BuscarPorId(id string) (*models.Sala, error)
}

type salaRepositoryMemoria struct {
	salas map[string]models.Sala
	mutex sync.RWMutex
}

func NovaSalaRepositoryMemoria() SalaRepository {
	return &salaRepositoryMemoria{
		salas: make(map[string]models.Sala),
	}
}

func (r *salaRepositoryMemoria) Criar(sala *models.Sala) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, existe := r.salas[sala.ID]; existe {
		return apperrors.ErrSalaJaExiste
	}
	r.salas[sala.ID] = *sala
	return nil
}

func (r *salaRepositoryMemoria) ListarTodas() []models.Sala {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var salas []models.Sala
	for _, sala := range r.salas {
		salas = append(salas, sala)
	}
	return salas
}

func (r *salaRepositoryMemoria) BuscarPorId(id string) (*models.Sala, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if sala, existe := r.salas[id]; existe {
		return &sala, nil
	}
	return nil, apperrors.ErrSalaNaoEncontrada
}