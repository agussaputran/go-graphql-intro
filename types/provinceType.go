package types

import "github.com/graphql-go/graphql"

// ProvinceType type
func ProvinceType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Province",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"districts": &graphql.Field{
					Type: graphql.NewList(DistrictType()),
				},
			},
		},
	)
}
