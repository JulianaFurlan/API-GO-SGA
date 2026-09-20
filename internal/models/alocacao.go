package models

type Alocacao struct {
	ID            string `json:"id"`
	TurmaID       string `json:"turma_id"`
	SalaID        string `json:"sala_id" binding:"required"`
	DiaSemana     string `json:"dia_semana" binding:"required"`
	HorarioInicio string `json:"horario_inicio" binding:"required"`
	HorarioFim    string `json:"horario_termino" binding:"required"`
}