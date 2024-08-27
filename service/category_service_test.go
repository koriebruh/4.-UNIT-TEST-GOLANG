package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"unit-test/entity"
	"unit-test/repository"
)

// <-- package ini digunakan unutk menguji kinerja dari method yang kita buat di caregory_service.go

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}} // <-- dari mock bukan dari database
var categoryService = CategoryService{Repository: categoryRepository}          // <-- set repo nya menjadi dari mock

func TestCategoryService_GetNotFound(t *testing.T) {
	// program mock -->
	categoryRepository.Mock.On("FindById", "1").Return(nil)
	// <-- ketika FindById dipanggil dengan param == 1 maka akan mereturn nil

	category, err := categoryService.Get("1")
	assert.NotNil(t, err)
	assert.Nil(t, category)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id:   "2",
		Name: "HandPhone",
	}

	categoryRepository.Mock.On("FindById", "2").Return(category)
	result, err := categoryService.Get("2")
	assert.Nil(t, err)       // <-- tidak ada error
	assert.NotNil(t, result) // <-- result punya data
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
