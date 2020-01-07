package model

import (
	"testing"
)

func TestResource(t *testing.T) *Resource {
	t.Helper()

	return &Resource{}
}

func TestRevision(t *testing.T) *Revision {
	t.Helper()

	return &Revision{}
}

func TestOperation(t *testing.T) *Operation {
	t.Helper()

	return &Operation{}
}
