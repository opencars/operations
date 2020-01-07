package model

import (
	"strings"
	"time"

	"github.com/opencars/govdata"
)

// Revision represents entity in the store.Store.
type Revision struct {
	ID          string    `json:"id" db:"id"`
	URL         string    `json:"url" db:"url"`
	FileHashSum *string   `json:"file_hash_sum" db:"file_hash_sum"`
	Deleted     int32     `json:"deleted" db:"deleted"`
	Added       int32     `json:"added" db:"added"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// RevisionFromGov returns new instance of Revision from govdata.Revision.
func RevisionFromGov(revision *govdata.Revision) *Revision {
	parts := strings.Split(revision.URL, "/")

	return &Revision{
		ID:          parts[len(parts)-1],
		URL:         revision.URL,
		FileHashSum: revision.FileHashSum,
		CreatedAt:   revision.ResourceCreated.Time,
	}
}
