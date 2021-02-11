package main

import (
	"graphql-intro/connection"
	"graphql-intro/gqlargs"
	"graphql-intro/models"
	"graphql-intro/seeders"
	"graphql-intro/types"
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

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"product_list": &graphql.Field{
				Type:        graphql.NewList(types.ProductType()),
				Description: "Get Product List",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return products, nil
				},
			},
			"user_list": &graphql.Field{
				Type:        graphql.NewList(types.UserType()),
				Description: "Get User List",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return users, nil
				},
			},
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

var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			//* =================== PRODUCT MUTATION ===================================== //
			"create_product": &graphql.Field{
				Type:        types.ProductType(),
				Description: "Create new product",
				Args:        gqlargs.CreateProductArgs(),
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
				Type:        types.ProductType(),
				Description: "update product",
				Args:        gqlargs.UpdateProductArgs(),
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
				Type:        types.ProductType(),
				Description: "delete product",
				Args:        gqlargs.DeleteProductArgs(),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)
					var product = Product{}

					for i, v := range products {
						if int64(id) == v.ID {
							product = products[i]
							products = append(products[:i], products[i+1:]...)
							break
						}
					}

					return product, nil
				},
			},

			//* =================== END OF PRODUCT MUTATION ===================================== //

			// *  =================== USER MUTATION ===================================== //
			"create_user": &graphql.Field{
				Type:        types.UserType(),
				Description: "Create new user",
				Args:        gqlargs.DeleteUserArgs(),
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
				Type:        types.UserType(),
				Description: "update user",
				Args:        gqlargs.UpdateUserArgs(),
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
			"delete_user": &graphql.Field{
				Type:        types.UserType(),
				Description: "delete user",
				Args:        gqlargs.DeleteUserArgs(),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)
					var user = User{}

					for i, v := range users {
						if int64(id) == v.ID {
							user = users[i]
							users = append(users[:i], users[i+1:]...)
							break
						}
					}

					return user, nil
				},
			},
			// *  =================== END OF USER MUTATION ===================================== //

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
	pgDB := connection.Connect()

	models.Migrations(pgDB)
	seeders.SeedProvince(pgDB)
	seeders.SeedDistrict(pgDB)

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
		// fmt.Println(result)
		c.JSON(http.StatusOK, result)
	})
	r.Run()
}
