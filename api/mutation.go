package api

import (
	"fmt"
	"graphql-intro/connection"
	"graphql-intro/gqlargs"
	jwtauth "graphql-intro/jwt-auth"
	"graphql-intro/models"
	"graphql-intro/types"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
)

// MutationType global
var MutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"create_province": &graphql.Field{
				Type:        types.ProvinceType(),
				Description: "Create new province",
				Args:        gqlargs.CreateProvinceArgs(),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					db := *connection.GetConnection()

					token := params.Context.Value("token").(string)
					verifToken, err := jwtauth.VerifyToken(token)
					if err != nil {
						return nil, err
					}
					fmt.Println(verifToken["role"])
					if verifToken["role"] == "guest" {
						return nil, err
					}

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
					db := *connection.GetConnection()

					token := params.Context.Value("token").(string)
					verifToken, err := jwtauth.VerifyToken(token)
					if err != nil {
						return nil, err
					}
					fmt.Println(verifToken["role"])
					if verifToken["role"] == "guest" {
						return nil, err
					}

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
					db := *connection.GetConnection()

					token := params.Context.Value("token").(string)
					verifToken, err := jwtauth.VerifyToken(token)
					if err != nil {
						return nil, err
					}
					fmt.Println(verifToken["role"])
					if verifToken["role"] != "admin" {
						return nil, err
					}

					id, _ := params.Args["id"].(int)
					var province = models.Provinces{}
					db.Delete(&province, id)

					return province, nil
				},
			},
			// * ==========================================================
			"login": &graphql.Field{
				Type:        types.UserLoginType(),
				Description: "login",
				Args:        gqlargs.LoginArgs(),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					var (
						db     = *connection.GetConnection()
						user   models.User
						result interface{}
					)

					email, _ := params.Args["email"].(string)
					password, _ := params.Args["password"].(string)

					db.Where("email = ?", email).First(&user)

					if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
						log.Println("Email ", user.Email, " Password salah")
						result = map[string]interface{}{
							"message": "email atau password salah",
						}
					} else {
						type authCustomClaims struct {
							Email string `json:"email"`
							Role  string `json:"role"`
							jwt.StandardClaims
						}

						claims := &authCustomClaims{
							user.Email,
							user.Role,
							jwt.StandardClaims{
								ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
								IssuedAt:  time.Now().Unix(),
							},
						}
						sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
						token, err := sign.SignedString([]byte(os.Getenv("JWT_SECRET")))
						if err != nil {
							log.Println("Gagal create token, message ", err.Error())
							result = map[string]interface{}{
								"token": nil,
							}
						} else {
							log.Println("Email ", user.Email, " Berhasil login")
							result = map[string]interface{}{
								"email": user.Email,
								"token": token,
							}
						}
					}
					return result, nil
				},
			},
		},
	},
)
