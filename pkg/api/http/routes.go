package http

import (
	"github.com/opencars/operations/pkg/version"
)

func (s *server) configureRoutes() {
	// GET /api/v1/operations/version.
	s.router.Handle("/api/v1/operations/version", version.Handler{}).Methods("GET", "OPTIONS")

	// GET /api/v1/operations?number={number}.
	s.router.Handle("/api/v1/operations", s.operationsByNumber()).Queries("number", "{number}").Methods("GET", "OPTIONS")
}
