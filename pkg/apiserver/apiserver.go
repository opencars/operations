package apiserver

import (
	"net/http"

	"github.com/opencars/operations/pkg/config"
	"github.com/opencars/operations/pkg/logger"
	"github.com/opencars/operations/pkg/store/sqlstore"
)

// Start prepares and starts the server.
func Start(addr string, settings *config.Settings) error {
	store, err := sqlstore.New(&settings.DB)
	if err != nil {
		return err
	}

	srv := newServer(store)
	server := http.Server{
		Addr:    addr,
		Handler: srv,
	}

	logger.Info("Server is listening on %s...", addr)
	return server.ListenAndServe()
}
