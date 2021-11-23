package http

import (
	"github.com/opencars/operations/pkg/version"
)

func (s *server) configureRoutes() {
	// GET /api/v1/operations/version.
	s.router.Handle("/api/v1/operations/version", version.Handler{}).Methods("GET")

	// GET /api/v1/operations?number={number}.
	s.router.Handle("/api/v1/operations", s.operationsByNumber()).Queries("number", "{number}").Methods("GET")

	// GET /api/v1/operations?vin={vin}.
	s.router.Handle("/api/v1/operations", s.operationsByNumber()).Queries("vin", "{vin}").Methods("GET")
}
