package service

import	"api-gin/internal/apperrors"
import	"api-gin/internal/models"
import "api-gin/internal/repository"

type TurmaService struct {
	turmaRepo      repository.TurmaRepository
	alunoRepo     repository.AlunoRepository
	salaRepo      repository.SalaRepository
	matriculaRepo repository.MatriculaRepository
	alocacaoRepo  repository.AlocacaoRepository
}

func NovaTurmaService(
	turmaRepo repository.TurmaRepository,
	alunoRepo repository.AlunoRepository,
	salaRepo repository.SalaRepository,
	matriculaRepo repository.MatriculaRepository,
	alocacaoRepo repository.AlocacaoRepository,
) *TurmaService {
	return &TurmaService{
		turmaRepo:     turmaRepo,
		alunoRepo:     alunoRepo,
		salaRepo:      salaRepo,
		matriculaRepo: matriculaRepo,
		alocacaoRepo:  alocacaoRepo,
	}
}

func (s *TurmaService) Criar(turma *models.Turma) error {
	return s.turmaRepo.Criar(turma)
}

type TurmaResumo struct {
	ID               string `json:"id"`
	Nome             string `json:"nome"`
	Disciplina       string `json:"disciplina"`
	Docente          string `json:"docente"`
	QuantidadeAlunos int    `json:"quantidade_alunos"`
	Alocada          bool   `json:"alocada"`
}

func (s *TurmaService) Listar() []TurmaResumo {
	turmas := s.turmaRepo.ListarTodas()

	var resumos []TurmaResumo
	for _, turma := range turmas {
		matriculas := s.matriculaRepo.ListarPorTurma(turma.ID)
		_, errAlocacao := s.alocacaoRepo.BuscarPorTurma(turma.ID)

		resumos = append(resumos, TurmaResumo{
			ID:               turma.ID,
			Nome:             turma.Nome,
			Disciplina:       turma.Disciplina,
			Docente:          turma.Docente,
			QuantidadeAlunos: len(matriculas),
			Alocada:          errAlocacao == nil,
		})
	}
	return resumos
}

func (s *TurmaService) MatricularAluno(turmaID string, alunoID string) error {
	if _, err := s.turmaRepo.BuscarPorId(turmaID); err != nil {
		return err
	}
	if _, err := s.alunoRepo.BuscarPorId(alunoID); err != nil {
		return err
	}

	matriculasDaTurma := s.matriculaRepo.ListarPorTurma(turmaID)

	for _, m := range matriculasDaTurma {
		if m.AlunoID == alunoID {
			return apperrors.ErrAlunoJaMatriculado
		}
	}

	alocacaoAtual, errAlocacao := s.alocacaoRepo.BuscarPorTurma(turmaID)
	if errAlocacao == nil {
		sala, err := s.salaRepo.BuscarPorId(alocacaoAtual.SalaID)
		if err != nil {
			return err
		}

		if len(matriculasDaTurma)+1 > sala.Capacidade {
			return apperrors.ErrCapacidadeExcedida
		}

		matriculasDoAluno := s.matriculaRepo.ListarPorAluno(alunoID)
		for _, m := range matriculasDoAluno {
			alocacaoOutraTurma, err := s.alocacaoRepo.BuscarPorTurma(m.TurmaID)
			if err != nil {
				continue
			}
			if horariosSobrepoem(alocacaoAtual.HorarioInicio, alocacaoAtual.HorarioFim, alocacaoOutraTurma.HorarioInicio, alocacaoOutraTurma.HorarioFim) {
				return apperrors.ErrConflitoHorarioAluno
			}
		}
	}

	novaMatricula := models.Matricula{TurmaID: turmaID, AlunoID: alunoID}
	return s.matriculaRepo.Criar(&novaMatricula)
}

func (s *TurmaService) ListarAlunosDaTurma(turmaID string) ([]models.Aluno, error) {
	if _, err := s.turmaRepo.BuscarPorId(turmaID); err != nil {
		return nil, err
	}

	matriculas := s.matriculaRepo.ListarPorTurma(turmaID)

	var alunos []models.Aluno
	for _, m := range matriculas {
		aluno, err := s.alunoRepo.BuscarPorId(m.AlunoID)
		if err == nil {
			alunos = append(alunos, *aluno)
		}
	}
	return alunos, nil
}

func (s *TurmaService) AlocarSala(turmaID string, novaAlocacao models.Alocacao) error {
	if _, err := s.turmaRepo.BuscarPorId(turmaID); err != nil {
		return err
	}

	sala, err := s.salaRepo.BuscarPorId(novaAlocacao.SalaID)
	if err != nil {
		return err
	}

	matriculas := s.matriculaRepo.ListarPorTurma(turmaID)
	if len(matriculas) > sala.Capacidade {
		return apperrors.ErrCapacidadeExcedida
	}

	alocacoesDaSala := s.alocacaoRepo.ListarPorSala(novaAlocacao.SalaID)
	for _, existente := range alocacoesDaSala {
		if existente.TurmaID == turmaID {
			continue
		}
		if existente.DiaSemana != novaAlocacao.DiaSemana {
			continue
		}
		if horariosSobrepoem(novaAlocacao.HorarioInicio, novaAlocacao.HorarioFim, existente.HorarioInicio, existente.HorarioFim) {
			return apperrors.ErrConflitoHorarioSala
		}
	}

	novaAlocacao.ID = turmaID
	novaAlocacao.TurmaID = turmaID
	return s.alocacaoRepo.Criar(&novaAlocacao)
}

func horariosSobrepoem(novaInicio, novaFim, existenteInicio, existenteFim string) bool {
	return novaInicio < existenteFim && novaFim > existenteInicio
}
