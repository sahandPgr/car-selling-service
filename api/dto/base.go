package dto

// CreateUpdateCountryRequest is the payload for creating or updating a country
type CreateUpdateCountryRequest struct {
	Name string `json:"name" binding:"required,alpha,min=2,max=20"`
}

// CountryResponse is the response for a country
type CountryResponse struct {
	Id        int               `json:"id"`
	Name      string            `json:"name"`
	Cities    []CityResponse    `json:"cities,omitempty"`
	Companies []CompanyResponse `json:"companies,omitempty"`
}

// UpdateCityRequest is the payload for updating a city
type UpdateCityRequest struct {
	Name string `json:"name" binding:"required,alpha,min=2,max=20"`
}

// CreateCityRequest is the payload for creating a city
type CreateCityRequest struct {
	Name      string `json:"name" binding:"required,alpha,min=2,max=20"`
	CountryId int    `json:"countryId"`
}

// CityResponse is the response for a city
type CityResponse struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Country CountryResponse `json:"country"`
}

// CreateCompanyRequest is the payload for creating a company
type CreateCompanyRequest struct {
	CountryId int    `json:"countryId" binding:"required"`
	Name      string `json:"company" binding:"required,alpha,max=15"`
}

// UpdateCompanyRequest is the payload for updating a company
type UpdateCompanyRequest struct {
	CountryId int    `json:"countryId" binding:"required"`
	Name      string `json:"company" binding:"required,alpha,max=15"`
}

// CompanyResponse is the response for a company
type CompanyResponse struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Country CountryResponse `json:"country,omitempty"`
}
