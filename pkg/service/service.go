package service

import (
	"context"

	"github.com/VitalyDorozhkin/auto-crud/pkg/models"
)

type storage interface {
	GetAuto(ctx context.Context, id uint32) (auto models.Auto, err error)
	CreateAuto(ctx context.Context, auto models.AutoRequest) (id uint32, err error)
	UpdateAuto(ctx context.Context, auto models.Auto) (err error)
	DeleteAuto(ctx context.Context, id uint32) (err error)
}

type Service interface {
	GetAuto(ctx context.Context, id uint32) (response models.AutoResponse, err error)
	CreateAuto(ctx context.Context, request *models.AutoRequest) (response models.IDResponse, err error)
	UpdateAuto(ctx context.Context, id uint32, request *models.AutoRequest) (response models.StatusResponse, err error)
	DeleteAuto(ctx context.Context, id uint32) (response models.StatusResponse, err error)
}

type service struct {
	storage storage
}

func (s *service) GetAuto(ctx context.Context, id uint32) (response models.AutoResponse, err error) {
	auto, err := s.storage.GetAuto(ctx, id)
	if err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Data = &auto
	return
}

func (s *service) CreateAuto(ctx context.Context, request *models.AutoRequest) (response models.IDResponse, err error) {
	id, err := s.storage.CreateAuto(ctx, *request)
	if err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Data = &models.IDStruct{ID: id}
	return
}

func (s *service) UpdateAuto(ctx context.Context, id uint32, request *models.AutoRequest) (response models.StatusResponse, err error) {
	auto := models.Auto{
		ID:      id,
		Brand:   request.Brand,
		Model:   request.Model,
		Price:   request.Price,
		Status:  request.Status,
		Mileage: request.Mileage,
	}
	err = s.storage.UpdateAuto(ctx, auto)
	if err != nil {
		response.Error = err.Error()
		return response, err
	}
	res := true
	response.Data = &res
	return
}

func (s *service) DeleteAuto(ctx context.Context, id uint32) (response models.StatusResponse, err error) {
	err = s.storage.DeleteAuto(ctx, id)
	if err != nil {
		response.Error = err.Error()
		return response, err
	}
	res := true
	response.Data = &res
	return
}

func NewService(storage storage) Service {
	return &service{storage: storage}
}
