package dto

type CreateUpdateCountryRequest struct {
	Name string `json:"name" binding:"required,alpha,min=2,max=20"`
}

type CountryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
