package models

// Districts model
type Districts struct {
	ID         uint   `gorm:"primarykey" json:"id"`
	Name       string `json:"name"`
	ProvinceID uint   `json:"province_id"`
}
