package repo

import "test_project/model"

type UserRepo interface {
	Save(user model.User) error
	GetByID(id int) (model.User, error)
}
