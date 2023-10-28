package apiserver

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"main/internal/app/model"
	"main/internal/app/service"
	"main/internal/app/store"
	"net/http"
)

type APIServer struct {
	config      *Config
	logger      *logrus.Logger
	router      *mux.Router
	userService *service.UserService
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.NewLogger(); err != nil {
		return nil
	}

	s.configureRouter()
	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Infof("starting server at %s", s.config.BindAddr)

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/user", s.userHandler)
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)

	if err := st.Open(); err != nil {
		return err
	}

	fmt.Println("database connected successfully")

	s.configureServices(st)

	return nil
}

func (s *APIServer) configureServices(st *store.Store) {
	s.userService = service.CreateUserService(st)
}

func (s *APIServer) NewLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		user, err := s.userService.CreateUser(&model.User{
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
	case "GET":
		user, err := s.userService.FindUserByEmail("nxgr.dev@gmail.com")
		if err != nil {
			w.WriteHeader(404)
			return
		}

		fmt.Println(user)
		w.WriteHeader(200)

	default:
		w.WriteHeader(400)
	}
}
