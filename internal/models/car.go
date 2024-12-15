package models

import "time"

type CarModel struct {
	BaseModel
	Name                 string  `gorm:"size:15;type:string;not null; unique"`
	Company              Company `gorm:"foreignKey:CompanyId;constraint:OnDelete:No Action;OnUpdate:No Action"`
	CompanyId            int
	CarType              CarType `gorm:"foreignKey:CarTypeId;constraint:OnDelete:No Action;OnUpdate:No Action"`
	CarTypeId            int
	Gearbox              Gearbox `gorm:"foreignKey:GearboxId;constraint:OnDelete:No Action;OnUpdate:No Action"`
	GearboxId            int
	CarModelColors       []CarModelColor
	CarModelPersianYears []CarModelYear
	CarModelProperties   []CarModelProperty
	CarModelImages       []CarModelImage
	CarModelComments     []CarModelComment
}

type CarModelColor struct {
	BaseModel
	CarModel   CarModel `gorm:"foreignKey:CarModelId;constraint:OnDelete:No Action;OnUpdate:No Action"`
	CarModelId int
	Color      Color `gorm:"foreignKey:ColorId;constraint:OnDelete:No Action;OnUpdate:No Action"`
	ColorId    int
}

type CarModelYear struct {
	BaseModel
	CarModel               CarModel `gorm:"foreignKey:CarModelId;constraint:OnDelete:No Action;OnUpdate:No Action"`
	CarModelId             int
	Year                   PersianYear `gorm:"foreignKey:YearId;constraint:OnDelete:No Action;OnUpdate:No Action"`
	YearId                 int
	CarModelPriceHistories []CarModelPriceHistory
}

type CarModelImage struct {
	BaseModel
	CarModel   CarModel `gorm:"foreignKey:CarModelId;constraint:OnDelete:No Action;OnUpdate:No Action"`
	CarModelId int
	Image      File `gorm:"foreignKey:ImageId;constraint:OnDelete:No Action;OnUpdate:No Action"`
	ImageId    int
}

type CarModelPriceHistory struct {
	BaseModel
	CarModelYear   CarModelYear `gorm:"foreignKey:CarModelYearId;constraint:OnDelete:No Action;OnUpdate:No Action"`
	CarModelYearId int
	Price          float64   `gorm:"type:decimal(10,2);not null"`
	PriceAt        time.Time `gorm:"TIMESTAMP with time zone;not null"`
}

type CarModelProperty struct {
	BaseModel
	CarModel   CarModel `gorm:"foreignKey:CarModelId;constraint:OnDelete:No Action;OnUpdate:No Action"`
	CarModelId int
	Property   Property `gorm:"foreignKey:PropertyId;constraint:OnDelete:No Action;OnUpdate:No Action"`
	PropertyId int
	Value      string `gorm:"size:100;type:string;not null"`
}

type CarModelComment struct {
	BaseModel
	CarModel   CarModel `gorm:"foreignKey:CarModelId;constraint:OnDelete:No Action;OnUpdate:No Action"`
	CarModelId int
	User       User `gorm:"foreignKey:UserId;constraint:OnDelete:No Action;OnUpdate:No Action"`
	UserId     int
	Message    string `gorm:"size:500;type:string;not null"`
}

type Gearbox struct {
	BaseModel
	Name      string `gorm:"size:15;type:string;not null; unique"`
	CarModels []CarModel
}

type CarType struct {
	BaseModel
	Name      string `gorm:"size:15;type:string;not null; unique"`
	CarModels []CarModel
}

type Company struct {
	BaseModel
	Name      string  `gorm:"size:15;type:string;not null; unique"`
	Country   Country `gorm:"foreignKey:CountryId;constraint:OnDelete:No Action;OnUpdate:No Action"`
	CountryId int
	CarModels []CarModel
}
