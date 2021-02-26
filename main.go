package main

import (
	"graphql-intro/api"
	"graphql-intro/connection"
	"graphql-intro/models"
	"graphql-intro/seeders"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    api.QueryType,
		Mutation: api.MutationType,
	},
)

// ExecuteQuery func
func executeQuery(query string, schema graphql.Schema) *graphql.Result {
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
	seeders.SeedUser(pgDB)

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
		result := executeQuery(query.Query, schema)
		// fmt.Println(result)
		c.JSON(http.StatusOK, result)
	})
	r.Run()
}
