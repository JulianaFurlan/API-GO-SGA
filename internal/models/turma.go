package models

type Turma struct {
	ID         string `json:"id" binding:"required"`
	Nome       string `json:"nome" binding:"required"`
	Disciplina string `json:"disciplina" binding:"required"`
	Docente    string `json:"docente" binding:"required"`
}