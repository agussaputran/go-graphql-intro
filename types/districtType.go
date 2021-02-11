package types

import "github.com/graphql-go/graphql"

// DistrictType type
func DistrictType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "District",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}
