package api

import (
	"graphql-intro/connection"
	"graphql-intro/gqlargs"
	"graphql-intro/models"
	"graphql-intro/types"
	"math/rand"
	"time"

	"github.com/graphql-go/graphql"
)

// var products = []models.Product{
// 	{ID: 1,
// 		Name:  "Coki-coki",
// 		Info:  "Coklat + kacang mete",
// 		Price: 1000.00,
// 	},
// 	{
// 		ID:    2,
// 		Name:  "Indomie Goreng",
// 		Info:  "Mie goreng merk indomie",
// 		Price: 3000.00,
// 	},
// }

// var users = []models.User{
// 	{
// 		ID:   1,
// 		Name: "Agus Saputra",
// 		Age:  23,
// 	},
// 	{
// 		ID:   2,
// 		Name: "Raymond",
// 		Age:  40,
// 	},
// }

// MutationType global
var MutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"create_province": &graphql.Field{
				Type:        types.ProvinceType(),
				Description: "Create new province",
				Args:        gqlargs.CreateProvinceArgs(),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					db := connection.Connect()
					rand.Seed(time.Now().UnixNano())
					var province models.Provinces
					province.ID = uint(rand.Intn(100000))
					province.Name = params.Args["name"].(string)

					db.Create(&province)

					return province, nil
				},
			},

			"update_province": &graphql.Field{
				Type:        types.ProvinceType(),
				Description: "update province",
				Args:        gqlargs.UpdateProvinceArgs(),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					db := connection.Connect()
					id, _ := params.Args["id"].(int)
					name, _ := params.Args["name"].(string)

					province := models.Provinces{}
					db.Model(&province).Where("id = ?", id).Update("name", name)

					return province, nil
				},
			},

			"delete_province": &graphql.Field{
				Type:        types.ProvinceType(),
				Description: "delete province",
				Args:        gqlargs.DeleteProvinceArgs(),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					db := connection.Connect()
					id, _ := params.Args["id"].(int)
					var province = models.Provinces{}
					db.Delete(&province, id)

					return province, nil
				},
			},
			// * ==========================================================
		},
	},
)
