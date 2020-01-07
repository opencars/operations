package model

import (
	"time"
)

// Resource represents entity from the data.gov.ua web portal.
type Resource struct {
	ID           int64     `json:"id" db:"id"`
	UID          string    `json:"uid" db:"uid"`
	Name         string    `json:"name" db:"name"`
	URL          string    `json:"url" db:"url"`
	LastModified time.Time `json:"last_modified" db:"last_modified"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
