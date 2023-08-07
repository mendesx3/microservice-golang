package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mendesx3/microservice-golang/adapter/api/model"
	"github.com/mendesx3/microservice-golang/usecase"
	"net/http"
	"strconv"
)

type GetPersonController struct {
	getPersonUseCase usecase.GetPersonUseCase
}

func NewGetPersonController(getPersonUseCase usecase.GetPersonUseCase) GetPersonController {
	return GetPersonController{
		getPersonUseCase: getPersonUseCase,
	}
}

func (p GetPersonController) GetById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	response, err := p.getPersonUseCase.Execute(id)
	if err != nil {
		c.JSON(500, gin.H{"error when creating person": err.Error()})
		return
	}

	model.NewSuccess(response, http.StatusOK).Send(c)
}
