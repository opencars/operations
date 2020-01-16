package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/opencars/operations/pkg/handler"
	"github.com/opencars/operations/pkg/store"
	"github.com/opencars/translit"
)

func newServer(store store.Store) *server {
	srv := server{
		router: mux.NewRouter(),
		store:  store,
	}

	srv.configureRoutes()

	return &srv
}

type server struct {
	router *mux.Router
	store  store.Store
}

func (s *server) configureRoutes() {
	s.router.Handle("/api/v1/operations/{number}", s.operationsByNumber())
}

func (s *server) operationsByNumber() handler.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		number := translit.ToUA(mux.Vars(r)["number"])

		operation, err := s.store.Operation().FindByNumber(number)
		if err != nil {
			return err
		}

		if err := json.NewEncoder(w).Encode(operation); err != nil {
			return err
		}

		return nil
	}
}

// ServeHTTP implements http.Handler interface.
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.ServeHTTP(w, r)
}
