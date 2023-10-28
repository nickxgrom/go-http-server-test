package controller

import (
	"fmt"
	"main/internal/app/model"
	"main/internal/app/service"
	"main/internal/app/store"
	"net/http"
)

type UserController struct {
	service *service.UserService
}

func CreateUserController(s *store.Store) *UserController {
	return &UserController{
		service: service.CreateUserService(s),
	}
}

func (c *UserController) HandleFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		c.create(w, r)
	case "GET":
		c.findByEmail(w, r)
	default:
		w.WriteHeader(400)
	}
}

func (c *UserController) create(w http.ResponseWriter, r *http.Request) {
	user, err := c.service.CreateUser(&model.User{
		Email:    "nxgr.dev@gmail.com",
		Password: "5658",
	})
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}

	fmt.Println(user)
	w.WriteHeader(200)
}

func (c *UserController) findByEmail(w http.ResponseWriter, r *http.Request) {
	user, err := c.service.FindUserByEmail("nxgr.dev@gmail.com")
	if err != nil {
		w.WriteHeader(404)
		return
	}

	fmt.Println(user)
	w.WriteHeader(200)
}
