package router

import (
	"net/http"
	"time"

	"api-gin/internal/handler"

	"github.com/gin-gonic/gin"
)

func Configurar(salaHandler *handler.SalaHandler, alunoHandler *handler.AlunoHandler, turmaHandler *handler.TurmaHandler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":    "healthy",
				"timestamp": time.Now(),
				"version":   "1.0.0",
			})
		})

		v1.POST("/salas", salaHandler.Criar)
		v1.GET("/salas", salaHandler.Listar)
		v1.GET("/salas/:id/agenda", salaHandler.Agenda)

		v1.POST("/alunos", alunoHandler.Criar)
		v1.GET("/alunos", alunoHandler.Listar)
		v1.GET("/alunos/:id", alunoHandler.BuscarPorId)

		v1.POST("/turmas", turmaHandler.Criar)
		v1.GET("/turmas", turmaHandler.Listar)
		v1.POST("/turmas/:id/alunos", turmaHandler.MatricularAluno)
		v1.GET("/turmas/:id/alunos", turmaHandler.ListarAlunosDaTurma)
		v1.POST("/turmas/:id/alocar", turmaHandler.AlocarSala)
	}

	return r
}
