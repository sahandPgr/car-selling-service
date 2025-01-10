package dto

import "time"

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

// Struct to create a Persian year.
type CreatePersianYearRequest struct {
	PersianTitle string    `json:"persianYear" binding:"min=4,max=4"`
	Year         int       `json:"year"`
	StartAt      time.Time `json:"startAt"`
	EndAt        time.Time `json:"endAt"`
}

// Struct to update a Persian year (optional fields).
type UpdatePersianYearRequest struct {
	PersianTitle string    `json:"persianYear,omitempty" binding:"min=4,max=4"`
	Year         int       `json:"year,omitempty"`
	StartAt      time.Time `json:"startAt,omitempty"`
	EndAt        time.Time `json:"endAt,omitempty"`
}

// Struct to return Persian year details.
type PersianYearResponse struct {
	Id      int       `json:"id"`
	Year    int       `json:"year,omitempty"`
	StartAt time.Time `json:"startAt,omitempty"`
	EndAt   time.Time `json:"endAt,omitempty"`
}

// Struct to return Persian year without dates.
type PersianYearWithoutDateResponse struct {
	Id   int `json:"id"`
	Year int `json:"year,omitempty"`
}
