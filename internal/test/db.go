package test

import (
	"testing"

	"github.com/ericbutera/amalgam/internal/db"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

const UserID = "a3dce526-adf2-4f2d-bda9-e42dfd520ba5"

func NewDB(t *testing.T) *gorm.DB {
	// Note: this is a little "slow" for a unit test, but it's much nicer than maintaining mocks.
	d, err := db.NewSqlite(
		"file::memory:",
		db.WithAutoMigrate(),
		// db.WithTraceAll(),
	)
	require.NoError(t, err)
	return d
}
