package service

import (
	"github.com/boris989/ai-assistent/test_project/model"
	"github.com/boris989/ai-assistent/test_project/repo"
)

type UserService struct {
	repo repo.UserRepo
}

func (s *UserService) CreateUser(name string) (model.User, error) {
	user := model.User{
		Name: name,
	}

	err := s.repo.Save(user)
	return user, err
}
