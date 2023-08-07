package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mendesx3/microservice-golang/adapter/api"
	"github.com/mendesx3/microservice-golang/repository"
	"github.com/mendesx3/microservice-golang/usecase"
)

func Initialize() {
	// Crie uma instância de CreatePerson

	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())
	initRoutes(&r.RouterGroup)
	r.Run(":8080")

}

func initRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/api/v1")
	initUserRouter(v1)
}

func initUserRouter(v1 *gin.RouterGroup) {

	// Crie a instância de PersonRepository e CreatePersonUseCase aqui
	personRepository := repository.NewPersonRepository( /* passa a dependência do repositório aqui, tal como a conexão do DB */ )
	createPersonUseCase := usecase.NewCreatePersonUseCase(personRepository)
	getPersonUseCase := usecase.NewGetPersonUseCase(personRepository)

	// Crie a instância de CreatePersonController
	create := api.NewCreatePersonController(createPersonUseCase)
	get := api.NewGetPersonController(getPersonUseCase)

	u := v1.Group("/person")
	u.POST("/", create.Create)
	//u.PUT("/:id", api.UpdateById)
	//u.GET("/", api.ReadAll)
	u.GET("/:id", get.GetById)
	//u.DELETE("/:id", api.DeleteById)
}
