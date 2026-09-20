package handler

import (
	"net/http"
	"api-gin/internal/models"
	"api-gin/internal/service"
	"github.com/gin-gonic/gin"
)

type AlunoHandler struct {
	service *service.AlunoService
}

func NovoAlunoHandler(service *service.AlunoService) *AlunoHandler {
	return &AlunoHandler{service: service}
}

func (h *AlunoHandler) Criar(c *gin.Context) {
	var novoAluno models.Aluno
	if err := c.ShouldBindJSON(&novoAluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := h.service.Criar(&novoAluno); err != nil {
		c.JSON(statusParaErro(err), gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, novoAluno)
}

func (h *AlunoHandler) Listar(c *gin.Context) {
	alunos := h.service.Listar()
	c.JSON(http.StatusOK, alunos)
}

func (h *AlunoHandler) BuscarPorId(c *gin.Context) {
	id := c.Param("id")

	aluno, err := h.service.BuscarPorId(id)
	if err != nil {
		c.JSON(statusParaErro(err), gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, aluno)
}
