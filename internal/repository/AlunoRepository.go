package repository

import "api-gin/internal/models"
import "api-gin/internal/apperrors"
import "sync"

type AlunoRepository interface {
	Criar (aluno *models.Aluno) error
	ListarTodas() []models.Aluno
	BuscarPorId(id string) (*models.Aluno, error)
}

type alunoRepositoryMemoria struct {
	alunos map[string]models.Aluno
	mutex sync.RWMutex
}

func NovoAlunoRepositoryMemoria() AlunoRepository {
	return &alunoRepositoryMemoria{
		alunos: make(map[string]models.Aluno),
	}
}

func (r *alunoRepositoryMemoria) Criar(aluno *models.Aluno) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, existe := r.alunos[aluno.Matricula]; existe {
		return apperrors.ErrAlunoJaExiste
	}
	r.alunos[aluno.Matricula] = *aluno
	return nil
}

func (r *alunoRepositoryMemoria) ListarTodas() []models.Aluno {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var alunos []models.Aluno
	for _, aluno := range r.alunos {
		alunos = append(alunos, aluno)
	}
	return alunos
}

func (r *alunoRepositoryMemoria) BuscarPorId(id string) (*models.Aluno, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if aluno, existe := r.alunos[id]; existe {
		return &aluno, nil
	}
	return nil, apperrors.ErrAlunoNaoEncontrado
}
