package models

type Aluno struct {
	Nome               string `json:"nome" binding:"required"`
	Matricula          string `json:"matricula" binding:"required"`
	EmailInstitucional string `json:"email_institucional" binding:"required,email"`
}
