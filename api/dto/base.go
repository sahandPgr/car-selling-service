package dto

type CreateUpdateCountryRequest struct {
	Name string `json:"name" binding:"required,alpha,min=2,max=20"`
}

type CountryResponse struct {
	Id        int               `json:"id"`
	Name      string            `json:"name"`
	Cities    []CityResponse    `json:"cities,omitempty"`
	Companies []CompanyResponse `json:"companies,omitempty"`
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

type CreateCompanyRequest struct {
	CountryId int    `json:"countryId" binding:"required"`
	Name      string `json:"company" binding:"required,alpha,max=15"`
}

type UpdateCompanyRequest struct {
	CountryId int    `json:"countryId" binding:"required"`
	Name      string `json:"company" binding:"required,alpha,max=15"`
}

type CompanyResponse struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Country CountryResponse `json:"country,omitempty"`
}
