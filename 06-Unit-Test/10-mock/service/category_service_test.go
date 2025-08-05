package service

import (
	"learning-golang/06-Unit-Test/10-mock/entity"
	"learning-golang/06-Unit-Test/10-mock/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryServiceNotFound(t *testing.T) {
	// Program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryServiceSuccess(t *testing.T) {
	category := entity.Category{
		Id:   "1",
		Name: "Sabun Cuci",
	}

	categoryRepository.Mock.On("FindById", "2").Return(category)

	result, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, category.Id)
	assert.Equal(t, category.Name, category.Name)
}
