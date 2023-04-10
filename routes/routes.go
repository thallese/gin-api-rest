package routes

import (
	"github.com/creycrey/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.GET("/:nome", controllers.Saudacao)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.Run()
}
