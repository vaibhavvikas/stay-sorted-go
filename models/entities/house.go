package entities

import "time"

// House represents the House table in the database
type House struct {
	Id            uint           `gorm:"column:house_id;type:int;primaryKey;autoIncrement;"`
	Pid           string         `gorm:"column:house_pid;type:varchar(255);unique;not null;"`
	UserPid       string         `gorm:"column:user_pid;type:varchar(255);not null;"`
	Name          string         `gorm:"column:name;type:varchar(255);not null;"`
	Description   string         `gorm:"column:descripion;not null;"`
	Address       string         `gorm:"column:address;not null"`
	NumBedrooms   int            `gorm:"column:num_bedrooms;not null"`
	NumBathrooms  int            `gorm:"column:num_bathrooms;not null"`
	SquareFootage int            `gorm:"column:square_footage;not null"`
	Price         int            `gorm:"column:price;not null"`
	Amenities     string         `gorm:"column:amenities;not null"`
	City          string         `gorm:"column:city;not null;"`
	State         string         `gorm:"column:state;not null;"`
	Country       string         `gorm:"column:country;not null;"`
	Pincode       string         `gorm:"column:pincode;not null;"`
	HousePictures []HousePicture `gorm:"foreignKey:HousePid;references:Pid;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt     time.Time      `gorm:"column:created_at;type:timestamp;default:current_timestamp"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:timestamp;default:current_timestamp"`
}

// HousePicture represents the House Pictures table in the database
type HousePicture struct {
	ID         uint      `gorm:"column:picture_id;type:int;primaryKey;autoIncrement;"`
	Pid        string    `gorm:"column:picture_pid;type:varchar(255);unique;not null;"`
	HousePid   string    `gorm:"column:house_pid;type:varchar(255);not null;"`
	PictureURL string    `gorm:"column:picture_url;not null"`
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamp;default:current_timestamp"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:timestamp;default:current_timestamp"`
}
