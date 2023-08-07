package usecase

import (
	"github.com/mendesx3/microservice-golang/adapter/api/model"
	"github.com/mendesx3/microservice-golang/repository"
	"strconv"
)

type (
	GetPersonUseCase interface {
		Execute(id int) (*model.PersonResponse, error)
	}

	getPersonUseCase struct {
		personRepository repository.PersonRepository
	}
)

func NewGetPersonUseCase(personRepository repository.PersonRepository) GetPersonUseCase {
	return getPersonUseCase{
		personRepository: personRepository,
	}
}

func (r getPersonUseCase) Execute(id int) (*model.PersonResponse, error) {
	response, err := r.personRepository.GetPerson(id)
	if err != nil {
		return nil, err
	}

	personResponse := model.PersonResponse{
		ID:       strconv.Itoa(int(response.ID)),
		Name:     response.Name,
		Email:    response.Email,
		Password: response.Password,
	}

	return &personResponse, nil
}
