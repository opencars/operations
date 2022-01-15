package http

import (
	"github.com/opencars/operations/pkg/version"
)

func (s *server) configureRoutes() {
	// GET /api/v1/operations/version.
	s.router.Handle("/api/v1/operations/version", version.Handler{}).Methods("GET")

	router := s.router.PathPrefix("/api/v1/").Subrouter()
	router.Use(
		AuthorizationMiddleware(),
	)

	// GET /api/v1/operations?number={number}.
	router.Handle("/operations", s.listByNumber()).Queries("number", "{number}").Methods("GET")

	// GET /api/v1/operations?vin={vin}.
	router.Handle("/operations", s.listByVIN()).Queries("vin", "{vin}").Methods("GET")
}
