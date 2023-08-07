package usecase

import (
	"context"
	"github.com/mendesx3/microservice-golang/repository"
	"github.com/mendesx3/microservice-golang/repository/schemas"
)

type CreatePerson struct {
	Name     string
	Email    string
	Password string
}

type CreatePersonUseCase interface {
	Execute(context.Context, CreatePerson) (uint, error)
}

type createPersonUseCase struct {
	personRepository repository.PersonRepository
}

func NewCreatePersonUseCase(personRepository repository.PersonRepository) CreatePersonUseCase {
	return createPersonUseCase{
		personRepository: personRepository,
	}
}

func (r createPersonUseCase) Execute(ctx context.Context, request CreatePerson) (uint, error) {
	p := schemas.Person{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	id, err := r.personRepository.CreatePerson(p)
	if err != nil {
		return 0, err
	}
	return id, nil
}
