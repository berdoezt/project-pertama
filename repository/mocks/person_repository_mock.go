package mocks

import (
	"project-pertama/model"

	"github.com/stretchr/testify/mock"
)

type personRepositoryMock struct {
	mock.Mock
}

func NewPersonRepositoryMock() *personRepositoryMock {
	return &personRepositoryMock{}
}

func (pr *personRepositoryMock) Create(newPerson model.Person) (model.Person, error) {

	ret := pr.Called(newPerson)

	if ret.Get(0) == nil {
		return model.Person{}, nil
	}

	if ret.Get(1) == nil {
		return model.Person{}, nil
	}

	return ret.Get(0).(model.Person), ret.Get(1).(error)
}

func (pr *personRepositoryMock) GetAll() ([]model.Person, error) {
	var persons = []model.Person{
		model.Person{
			Name: "Budi",
			UUID: "123",
		},
		model.Person{
			Name: "Ani",
			UUID: "345",
		},
	}

	return persons, nil
}

func (pr *personRepositoryMock) Delete(uuid string) error {
	return nil
}
