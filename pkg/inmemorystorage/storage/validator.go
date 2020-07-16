package storage

import "github.com/VitalyDorozhkin/auto-crud/pkg/models"

type Validator interface {
	CheckBrand(brand string) bool
	CheckModel(model string) bool
	CheckStatus(status string) bool
	CheckMileage(mileage uint32) bool
	CheckPrice(price uint32) bool
}

type validator struct {
	config models.InMemoryValidatorConfig
}

func (v *validator) CheckBrand(brand string) bool {
	return v.config.BrandLength.InInterval(int32(len(brand)))
}

func (v *validator) CheckModel(model string) bool {
	return v.config.ModelLength.InInterval(int32(len(model)))
}

func (v *validator) CheckStatus(status string) bool {
	for _, v := range v.config.Statuses {
		if status == v {
			return true
		}
	}
	return false
}

func (v *validator) CheckMileage(mileage uint32) bool {
	return v.config.Mileage.InInterval(int32(mileage))
}

func (v *validator) CheckPrice(price uint32) bool {
	return v.config.Price.InInterval(int32(price))
}

func NewValidator(config models.InMemoryValidatorConfig) Validator {
	return &validator{config: config}
}
