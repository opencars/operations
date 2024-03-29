package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/opencars/httputil"

	"github.com/opencars/operations/pkg/domain"
	"github.com/opencars/operations/pkg/domain/query"
)

type server struct {
	router *mux.Router

	svc domain.CustomerService
}

func newServer(svc domain.CustomerService) *server {
	srv := server{
		router: mux.NewRouter(),
		svc:    svc,
	}

	srv.configureRoutes()

	return &srv
}

func (s *server) listByNumber() httputil.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		q := query.ListByNumber{
			UserID:  UserIDFromContext(r.Context()),
			TokenID: TokenIDFromContext(r.Context()),
			Number:  r.URL.Query().Get("number"),
		}

		operations, err := s.svc.ListByNumber(r.Context(), &q)
		if err != nil {
			return handleErr(err)
		}

		return json.NewEncoder(w).Encode(operations)
	}
}

func (s *server) listByVIN() httputil.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		q := query.ListByVIN{
			UserID:  UserIDFromContext(r.Context()),
			TokenID: TokenIDFromContext(r.Context()),
			VIN:     r.URL.Query().Get("vin"),
		}

		operations, err := s.svc.ListByVIN(r.Context(), &q)
		if err != nil {
			return handleErr(err)
		}

		return json.NewEncoder(w).Encode(operations)
	}
}
