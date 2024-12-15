package models

// Country is a model for country
type Country struct {
	BaseModel
	Name      string `gorm:"size:15;type:string;not null"`
	Cities    []City
	Companies []Company
}
