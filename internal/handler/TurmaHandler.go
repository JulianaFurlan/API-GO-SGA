package handler

import (
	"net/http"
	"api-gin/internal/models"
	"api-gin/internal/service"
	"github.com/gin-gonic/gin"
)

type TurmaHandler struct {
	service *service.TurmaService
}

func NovoTurmaHandler(service *service.TurmaService) *TurmaHandler {
	return &TurmaHandler{service: service}
}

func (h *TurmaHandler) Criar(c *gin.Context) {
	var novaTurma models.Turma
	if err := c.ShouldBindJSON(&novaTurma); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := h.service.Criar(&novaTurma); err != nil {
		c.JSON(statusParaErro(err), gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, novaTurma)
}

func (h *TurmaHandler) Listar(c *gin.Context) {
	turmas := h.service.Listar()
	c.JSON(http.StatusOK, turmas)
}

type matriculaRequest struct {
	AlunoID string `json:"aluno_id" binding:"required"`
}

func (h *TurmaHandler) MatricularAluno(c *gin.Context) {
	turmaID := c.Param("id")

	var req matriculaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := h.service.MatricularAluno(turmaID, req.AlunoID); err != nil {
		c.JSON(statusParaErro(err), gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"mensagem": "aluno matriculado com sucesso"})
}

func (h *TurmaHandler) ListarAlunosDaTurma(c *gin.Context) {
	turmaID := c.Param("id")

	alunos, err := h.service.ListarAlunosDaTurma(turmaID)
	if err != nil {
		c.JSON(statusParaErro(err), gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, alunos)
}

func (h *TurmaHandler) AlocarSala(c *gin.Context) {
	turmaID := c.Param("id")

	var novaAlocacao models.Alocacao
	if err := c.ShouldBindJSON(&novaAlocacao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := h.service.AlocarSala(turmaID, novaAlocacao); err != nil {
		c.JSON(statusParaErro(err), gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"mensagem": "sala alocada com sucesso"})
}
