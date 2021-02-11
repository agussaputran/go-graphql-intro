package seeders

import (
	"encoding/json"
	"fmt"
	"graphql-intro/models"
	"graphql-intro/service"
	"log"
	"strconv"

	"gorm.io/gorm"
)

// SeedDistrict func
func SeedDistrict(db *gorm.DB) {
	resBody := service.FetchFromRajaongkir("/city")
	var (
		response models.RajaongkirCity
		dist     models.Districts
	)

	if err := json.Unmarshal(resBody, &response); err != nil {
		log.Fatalln("Error -> ", err.Error())
	}

	for i := 0; i < 100; i++ {
		uProvID, _ := strconv.ParseUint(response.RajaOngkir.CityResults[i].ProvinceID, 10, 32)
		dist.Name = response.RajaOngkir.CityResults[i].CityName
		dist.ProvinceID = uint(uProvID)
		dist.ID = 0
		db.Create(&dist)
	}
	fmt.Println("Seed District created")
}
