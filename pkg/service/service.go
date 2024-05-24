package service

import (
	"context"
	"strconv"
	"warehouse-api/pkg/model"
	"warehouse-api/pkg/repository"
)

// контракт для сервиса
type ServiceInterface interface {
	ReserveProducts(ctx context.Context, codes []string) error
	ReleaseProducts(ctx context.Context, codes []string) error
	GetWarehouseStock(ctx context.Context, warehouseID int) ([]model.Product, error)
	GetReservedStock(ctx context.Context, warehouseID int) ([]model.Product, error)
}

// методы для взаимодействия с хранилищем
type Service struct {
	repo *repository.Repository
}

// новый экземпляр сервиса
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

func (s *Service) GetReservedStock(ctx context.Context, warehouseID int) ([]model.Product, error) {
	return s.repo.GetReservedStock(ctx, warehouseID)
}
