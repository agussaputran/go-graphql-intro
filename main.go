package main

import (
	"context"
	"graphql-intro/api"
	"graphql-intro/connection"
	"graphql-intro/models"
	"graphql-intro/seeders"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    api.QueryType,
		Mutation: api.MutationType,
	},
)

func main() {
	pgDB := connection.Connect()

	models.Migrations(pgDB)
	seeders.SeedProvince(pgDB)
	seeders.SeedDistrict(pgDB)
	seeders.SeedUser(pgDB)

	routes := chi.NewRouter()
	r := registerRoutes(routes)
	log.Println("Server up and run at " + os.Getenv("HOSTNAME") + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}

// RegisterRoutes func
func registerRoutes(r *chi.Mux) *chi.Mux {
	/* GraphQL */
	graphQL := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
		// GraphiQL: true,
		Playground: true,
	})
	// r.Use(middleware.Logger)
	r.Handle("/query", headerAuthorization(graphQL))
	return r
}

// Header Authorization
func headerAuthorization(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context
		auth := r.Header.Get("Authorization")
		ctx = context.WithValue(r.Context(), "token", auth)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
