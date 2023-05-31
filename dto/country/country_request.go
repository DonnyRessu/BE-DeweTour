package tripdto

type CreateCountryRequest struct {
	ID   int    `json:"id" gorm:"primary_key:auto_increment"`
	Name string `json:"name" gorm:"type:varchar"`
}

type UpdateCountryRequest struct {
	ID   int    `json:"id" gorm:"primary_key: auto_increment"`
	Name string `json:"name" gorm:"type:varchar(255)"`
}