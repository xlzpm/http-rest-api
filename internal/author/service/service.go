package service

import (
	"context"
	"fmt"

	"github.com/xlzpm/internal/author/model"
	"github.com/xlzpm/internal/author/storage"
	modelStorage "github.com/xlzpm/internal/author/storage/model"
	"github.com/xlzpm/pkg/api/filter"
	"github.com/xlzpm/pkg/api/sort"
	"github.com/xlzpm/pkg/logging"
)

type Service struct {
	repository storage.Repository
	logger     *logging.Logger
}

func NewService(repository storage.Repository, logger *logging.Logger) *Service {
	return &Service{
		repository: repository,
		logger:     logger,
	}
}

func (s *Service) GetAll(ctx context.Context, filterOptions filter.Options, sortOptions sort.Options) ([]model.Author, error) {
	options := modelStorage.NewSortOptions(sortOptions.Field, sortOptions.Order)

	all, err := s.repository.FindAll(ctx, options)
	if err != nil {
		return nil, fmt.Errorf("failed to get all authors due to error: %v", err)
	}

	return all, nil
}
