package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

// BaseModel is a base model for all models
type BaseModel struct {
	ID         int            `gorm:"primarykey"`
	CreatedAt  time.Time      `gorm:"type:TIMESTAMP with time zone;not null"`
	ModifiedAt *sql.NullTime  `gorm:"type:TIMESTAMP with time zone;null"`
	DeletedAt  *sql.NullTime  `gorm:"type:TIMESTAMP with time zone;null"`
	CreatedBy  int            `gorm:"not null"`
	ModifiedBy *sql.NullInt64 `gorm:"null"`
	DeletedBy  *sql.NullInt64 `gorm:"null"`
}

// BeforeCreate is a hook that is called before creating a new model
func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var userId = -1
	if value != nil {
		userId = int(value.(float64))
	}
	m.CreatedAt = time.Now().UTC()
	m.CreatedBy = userId
	return
}

// BeforeUpdate is a hook that is called before updating a model
func (m *BaseModel) BeforeUpdate(db *gorm.DB) {
	value := db.Statement.Context.Value("UserId")
	var userId = &sql.NullInt64{Valid: false}
	if value != nil {
		userId = &sql.NullInt64{Int64: int64(value.(float64)), Valid: true}
	}
	m.ModifiedAt = &sql.NullTime{Time: time.Now().UTC(), Valid: true}
	m.ModifiedBy = userId
}

// BeforeDelete is a hook that is called before deleting a model
func (m *BaseModel) BeforeDelete(db *gorm.DB) {
	value := db.Statement.Context.Value("UserId")
	var userId = &sql.NullInt64{Valid: false}
	if value != nil {
		userId = &sql.NullInt64{Int64: int64(value.(float64)), Valid: true}
	}
	m.DeletedAt = &sql.NullTime{Time: time.Now().UTC(), Valid: true}
	m.DeletedBy = userId
}
