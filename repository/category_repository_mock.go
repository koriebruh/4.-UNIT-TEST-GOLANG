package repository

import (
	"github.com/stretchr/testify/mock"
	"unit-test/entity"
)

// <-- package ini digunakn untuk menguji method (yang sama seperti category_service.go)
// tanpa connect ke data base

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	arguments := repository.Mock.Called(id) // <-- param yang dibutuhkan
	if arguments.Get(0) == nil {
		return nil
	}
	catrgory := arguments.Get(0).(entity.Category)
	return &catrgory
}
