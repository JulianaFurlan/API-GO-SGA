package handler

import (
	"net/http"
	"api-gin/internal/models"
	"api-gin/internal/service"
	"github.com/gin-gonic/gin"
)

type SalaHandler struct {
	service *service.SalaService
}

func NovoSalaHandler(service *service.SalaService) *SalaHandler {
	return &SalaHandler{service: service}
}

func (h *SalaHandler) Criar(c *gin.Context) {
	var novaSala models.Sala
	if err := c.ShouldBindJSON(&novaSala); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := h.service.Criar(&novaSala); err != nil {
		c.JSON(statusParaErro(err), gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, novaSala)
}

func (h *SalaHandler) Listar(c *gin.Context) {
	salas := h.service.Listar()
	c.JSON(http.StatusOK, salas)
}

func (h *SalaHandler) Agenda(c *gin.Context) {
	id := c.Param("id")

	agenda, err := h.service.Agenda(id)
	if err != nil {
		c.JSON(statusParaErro(err), gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, agenda)
}
