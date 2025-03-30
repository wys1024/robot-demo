package model

type User struct {
	//gorm.Model
	ID           uint   `gorm:"primarykey"`
	Username     string `gorm:"type:varchar(255);uniqueIndex"`
	PasswordHash string `gorm:"type:varchar(255)"`
}
