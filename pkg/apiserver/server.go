package apiserver

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/opencars/operations/pkg/handler"
	"github.com/opencars/operations/pkg/store"
	"github.com/opencars/operations/pkg/version"
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
	// api := s.router.PathPrefix("").Subrouter()
	// GET /api/v1/operations/version
	s.router.Handle("/api/v1/operations/version", version.Handler{}).Methods("GET", "OPTIONS")

	// GET /api/v1/operations?number={number}.
	s.router.Handle("/api/v1/operations", s.operationsByNumber()).Queries("number", "{number}").Methods("GET", "OPTIONS")
}

func (s *server) operationsByNumber() handler.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		number := strings.ToUpper(translit.ToUA(r.URL.Query().Get("number")))

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
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"Api-Key"})

	cors := handlers.CORS(origins, methods, headers)(s.router)
	cors.ServeHTTP(w, r)
}
