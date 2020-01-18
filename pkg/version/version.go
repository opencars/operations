package version

import (
	"log"
	"net/http"
	"runtime"
	"encoding/json"
)

var (
	// Version holds the current version of the app.
	Version = "dev"
)

// Handler expose version routes.
type Handler struct{}

// ServeHTTP serves HTTP.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v := struct {
		Version string `json:"version"`
		Go      string `json:"go"`
	}{
		Version: Version,
		Go:      runtime.Version(),
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("version: %v", err)
	}
}
