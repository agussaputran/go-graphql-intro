package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

// Product model
type Product struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Info  string  `json:"info"`
	Price float64 `json:"price"`
}

// User model
type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

var products = []Product{
	{ID: 1,
		Name:  "Coki-coki",
		Info:  "Coklat + kacang mete",
		Price: 1000.00,
	},
	{
		ID:    2,
		Name:  "Indomie Goreng",
		Info:  "Mie goreng merk indomie",
		Price: 3000.00,
	},
}

var users = []User{
	{
		ID:   1,
		Name: "Agus Saputra",
		Age:  23,
	},
	{
		ID:   2,
		Name: "Raymond",
		Age:  40,
	},
}

var productType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"info": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var userType = graphql.NewObject(
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

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"product_list": &graphql.Field{
				Type:        graphql.NewList(productType),
				Description: "Get Product List",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return products, nil
				},
			},
			"user_list": &graphql.Field{
				Type:        graphql.NewList(userType),
				Description: "Get User List",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return users, nil
				},
			},
		},
	},
)

var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			//* =================== PRODUCT MUTATION ===================================== //
			"create_product": &graphql.Field{
				Type:        productType,
				Description: "Create new product",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"info": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					rand.Seed(time.Now().UnixNano())
					product := Product{
						ID:    int64(rand.Intn(100000)),
						Name:  params.Args["name"].(string),
						Info:  params.Args["info"].(string),
						Price: params.Args["price"].(float64),
					}
					products = append(products, product)
					return product, nil
				},
			},
			"update_product": &graphql.Field{
				Type:        productType,
				Description: "Update product",
				Args: graphql.FieldConfigArgument{
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
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)
					name, _ := params.Args["name"].(string)
					info, _ := params.Args["info"].(string)
					price, _ := params.Args["price"].(float64)

					var product = Product{}

					for i, v := range products {
						if int64(id) == v.ID {
							products[i].Name = name
							products[i].Info = info
							products[i].Price = price

							product = products[i]
							break
						}
					}

					return product, nil
				},
			},

			"delete_product": &graphql.Field{
				Type:        productType,
				Description: "delete product",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)

					var product = Product{}

					for i, v := range products {
						if int64(id) == v.ID {
							products = RemoveIndex(products, i)
							break
						}
					}

					return product, nil
				},
			},
			//* =================== END OF PRODUCT MUTATION ===================================== //

			// *  =================== USER MUTATION ===================================== //
			"create_user": &graphql.Field{
				Type:        userType,
				Description: "Create new user",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"age": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					rand.Seed(time.Now().UnixNano())
					user := User{
						ID:   int64(rand.Intn(100000)),
						Name: params.Args["name"].(string),
						Age:  int64(params.Args["age"].(int)),
					}
					users = append(users, user)
					return user, nil
				},
			},

			"update_user": &graphql.Field{
				Type:        userType,
				Description: "Create new product",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"age": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)
					name, _ := params.Args["name"].(string)
					age, _ := params.Args["age"].(int)

					var user = User{}

					for i, v := range users {
						if int64(id) == v.ID {
							users[i].Name = name
							users[i].Age = int64(age)

							user = users[i]
							break
						}
					}

					return user, nil
				},
			},
			// *  =================== END OF USER MUTATION ===================================== //
		},
	},
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)

// ExecuteQuery func
func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		log.Println(result.Errors)
	}
	return result
}

func main() {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {

		type Query struct {
			Query string
		}

		// buf, _ := ioutil.ReadAll(c.Request.Body)
		var query Query
		c.Bind(&query)
		// err := json.Unmarshal(buf, &query)
		// if err != nil {
		// 	log.Println("error ->", err.Error())
		// }
		// query.Query
		result := ExecuteQuery(query.Query, schema)
		fmt.Println(result)
		c.JSON(http.StatusOK, result)
	})
	r.Run()
}

// RemoveIndex func
func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
