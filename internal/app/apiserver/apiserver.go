package apiserver

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"main/internal/app/model"
	"main/internal/app/store"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
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

	s.store = st
	fmt.Println("database connected successfully")

	return nil
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
		if user, err := s.store.User().Create(&model.User{
			Email:             "nxgr.dev@gmail.com",
			EncryptedPassword: "5658",
		}); err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			return
		} else {
			fmt.Println(user)
		}

		w.WriteHeader(200)
	case "GET":
		if _, err := fmt.Println(s.store.User().FindByEmail("nxgr2.dev@gmail.com")); err != nil {
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(200)
	default:
		w.WriteHeader(400)
	}
}
