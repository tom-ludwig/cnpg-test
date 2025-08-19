package main

import (
	"github.com/jmoiron/sqlx"
	"testing"
)

func Test_setupHTTP(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		db *sqlx.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupHTTP(tt.db)
		})
	}
}
