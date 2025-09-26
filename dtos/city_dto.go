package dtos

// CityResponse is used when sending city data to the client
type CityResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// CreateCityRequest is used for seeding or admin panel
type CreateCityRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
}
