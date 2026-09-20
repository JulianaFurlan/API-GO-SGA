package handler

import (
	"net/http"
	"api-gin/internal/apperrors"
)

func statusParaErro(err error) int {
	switch err {
	case apperrors.ErrSalaNaoEncontrada,
		apperrors.ErrAlunoNaoEncontrado,
		apperrors.ErrTurmaNaoEncontrada,
		apperrors.ErrAlocacaoNaoEncontrada:
		return http.StatusNotFound

	case apperrors.ErrAlunoJaMatriculado,
		apperrors.ErrSalaJaExiste,
		apperrors.ErrAlunoJaExiste,
		apperrors.ErrTurmaJaExiste,
		apperrors.ErrAlocacaoJaExiste,
		apperrors.ErrConflitoHorarioSala,
		apperrors.ErrConflitoHorarioAluno:
		return http.StatusConflict

	case apperrors.ErrCapacidadeExcedida:
		return http.StatusUnprocessableEntity

	default:
		return http.StatusInternalServerError
	}
}
