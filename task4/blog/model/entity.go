package model

import "gorm.io/gorm"

var DB *gorm.DB

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
}
type Post struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	UserID  uint
	User    User
}
type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint
	PostID  uint
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{}, &Post{}, &Comment{})
}
