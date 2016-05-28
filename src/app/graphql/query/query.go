package query

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
	"gopkg.in/hlandau/passlib.v1"
	"io/ioutil"
	"os"
	"strconv"
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
				"ID": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},

			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				idQuery, isOK := params.Args["ID"].(int)
				if !isOK {
					return nil, errors.New("Error: id param")
				}

				db := database.GetDB()
				user := new(database.User)
				if db.First(&user, idQuery).RecordNotFound() {
					return nil, errors.New("Error: not found for id = " + strconv.Itoa(idQuery))
				}

				return user, nil
			},
		},

		// 用户登陆
		// curl -g 'http://localhost:8080/graphql?query={login(Phone:"135sd7223",Password:"12321"){ID,Token}}'
		"login": &graphql.Field{
			Type: gtype.LoginType,
			Args: graphql.FieldConfigArgument{
				"Phone": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"Password": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},

			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				phoneQuery, isPhoneOK := params.Args["Phone"].(string)
				passwordQuery, isPasswordOK := params.Args["Password"].(string)

				if !isPasswordOK || !isPhoneOK {
					return nil, errors.New("Error: params")
				}

				// 判断密码是否正确
				user := new(database.User)
				db := database.GetDB()
				//if db.Select("id", "password").Where("phone = ?", phoneQuery).First(&user).RecordNotFound() {
				if db.Select([]string{"id", "password"}).Where("phone = ?", phoneQuery).First(&user).RecordNotFound() {
					return nil, errors.New("Error: not found for phone = " + phoneQuery)
				}
				println(user.Password)

				_, err := passlib.Verify(passwordQuery, user.Password)
				if err != nil {
					return nil, errors.New("Error: password verify")
				}

				// 生成token
				token := jwt.New(jwt.SigningMethodHS256)
				token.Claims["user"] = "1"
				token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
				tokenString, err := token.SignedString(ibbdSecretKey)
				if err != nil {
					return nil, errors.New("Error: create Token")
				}
				return loginUser{ID: 1, Token: tokenString}, nil
			},
		},
	},
})
