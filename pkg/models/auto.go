package models

type Auto struct {
	ID      int32  `json:"id"`
	Brand   string `json:"brand"`
	Model   string `json:"model"`
	Price   uint32 `json:"price"`
	Status  string `json:"status"`
	Mileage int32  `json:"mileage"`
}

type AutoRequest struct {
	Brand   string `json:"brand"`
	Model   string `json:"model"`
	Price   uint32 `json:"price"`
	Status  string `json:"status"`
	Mileage int32  `json:"mileage"`
}

type StatusResponse struct {
	Data  *bool   `json:"data,omitempty"`
	Error *string `json:"error,omitempty"`
}

type AutoResponse struct {
	Data  *Auto   `json:"data,omitempty"`
	Error *string `json:"error,omitempty"`
}
