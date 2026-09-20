package main

import (
	"api-gin/internal/handler"
	"api-gin/internal/repository"
	"api-gin/internal/router"
	"api-gin/internal/service"
)

func main() {
	salaRepo := repository.NovaSalaRepositoryMemoria()
	alunoRepo := repository.NovoAlunoRepositoryMemoria()
	turmaRepo := repository.NovaTurmaRepositoryMemoria()
	matriculaRepo := repository.NovaMatriculaRepositoryMemoria()
	alocacaoRepo := repository.NovaAlocacaoRepositoryMemoria()

	salaService := service.NovaSalaService(salaRepo, turmaRepo, alocacaoRepo)
	alunoService := service.NovoAlunoService(alunoRepo)
	turmaService := service.NovaTurmaService(turmaRepo, alunoRepo, salaRepo, matriculaRepo, alocacaoRepo)

	salaHandler := handler.NovoSalaHandler(salaService)
	alunoHandler := handler.NovoAlunoHandler(alunoService)
	turmaHandler := handler.NovoTurmaHandler(turmaService)

	r := router.Configurar(salaHandler, alunoHandler, turmaHandler)
	r.Run(":8080")
}
