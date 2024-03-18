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

	r := pr.Called(newPerson)

	var r1 model.Person
	var r2 error

	if r.Get(0) != nil {
		r1 = r.Get(0).(model.Person)
	}

	if r.Get(1) != nil {
		r2 = r.Get(1).(error)
	}

	return r1, r2
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
