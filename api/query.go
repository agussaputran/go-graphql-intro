package api

import (
	"graphql-intro/connection"
	"graphql-intro/models"
	"graphql-intro/types"

	"github.com/graphql-go/graphql"
)

// QueryType global
var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"province_list": &graphql.Field{
				Type:        graphql.NewList(types.ProvinceType()),
				Description: "Get Province List",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					db := connection.Connect()
					var province []models.Provinces
					db.Preload("District").Find(&province)
					return province, nil
				},
			},
		},
	},
)
