package main

import (
	"github.com/VitalyDorozhkin/auto-crud/pkg/inmemorystorage/storage"
	"github.com/VitalyDorozhkin/auto-crud/pkg/models"
	"github.com/VitalyDorozhkin/auto-crud/pkg/service"
)

func main() {
	var (
		minBrandLength int32 = 2
		maxBrandLength int32 = 32
		minModelLength int32 = 1
		maxModelLength int32 = 64
		statuses             = []string{"in transit", "in stock", "sold out", "withdrawn"}
	)

	validator := storage.NewValidator(models.InMemoryValidatorConfig{
		BrandLength: models.Interval{Min: &minBrandLength, Max: &maxBrandLength},
		ModelLength: models.Interval{Min: &minModelLength, Max: &maxModelLength},
		Statuses:    statuses,
		Mileage:     models.Interval{},
		Price:       models.Interval{},
	})

	inMemoryStorage := storage.NewStorage(validator)

	svc := service.NewService(inMemoryStorage)

	println(svc)
}
