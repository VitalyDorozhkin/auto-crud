package storage

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/VitalyDorozhkin/auto-crud/pkg/models"
)

type Storage interface {
	GetAuto(ctx context.Context, id uint32) (auto models.Auto, err error)
	CreateAuto(ctx context.Context, auto models.CreateAutoRequest) (id uint32, err error)
	UpdateAuto(ctx context.Context, id uint32, auto models.UpdateAutoRequest) (err error)
	DeleteAuto(ctx context.Context, id uint32) (err error)
}

type storage struct {
	items     map[uint32]models.Auto
	validator Validator
}

func (s *storage) GetAuto(ctx context.Context, id uint32) (auto models.Auto, err error) {
	return s.items[id], nil
}

func (s *storage) CreateAuto(ctx context.Context, auto models.CreateAutoRequest) (id uint32, err error) {
	id = uuid.New().ID()
	if _, ok := s.items[id]; ok {
		return 0, fmt.Errorf("generate id error")
	}
	s.items[id] = models.Auto{
		ID:      id,
		Brand:   auto.Brand,
		Model:   auto.Model,
		Price:   auto.Price,
		Status:  auto.Status,
		Mileage: auto.Mileage,
	}
	return id, nil
}

func (s *storage) UpdateAuto(ctx context.Context, id uint32, auto models.UpdateAutoRequest) (err error) {
	if _, ok := s.items[id]; !ok {
		return fmt.Errorf("no such auto")
	}
	s.items[id], err = s.copyWithValidate(s.items[id], auto)
	return
}

func (s *storage) DeleteAuto(ctx context.Context, id uint32) (err error) {
	delete(s.items, id)
	return nil
}

func (s *storage) copyWithValidate(dest models.Auto, source models.UpdateAutoRequest) (res models.Auto, err error) {
	res = dest
	if source.Brand != nil {
		if !s.validator.CheckBrand(*source.Brand) {
			return dest, fmt.Errorf("invalid brand")
		}
		res.Brand = *source.Brand
	}
	if source.Model != nil {
		if !s.validator.CheckModel(*source.Model) {
			return dest, fmt.Errorf("invalid model")
		}
		res.Model = *source.Model
	}
	if source.Status != nil {
		if !s.validator.CheckStatus(*source.Status) {
			return dest, fmt.Errorf("invalid status")
		}
		res.Status = *source.Status
	}
	if source.Mileage != nil {
		if !s.validator.CheckMileage(uint32(*source.Mileage)) {
			return dest, fmt.Errorf("invalid mileage")
		}
		res.Mileage = *source.Mileage
	}
	if source.Price != nil {
		if !s.validator.CheckPrice(uint32(*source.Price)) {
			return dest, fmt.Errorf("invalid price")
		}
		res.Price = *source.Price
	}
	return
}

func NewStorage(validator Validator) Storage {
	return &storage{
		items:     make(map[uint32]models.Auto),
		validator: validator,
	}
}
