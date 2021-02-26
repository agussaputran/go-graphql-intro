package types

import "github.com/graphql-go/graphql"

// UserLoginType func
func UserLoginType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "UserLogin",
			Fields: graphql.Fields{
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"token": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}
