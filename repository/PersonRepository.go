package repository

import (
	"project-pertama/model"

	"gorm.io/gorm"
)

type personRepository struct {
	db *gorm.DB
}

func NewPersonRepository(db *gorm.DB) *personRepository {
	return &personRepository{
		db: db,
	}
}

func (pr *personRepository) Create(newPerson model.Person) (model.Person, error) {
	tx := pr.db.Create(&newPerson)
	return newPerson, tx.Error
}

func (pr *personRepository) GetAll() ([]model.Person, error) {
	var persons = []model.Person{}

	tx := pr.db.Find(&persons)
	return persons, tx.Error
}

func (pr *personRepository) Delete(uuid string) error {
	// fmt.Println(uuid)
	tx := pr.db.Unscoped().Delete(&model.Person{}, "uuid = ?", uuid)
	return tx.Error
}
