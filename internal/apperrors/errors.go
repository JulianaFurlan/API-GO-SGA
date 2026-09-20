package apperrors

import "errors"

var (
	// Recurso não encontrado 404
	ErrSalaNaoEncontrada     = errors.New("sala não encontrada")
	ErrAlunoNaoEncontrado    = errors.New("aluno não encontrado")
	ErrTurmaNaoEncontrada    = errors.New("turma não encontrada")
	ErrAlocacaoNaoEncontrada = errors.New("alocação não encontrada")

	// Duplicidade  409
	ErrAlunoJaMatriculado = errors.New("aluno já matriculado nesta turma")
	ErrAlunoJaExiste      = errors.New("aluno já existe")
	ErrSalaJaExiste       = errors.New("sala já existe")
	ErrTurmaJaExiste      = errors.New("turma já existe")
	ErrAlocacaoJaExiste   = errors.New("alocação já existe")

	// Capacidade insuficiente 422
	ErrCapacidadeExcedida = errors.New("capacidade máxima da sala excedida")

	// Conflito de agenda 409
	ErrConflitoHorarioSala  = errors.New("conflito de horário: sala já alocada nesse dia/horário")
	ErrConflitoHorarioAluno = errors.New("conflito de horário: aluno já matriculado em turma com horário sobreposto")
)