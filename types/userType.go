package types

import "github.com/graphql-go/graphql"

// UserType func
func UserType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"age": &graphql.Field{
					Type: graphql.Int,
				},
			},
		},
	)
}
