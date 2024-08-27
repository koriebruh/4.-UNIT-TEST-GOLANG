package service

import (
	"errors"
	"unit-test/entity"
	"unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

/* <-- reminder me
(service CategoryService) : Receiver artinya mengimplementasikan method get ke struct CategoryService
Get(id string)            : method yang akan di impl ke category
(*entity.Category, error) : expect return

--> summary jadi kita mengimplemetasikan method Get Ke struct CategoryService dan return dari method tersebut
category dan error
*/

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return category, errors.New("Category not found ")
	} else {
		return category, nil
	}

}
