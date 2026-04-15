package handler

import "test_project/service"

type UserHandler struct {
	service *service.UserService
}

func (h *UserHandler) HandleCreateUser(name string) {
	h.service.CreateUser(name)
}
