package repository

import "api-gin/internal/models"
import "api-gin/internal/apperrors"
import "sync"

type MatriculaRepository interface {
	Criar(matricula *models.Matricula) error
	ListarPorTurma(turmaID string) []models.Matricula
	ListarPorAluno(alunoID string) []models.Matricula
}

type matriculaRepositoryMemoria struct {
	matriculas []models.Matricula
	mutex sync.RWMutex
}

func NovaMatriculaRepositoryMemoria() MatriculaRepository {
	return &matriculaRepositoryMemoria{
		matriculas: []models.Matricula{},
	}
}

func (r *matriculaRepositoryMemoria) Criar(matricula *models.Matricula) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, m := range r.matriculas {
		if m.AlunoID == matricula.AlunoID && m.TurmaID == matricula.TurmaID {
			return apperrors.ErrAlunoJaMatriculado
		}
	}
	r.matriculas = append(r.matriculas, *matricula)
	return nil
}

func (r *matriculaRepositoryMemoria) ListarPorTurma(turmaID string) []models.Matricula {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	var matriculas []models.Matricula
	for _, matricula := range r.matriculas {
		if matricula.TurmaID == turmaID {
			matriculas = append(matriculas, matricula)
		}
	}
	return matriculas
}

func (r *matriculaRepositoryMemoria) ListarPorAluno(alunoID string) []models.Matricula {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	var matriculas []models.Matricula
	for _, matricula := range r.matriculas {
		if matricula.AlunoID == alunoID {
			matriculas = append(matriculas, matricula)
		}
	}
	return matriculas
}
