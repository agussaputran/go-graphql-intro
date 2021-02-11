package models

// Provinces model
type Provinces struct {
	ID       uint        `gorm:"primarykey" json:"id"`
	Name     string      `json:"name"`
	District []Districts `gorm:"ForeignKey:ProvinceID" json:"districts"`
}
