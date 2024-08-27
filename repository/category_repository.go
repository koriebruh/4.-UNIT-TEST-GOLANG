package repository

import "unit-test/entity"

// <-- membuat method nya dulu yang akan digunakan

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
