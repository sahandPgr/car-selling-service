package models

import "time"

type PersianYear struct {
	BaseModel
	PersianTitle string    `gorm:"size:10;type:string;not null; unique"`
	Year         int       `gorm:"size:int;uniqueIndex;not null"`
	StartAt      time.Time `gorm:"type:TIMESTAMP with time zone; not null; unique"`
	EndAt        time.Time `gorm:"type:TIMESTAMP with time zone; not null; unique"`
}

type Color struct {
	BaseModel
	Name           string `gorm:"size:15;type:string;not null; unique"`
	HexCode        string `gorm:"size:15;type:string;not null; unique"`
	CarModelColors []CarModelColor
}

type File struct {
	BaseModel
	Name        string `gorm:"size:100;type:string;not null"`
	Directory   string `gorm:"size:100;type:string;not null"`
	Description string `gorm:"size:500;type:string;not null"`
	MediaType   string `gorm:"size:20;type:string;not null"`
}
