package tripdto

import "week2/models"


type CreateTripRequest struct {
	Title          string         `json:"title" form:"title" validate:"required"`
	CountryID      int            `json:"country_id" form:"country_id" validate:"required"`
	Country        models.Country `json:"country" form:"country"`
	Accomodation   string         `json:"accomodation" form:"accomodation" validate:"required"`
	Transportation string         `json:"transportation" form:"transportation" validate:"required"`
	Eat            string         `json:"eat" form:"eat" validate:"required"`
	Day            int            `json:"day" form:"day" validate:"required"`
	Night          int            `json:"night" form:"night" validate:"required"`
	Datetrip       string         `json:"datetrip" form:"datetrip" validate:"required"`
	Price          int            `json:"price" form:"price" validate:"required"`
	Quota          int            `json:"quota" form:"quota" validate:"required"`
	Description    string         `json:"description" form:"description" validate:"required"`
	Image          string         `json:"image" form:"image" validate:"required"`
}

type UpdateTripRequest struct {
	Title          string         `json:"title" form:"title" validate:"required"`
	CountryID      int            `json:"country_id" form:"country_id" validate:"required"`
	Country        models.Country `json:"country" form:"country"`
	Accomodation   string         `json:"accomodation" form:"accomodation" validate:"required"`
	Transportation string         `json:"transportation" form:"transportation" validate:"required"`
	Eat            string         `json:"eat" form:"eat" validate:"required"`
	Day            int            `json:"day" form:"day" validate:"required"`
	Night          int            `json:"night" form:"night" validate:"required"`
	Datetrip       string         `json:"datetrip" form:"datetrip" validate:"required"`
	Price          int            `json:"price" form:"price" validate:"required"`
	Quota          int            `json:"quota" form:"quota" validate:"required"`
	Description    string         `json:"description" form:"description" validate:"required"`
	Image          string         `json:"image" form:"image" validate:"required"`
}
