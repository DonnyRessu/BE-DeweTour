package repositories

import (
	"week2/models"

	"gorm.io/gorm"
)

type TripRepository interface {
	FindTrip() ([]models.Trip, error)
	GetTrip(ID int) (models.Trip, error)
	CreateTrip(trip models.Trip) (models.Trip, error)
	UpdateTrip(trip models.Trip) (models.Trip, error)
	DeleteTrip(trip models.Trip, ID int) (models.Trip, error)
}

func RepositoryTrip(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTrip() ([]models.Trip, error) {
	var trip []models.Trip
	err := r.db.Preload("Country").Find(&trip).Error

	return trip, err
}

func (r *repository) GetTrip(ID int) (models.Trip, error) {
	var trip models.Trip

	err := r.db.First(&trip, ID).Error

	return trip, err
}

func (r *repository) CreateTrip(trip models.Trip) (models.Trip, error) {
	err := r.db.Preload("Country").Create(&trip).Error

	return trip, err
}

func (r *repository) UpdateTrip(trip models.Trip) (models.Trip, error) {
	err := r.db.Preload("Country").Save(&trip).Error
	// err := r.db.Exec("UPDATE trips SET country_id=? WHERE id=?", trip.CountryID, trip.ID).Error

	return trip, err
}

func (r *repository) DeleteTrip(trip models.Trip, ID int) (models.Trip, error) {
	err := r.db.Delete(&trip, ID).Scan(&trip).Error

	return trip, err
}
