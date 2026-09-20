package service

import	"api-gin/internal/models"
import 	"api-gin/internal/repository"


type AlunoService struct {
	repo repository.AlunoRepository
}

func NovoAlunoService(repo repository.AlunoRepository) *AlunoService {
	return &AlunoService{repo: repo}
}

func (s *AlunoService) Criar(aluno *models.Aluno) error {
	return s.repo.Criar(aluno)
}

func (s *AlunoService) Listar() []models.Aluno {
	return s.repo.ListarTodas()
}

func (s *AlunoService) BuscarPorId(id string) (*models.Aluno, error) {
	return s.repo.BuscarPorId(id)
}
