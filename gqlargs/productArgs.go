package gqlargs

import "github.com/graphql-go/graphql"

// CreateProductArgs args
func CreateProductArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"info": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"price": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Float),
		},
	}
}

// UpdateProductArgs args
func UpdateProductArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"info": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"price": &graphql.ArgumentConfig{
			Type: graphql.Float,
		},
	}
}

// DeleteProductArgs args
func DeleteProductArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}
}
