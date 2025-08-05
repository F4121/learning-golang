package repository

import "learning-golang/06-Unit-Test/10-mock/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
