package model

import (
	"github.com/xlzpm/internal/author/storage"
	"github.com/xlzpm/pkg/api/filter"
)

type filterOptions struct {
	limit  int
	fields []filter.Field
}

func NewOptions(options filter.Options) storage.FilterOptions {
	return &filterOptions{limit: options.GetLimit(), fields: options.Fields()}
}
