package api

import (
	"fmt"
	"graphql-intro/connection"
	jwtauth "graphql-intro/jwt-auth"
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
					token := params.Context.Value("token").(string)
					verifToken, err := jwtauth.VerifyToken(token)
					if err != nil {
						return nil, err
					}
					fmt.Println(verifToken["role"])
					db := connection.Connect()
					var province []models.Provinces
					db.Preload("District").Find(&province)
					return province, nil
				},
			},
		},
	},
)
