package repository

import (
	"github.com/mendesx3/microservice-golang/config"
	"github.com/mendesx3/microservice-golang/repository/schemas"
)

type (
	PersonRepository interface {
		CreatePerson(schemas.Person) (uint, error)
		GetPerson(id int) (*schemas.Person, error)
		GetPersons() error
		UpdatePerson() error
		DeletePerson() error
	}

	personRepository struct {
	}
)

func NewPersonRepository() PersonRepository {
	return personRepository{}
}

func (p personRepository) CreatePerson(input schemas.Person) (uint, error) {
	person := schemas.Person{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	err := config.GetDB().Create(&person).Error
	return person.ID, err
}

func (p personRepository) GetPerson(id int) (*schemas.Person, error) {
	var person schemas.Person
	err := config.GetDB().Where("id = ?", id).First(&person).Error
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func (p personRepository) GetPersons() error {
	//TODO implement me
	panic("implement me")
}

func (p personRepository) UpdatePerson() error {
	//TODO implement me
	panic("implement me")
}

func (p personRepository) DeletePerson() error {
	//TODO implement me
	panic("implement me")
}

func (p personRepository) createPerson() error {
	return nil
}
