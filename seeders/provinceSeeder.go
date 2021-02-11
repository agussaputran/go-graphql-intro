package seeders

import (
	"encoding/json"
	"fmt"
	"graphql-intro/models"
	"graphql-intro/service"
	"log"

	"gorm.io/gorm"
)

// SeedProvince func
func SeedProvince(db *gorm.DB) {
	resBody := service.FetchFromRajaongkir("/province")
	var (
		response models.RajaongkirProvince
		prov     models.Provinces
	)

	if err := json.Unmarshal(resBody, &response); err != nil {
		log.Fatalln("Error -> ", err.Error())
	}

	for i := 0; i < len(response.RajaOngkir.ProvinceResults); i++ {
		prov.Name = response.RajaOngkir.ProvinceResults[i].Province
		prov.ID = 0
		db.Create(&prov)
	}
	fmt.Println("Seed Province created")
}
