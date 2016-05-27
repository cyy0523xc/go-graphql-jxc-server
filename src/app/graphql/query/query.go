package query

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
	"io/ioutil"
	"os"
	"time"

	"app/database"
	"app/graphql/gtype"
)

type loginUser struct {
	ID    int32
	Token string
}

var ibbdSecretKey []byte = nil

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file!")
	}

	ibbdSecretKey, err = ioutil.ReadFile(os.Getenv("IBBD_SECRET_FILE"))
	if err != nil {
		panic("error")
	}
}

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{

		// 查询单个用户
		"user": &graphql.Field{
			Type: gtype.UserType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},

			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				idQuery, isOK := params.Args["id"].(int)
				if isOK {
					println("is OK")
					println(idQuery)
					return database.User{}, nil
				}

				println("is not OK")
				return database.User{}, nil
			},
		},

		// 用户登陆
		"login": &graphql.Field{
			Type: gtype.LoginType,
			Args: graphql.FieldConfigArgument{
				"phone": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},

			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				phoneQuery, isPhoneOK := params.Args["phone"].(string)
				passwordQuery, isPasswordOK := params.Args["password"].(string)
				println(phoneQuery)
				println(passwordQuery)

				if isPasswordOK && isPhoneOK {
					token := jwt.New(jwt.SigningMethodHS256)
					token.Claims["user"] = "1"
					token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
					tokenString, err := token.SignedString(ibbdSecretKey)
					if err == nil {
						println("is OK")
					} else {
						println("token error")
					}
					println(tokenString)
					return loginUser{ID: 1, Token: tokenString}, nil
				}

				println("is not OK")
				return loginUser{}, nil
			},
		},
	},
})
