package service

import (
	"context"
	"strconv"
	"warehouse-api/pkg/model"
	"warehouse-api/pkg/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ReserveProducts(ctx context.Context, codes []string) error {
	return s.repo.ReserveProducts(ctx, codes)
}

func (s *Service) ReleaseProducts(ctx context.Context, codes []string) error {
	return s.repo.ReleaseProducts(ctx, codes)
}

func (s *Service) GetWarehouseStock(ctx context.Context, warehouseID int) ([]model.Product, error) {
	return s.repo.GetStock(ctx, strconv.Itoa(warehouseID))
}
