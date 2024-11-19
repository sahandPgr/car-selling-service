package models

// User struct
type User struct {
	BaseModel
	Username    string `gorm:"type:string;size:12;unique;not null"`
	FirstName   string `gorm:"type:string;size:12;null"`
	LastName    string `gorm:"type:string;size:12;null"`
	Email       string `gorm:"type:string;size:64;null;unique;default:null"`
	PhoneMobile string `gorm:"type:string;size:11;null;unique;default:null"`
	Password    string `gorm:"type:string;size:64;not null"`
	Available   bool   `gorm:"default:true"`
	UserRoles   *[]UserRole
}

// Role struct
type Role struct {
	BaseModel
	Name      string `gorm:"type:string;size:12;unique;not null"`
	UserRoles *[]UserRole
}

// UserRole struct
type UserRole struct {
	BaseModel
	User   User `gorm:"foreignKey:UserId;constraint:OnDelete:No Action;OnUpdate:No Action"`
	Role   Role `gorm:"foreignKey:RoleId;constraint:OnDelete:No Action;OnUpdate:No Action"`
	UserId int
	RoleId int
}
