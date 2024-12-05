package pagination_test

import (
	"fmt"
	"testing"

	"github.com/ericbutera/amalgam/internal/db/pagination"
	"github.com/ericbutera/amalgam/internal/test"
	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

var (
	options = pagination.ListOptions{
		Limit: 2,
	}
	rules = []paginator.Rule{
		{Key: "ID", Order: paginator.ASC},
	}
)

type TestModel struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

func newDB(t *testing.T) *gorm.DB {
	t.Helper()
	db := test.NewDB(t)
	lo.Must0(db.AutoMigrate(&TestModel{}))
	return db
}

func seed(t *testing.T, db *gorm.DB, count int) {
	t.Helper()
	for i := 1; i <= count; i++ {
		require.NoError(t, db.Create(&TestModel{
			ID:   i,
			Name: fmt.Sprintf("test %d", i),
		}).Error)
	}
}

func TestPager_ValidPagination(t *testing.T) {
	t.Parallel()
	db := newDB(t)
	seed(t, db, 3)

	// page 1
	query := db.Model(&TestModel{})
	result, err := pagination.Pager[TestModel](query, options, rules)
	require.NoError(t, err)
	assert.Len(t, result.Results, 2)
	assert.NotEmpty(t, result.Cursor.Next)
	assert.Empty(t, result.Cursor.Previous)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, 2, result.Results[1].ID)

	// page 2
	query = db.Model(&TestModel{})
	result2, err := pagination.Pager[TestModel](query, pagination.ListOptions{
		Cursor: pagination.Cursor{Next: result.Cursor.Next},
		Limit:  2,
	}, rules)
	require.NoError(t, err)
	assert.Len(t, result2.Results, 1)
	assert.Empty(t, result2.Cursor.Next)
	assert.NotEmpty(t, result2.Cursor.Previous)
	assert.Equal(t, 3, result2.Results[0].ID)

	// back to page 1
	query = db.Model(&TestModel{})
	result3, err := pagination.Pager[TestModel](query, pagination.ListOptions{
		Cursor: pagination.Cursor{Previous: result2.Cursor.Previous},
		Limit:  2,
	}, rules)
	require.NoError(t, err)
	assert.Len(t, result3.Results, 2)
	assert.Equal(t, 1, result3.Results[0].ID)
}

func TestPager_LimitExceedsMax(t *testing.T) {
	t.Parallel()
	db := newDB(t)
	seed(t, db, 3)

	options := pagination.ListOptions{
		Limit: 200,
	}

	query := db.Model(&TestModel{})
	result, err := pagination.Pager[TestModel](query, options, rules)
	require.NoError(t, err)
	assert.Len(t, result.Results, 3)
}

func TestPager_InvalidCursor(t *testing.T) {
	t.Parallel()
	db := newDB(t)

	options := pagination.ListOptions{
		Limit:  2,
		Cursor: pagination.Cursor{Next: "invalid-cursor"},
	}

	query := db.Model(&TestModel{})
	_, err := pagination.Pager[TestModel](query, options, rules)
	assert.Error(t, err)
}

func TestPager_NoRecords(t *testing.T) {
	t.Parallel()
	db := newDB(t)
	query := db.Model(&TestModel{})
	result, err := pagination.Pager[TestModel](query, options, rules)
	require.NoError(t, err)
	assert.Empty(t, result.Results)
	assert.Equal(t, "", result.Cursor.Next)
	assert.Equal(t, "", result.Cursor.Previous)
}

func TestPager_NegativeLimit(t *testing.T) {
	t.Parallel()
	db := newDB(t)
	seed(t, db, 3)

	options := pagination.ListOptions{
		Limit: -10,
	}

	query := db.Model(&TestModel{})
	result, err := pagination.Pager[TestModel](query, options, rules)
	require.NoError(t, err)
	assert.Len(t, result.Results, 3)
}
