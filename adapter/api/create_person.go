package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mendesx3/microservice-golang/adapter/api/model"
	"github.com/mendesx3/microservice-golang/usecase"
	"net/http"
)

type CreatePersonController struct {
	createPersonUseCase usecase.CreatePersonUseCase
}

func NewCreatePersonController(createPersonUseCase usecase.CreatePersonUseCase) CreatePersonController {
	return CreatePersonController{
		createPersonUseCase: createPersonUseCase,
	}
}

func (p CreatePersonController) Create(c *gin.Context) {
	var input model.CreatePersonInput
	if err := json.NewDecoder(c.Request.Body).Decode(&input); err != nil {
		c.JSON(400, gin.H{"error when decoding json": err.Error()})
		return
	}
	defer c.Request.Body.Close()

	personRequest := usecase.CreatePerson{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	id, err := p.createPersonUseCase.Execute(c.Request.Context(), personRequest)
	if err != nil {
		c.JSON(500, gin.H{"error when creating person": err.Error()})
		return
	}

	model.NewSuccess(id, http.StatusCreated).Send(c)
}
