package models

// City is a model for city
type City struct {
	BaseModel
	Name      string `gorm:"size:15;type:string;not null"`
	CountryId uint
	Country   Country `gorm:"foreignKey:CountryId"`
}
