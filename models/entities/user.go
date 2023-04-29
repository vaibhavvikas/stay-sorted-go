package entities

import "time"

type User struct {
	Id        uint      `gorm:"column:user_id;type:int;primaryKey;autoIncrement;"`
	Pid       string    `gorm:"column:user_pid;type:varchar(255);unique;not null;"`
	Name      string    `gorm:"column:name;type:varchar(255);not null;"`
	Gender    string    `gorm:"column:gender;type:varchar(255);not null;"`
	Email     string    `gorm:"column:email;type:varchar(255);unique;not null;"`
	Dob       string    `gorm:"column:dob;type:varchar(10);not null;"`
	Phone     string    `gorm:"column:phone;type:varchar(10);unique;not null;"`
	Rating    float64   `gorm:"column:rating;type:float;"`
	IsActive  bool      `gorm:"column:is_active;type:bool;"`
	Salt      string    `gorm:"column:salt;not null;"`
	Password  string    `gorm:"column:password;not null;"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:current_timestamp"`
}
