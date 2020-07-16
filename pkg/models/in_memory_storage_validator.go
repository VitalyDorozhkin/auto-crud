package models

type InMemoryValidatorConfig struct {
	BrandLength Interval
	ModelLength Interval
	Statuses    []string
	Mileage     Interval
	Price       Interval
}

type Interval struct {
	Min *int32
	Max *int32
}

func (i Interval) InInterval(item int32) bool {
	if i.Min != nil && item < *i.Min {
		return false
	}
	if i.Max != nil && item > *i.Max {
		return false
	}
	return true
}
