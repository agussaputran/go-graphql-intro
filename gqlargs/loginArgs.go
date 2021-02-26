package gqlargs

import "github.com/graphql-go/graphql"

// LoginArgs mutation args
func LoginArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	}
}
