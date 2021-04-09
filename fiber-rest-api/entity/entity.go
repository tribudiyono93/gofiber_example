package entity

import "time"

type Base struct {
	CreatedBy string    `json:"created_by" gorm:"size:255"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedBy string    `json:"updated_by" gorm:"size:255"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID       string `json:"id" gorm:"primaryKey;size:255"`
	Email    string `json:"email" gorm:"unique;size:255"`
	Password string `json:"-" gorm:"size:255"`
	Name     string `json:"name" gorm:"size:255"`
	Base
}

type UserModuleRole struct {
	Email  string `json:"email" gorm:"primaryKey;size:255"`
	Module string `json:"module" gorm:"primaryKey;size:255"`
	Role   string `json:"role" gorm:"size:255"`
	Base
}

type Book struct {
	ID     uint   `gorm:"primaryKey;size:255"`
	Title  string `json:"name" gorm:"size:255"`
	Author string `json:"author" gorm:"size:255"`
	Rating int    `json:"rating" gorm:"size:255"`
	Base
}
