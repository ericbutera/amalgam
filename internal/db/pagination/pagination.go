package pagination

import (
	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

const DefaultLimit = 25

type ListOptions struct {
	Cursor Cursor
	Limit  int // Number of records per page
	// Filters TODO
}

type Cursor struct {
	Previous string
	Next     string
}

type Result[T any] struct {
	Cursor  Cursor
	Results []T
}

func Pager[T any](query *gorm.DB, options ListOptions, rules []paginator.Rule) (*Result[T], error) {
	if options.Limit <= 0 || options.Limit > 100 {
		options.Limit = DefaultLimit
	}

	p := paginator.New(&paginator.Config{
		Rules: rules,
		Limit: options.Limit,
	})

	if options.Cursor.Previous != "" {
		p.SetBeforeCursor(options.Cursor.Previous)
	}

	if options.Cursor.Next != "" {
		p.SetAfterCursor(options.Cursor.Next)
	}

	var dest []T

	result, cursor, err := p.Paginate(query, &dest)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &Result[T]{
		Cursor: Cursor{
			Previous: lo.FromPtr(cursor.Before),
			Next:     lo.FromPtr(cursor.After),
		},
		Results: dest,
	}, err
}
