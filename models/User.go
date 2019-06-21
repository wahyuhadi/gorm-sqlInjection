package models

import (
	"time"
)

// User Models
type User struct {
	ID        int64      `json:"id" gorm:"primary_key"`
	Role      string     `json:"role,omitempty" gorm:"not null"`
	Email     string     `json:"email" gorm:"not null; size:255"`
	Address   string     `json:"address" gorm:"not null; size:255"` // Default size for string is 255, reset it with this tag
	Username  string     `json:"username" gorm:"not null; size:255"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" sql:"index"`
	//Profile   *Profile   `json:"profile"` //`gorm:"foreignkey:UserID;association_foreignkey:Refer"` // One-To-Many relationship (has many - use Email's UserID as foreign key)
}

// TableName Create Return Function
func (User) TableName() string {
	return "users" // table name when succesfully migrate
}
