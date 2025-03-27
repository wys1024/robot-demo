package model

import "time"

type Aichats struct {
	Id        int       `gorm:"primaryKey"`
	Model     string    `gorm:"type:varchar(255)"`
	Questions string    `gorm:"type:text"`
	Answers   string    `gorm:"type:text"`
	Time      time.Time `gorm:"type:time"`
}
