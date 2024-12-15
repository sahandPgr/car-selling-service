package models

// City is a model for city
type City struct {
	BaseModel
	Name      string  `gorm:"size:15;type:string;not null"`
	Country   Country `gorm:"foreignKey:CountryId;constraint:OnDelete:No Action;OnUpdate:No Action"`
	CountryId int
}
