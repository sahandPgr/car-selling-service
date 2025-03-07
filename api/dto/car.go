package dto

// CreateCarTypeRequest is the payload for creating a car type
type CreateCarTypeRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=15"`
}

// UpdateCarTypeRequest is the payload for updating a car type
type UpdateCarTypeRequest struct {
	Name string `json:"name,omitempty" binding:"alpha,min=3,max=15"`
}

// CarTypeResponse is the response for a car type
type CarTypeResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// CreateGearboxRequest is the payload for creating a gearbox
type CreateGearboxRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=15"`
}

// UpdateGearboxRequest is the payload for updating a gearbox
type UpdateGearboxRequest struct {
	Name string `json:"name,omitempty" binding:"alpha,min=3,max=15"`
}

// GearboxResponse is the response for a gearbox
type GearboxResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// CreateCarModelRequest is the payload for creating a car model
type CreateCarModelRequest struct {
	Name      string `json:"name" binding:"required,min=3,max=15"`
	CompanyId int    `json:"companyId" binding:"required"`
	CarTypeId int    `json:"carTypeId" binding:"required"`
	GearboxId int    `json:"gearboxId" binding:"required"`
}

// UpdateCarModelRequest is the payload for updating a car model
type UpdateCarModelRequest struct {
	Name      string `json:"name,omitempty"`
	CompanyId int    `json:"companyId,omitempty"`
	CarTypeId int    `json:"carTypeId,omitempty"`
	GearboxId int    `json:"gearboxId,omitempty"`
}

// CarModelResponse is the response for a car model
type CarModelResponse struct {
	Id             int                     `json:"id"`
	Name           string                  `json:"name"`
	Company        CompanyResponse         `json:"company"`
	CarType        CarTypeResponse         `json:"carType"`
	Gearbox        GearboxResponse         `json:"gearbox"`
	CarModelColors []CarModelColorResponse `json:"carModelColors,omitempty"`
}

// CreateCarModelColorRequest for creating a car model color.
type CreateCarModelColorRequest struct {
	ColorId    int `json:"colorId" binding:"required"`
	CarModelId int `json:"carModelId" binding:"required"`
}

// UpdateCarModelColorRequest for updating a car model color.
type UpdateCarModelColorRequest struct {
	ColorId    int `json:"colorId,omitempty"`
	CarModelId int `json:"carModelId,omitempty"`
}

// CarModelColorResponse for car model color details.
type CarModelColorResponse struct {
	Id    int           `json:"id"`
	Color ColorResponse `json:"color"`
}

// Struct to create a car model year relationship.
type CreateCarModelYearRequest struct {
	CarModelId int `json:"carModelId" binding:"required"`
	YearId     int `json:"yearId" binding:"required"`
}

// Struct to update a car model year relationship (optional fields).
type UpdateCarModelYearRequest struct {
	CarModelId int `json:"carModelId,omitempty"`
	YearId     int `json:"yearId,omitempty"`
}

// Struct to return car model year details with Persian year information.
type CarModelYearResponse struct {
	Id                  int                            `json:"id"`
	PersianYearResponse PersianYearWithoutDateResponse `json:"persianYear"`
}
