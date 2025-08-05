package service

import (
	"errors"
	"learning-golang/06-Unit-Test/10-mock/entity"
	"learning-golang/06-Unit-Test/10-mock/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("category not found")
	} else {
		return category, nil
	}
}
