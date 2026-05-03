package repo

import "github.com/boris989/ai-assistent/test_project/model"

type UserRepo interface {
	Save(user model.User) error
	GetByID(id int) (model.User, error)
}
