package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/opencars/httputil"

	"github.com/opencars/operations/pkg/domain"
)

const (
	ascending  string = "ASC"
	descending string = "DESC"
)

const defaultLimit = 10

type server struct {
	router *mux.Router

	svc domain.UserOperationService
}

func newServer(svc domain.UserOperationService) *server {
	srv := server{
		router: mux.NewRouter(),
		svc:    svc,
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
		return "", ErrInvalidOrder
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
		return 0, ErrInvalidLimit
	}

	return num, nil
}

func (s *server) operationsByNumber() httputil.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		number := strings.ToUpper(r.URL.Query().Get("number"))

		limit, err := s.limit(r)
		if err != nil {
			return err
		}

		order, err := s.order(r)
		if err != nil {
			return err
		}

		operations, err := s.svc.FindByNumber(r.Context(), number, limit, order)
		if err != nil {
			return err
		}

		return json.NewEncoder(w).Encode(operations)
	}
}
