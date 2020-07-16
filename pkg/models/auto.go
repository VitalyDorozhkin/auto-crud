package models

type Auto struct {
	ID      uint32 `json:"id"`
	Brand   string `json:"brand"`
	Model   string `json:"model"`
	Price   uint32 `json:"price"`
	Status  string `json:"status"`
	Mileage int32  `json:"mileage"`
}

type CreateAutoRequest struct {
	Brand   string `json:"brand"`
	Model   string `json:"model"`
	Price   uint32 `json:"price"`
	Status  string `json:"status"`
	Mileage int32  `json:"mileage"`
}

type UpdateAutoRequest struct {
	Brand   *string `json:"brand"`
	Model   *string `json:"model"`
	Price   *uint32 `json:"price"`
	Status  *string `json:"status"`
	Mileage *int32  `json:"mileage"`
}

type IDResponse struct {
	Data  *IDStruct `json:"data,omitempty"`
	Error string    `json:"error,omitempty"`
}

type StatusResponse struct {
	Data  *bool  `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

type AutoResponse struct {
	Data  *Auto  `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

type IDStruct struct {
	ID uint32 `json:"id"`
}
