package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/opencars/translit"

	"github.com/opencars/operations/pkg/handler"
	"github.com/opencars/operations/pkg/store"
)

const (
	ascending  string = "ASC"
	descending string = "DESC"
)

const defaultLimit = 10

type server struct {
	router *mux.Router
	store  store.Store
}

func newServer(s store.Store) *server {
	srv := server{
		router: mux.NewRouter(),
		store:  s,
	}

	srv.configureRoutes()

	return &srv
}

func (s *server) order(r *http.Request) (string, error) {
	order := strings.ToUpper(r.URL.Query().Get("order"))
	if order == "" {
		return descending, nil
	}

	if order != ascending && order != descending {
		return "", handler.ErrInvalidOrder
	}

	return order, nil
}

func (s *server) limit(r *http.Request) (uint64, error) {
	limit := strings.ToUpper(r.URL.Query().Get("limit"))
	if limit == "" {
		return defaultLimit, nil
	}

	num, err := strconv.ParseUint(limit, 10, 64)
	if err != nil {
		return 0, handler.ErrInvalidLimit
	}

	return num, nil
}

func (s *server) operationsByNumber() handler.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		number := translit.ToUA(strings.ToUpper(r.URL.Query().Get("number")))

		limit, err := s.limit(r)
		if err != nil {
			return err
		}

		order, err := s.order(r)
		if err != nil {
			return err
		}

		operations, err := s.store.Operation().FindByNumber(number, limit, order)
		if err != nil {
			return err
		}

		for i := range operations {
			operations[i].Person = operations[i].PrettyPerson()
		}

		return json.NewEncoder(w).Encode(operations)
	}
}

// ServeHTTP implements http.Handler interface.
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"X-Api-Key", "Api-Key"})

	cors := handlers.CORS(origins, methods, headers)(s.router)
	cors.ServeHTTP(w, r)
}
