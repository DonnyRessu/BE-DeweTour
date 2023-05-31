package repositories

import (
	"week2/models"

	"gorm.io/gorm"
)

type CountryRepository interface {
	FindCountry() ([]models.Country, error)
	CreateCountry(country models.Country) (models.Country, error)
	GetCountry(Id int) (models.Country, error)
	UpdateCountry(country models.Country, Id int) (models.Country, error)
	DeleteCountry(country models.Country, Id int) (models.Country, error)
}

func RepositoryCountry(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCountry() ([]models.Country, error) {
	var country []models.Country
	err := r.db.Find(&country).Error //using find method
	return country, err
}

func (r *repository) GetCountry(Id int) (models.Country, error) {
	var country models.Country
	err := r.db.First(&country, Id).Error
	return country, err
}

func (r *repository) CreateCountry(country models.Country) (models.Country, error) {
	err := r.db.Create(&country).Error
	return country, err
}

func (r *repository) UpdateCountry(country models.Country, Id int) (models.Country, error) {
	err := r.db.Save(&country).Error
	return country, err
}

func (r *repository) DeleteCountry(country models.Country, Id int) (models.Country, error) {
	err := r.db.Delete(&country).Error
	return country, err
}
