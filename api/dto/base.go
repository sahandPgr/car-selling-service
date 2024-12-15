package dto

type CreateUpdateCountryRequest struct {
	Name string `json:"name" binding:"required,alpha,min=2,max=20"`
}

type CountryResponse struct {
	Id     int            `json:"id"`
	Name   string         `json:"name"`
	Cities []CityResponse `json:"cities"`
}

type UpdateCityRequest struct {
	Name string `json:"name" binding:"required,alpha,min=2,max=20"`
}

type CreateCityRequest struct {
	Name      string `json:"name" binding:"required,alpha,min=2,max=20"`
	CountryId int    `json:"countryId"`
}
type CityResponse struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Country CountryResponse `json:"country"`
}
