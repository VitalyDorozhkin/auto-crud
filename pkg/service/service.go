package service

import (
	"context"

	"github.com/VitalyDorozhkin/auto-crud/pkg/models"
)

type Service interface {
	GetAuto(ctx context.Context, id int32) (response models.AutoResponse, err error)
	CreateAuto(ctx context.Context, request models.AutoRequest) (response models.StatusResponse, err error)
	UpdateAuto(ctx context.Context, id int32, request models.AutoRequest) (response models.StatusResponse, err error)
	DeleteAuto(ctx context.Context, id int32) (response models.StatusResponse, err error)
}
