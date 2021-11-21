package model

import "time"

// Revision represents entity in the domain.Store.
type Revision struct {
	ID          string    `json:"id" db:"id"`
	URL         string    `json:"url" db:"url"`
	FileHashSum *string   `json:"file_hash_sum" db:"file_hash_sum"`
	Deleted     int32     `json:"deleted" db:"deleted"`
	Added       int32     `json:"added" db:"added"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
