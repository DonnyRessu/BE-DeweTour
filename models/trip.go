package models

import "time"

type Trip struct {
	ID             int       `json:"id" gorm:"primary_key: auto_increment"`
	Title          string    `json:"title" binding:"required, title" gorm:"unique; not null"`
	CountryID      int       `json:"country_id"`
	Country        Country   `json:"country"`
	Accomodation   string    `json:"accomodation" gorm:"type VARCHAR(255)"`
	Transportation string    `json:"transportation" gorm:"type VARCHAR(255)"`
	Eat            string    `json:"eat" gorm:"type VARCHAR(255)"`
	Day            int       `json:"day"`
	Night          int       `json:"night"`
	DateTrip       string    `json:"datetrip" gorm:"type VARCHAR(255)"`
	Price          int       `json:"price"`
	Quota          int       `json:"quota"`
	Description    string    `json:"description" gorm:"type VARCHAR(255)"`
	Image          string    `json:"image" gorm:"type VARCHAR(255)"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
}

type TripResponse struct {
	ID             int     `json:"id"`
	Title          string  `json:"title"`
	CountryID      int     `json:"country_id"`
	Country        Country `json:"country"`
	Accomodation   string  `json:"accomodation"`
	Transportation string  `json:"transportation"`
	Eat            string  `json:"eat"`
	Day            int     `json:"day"`
	Night          int     `json:"night"`
	DateTrip       string  `json:"datetrip"`
	Price          int     `json:"price"`
	Quota          int     `json:"quota"`
	Description    string  `json:"description"`
	Image          string  `json:"image"`
}

func (TripResponse) TableName() string {
	return "Trip"
}
