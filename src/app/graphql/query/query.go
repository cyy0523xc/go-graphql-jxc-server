package query

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
	"gopkg.in/hlandau/passlib.v1"
	//"io/ioutil"
	//"os"
	"strconv"
	"time"

	"app/database"
	"app/graphql/gtype"
	"app/lib"
)

type loginUser struct {
	ID    int32
	Token string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
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

		// 查询用户列表: 分页，排序
		// 需要登陆状态才能操作
		// curl -g 'http://localhost:8080/graphql?query={userList{ID,Name}}'
		"userList": &graphql.Field{
			Type: graphql.NewList(gtype.UserType),
			Args: graphql.FieldConfigArgument{
				// 页码
				"Page": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				// 每页显示的条数，默认为20
				"Limit": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				// 排序字段
				"SortBy": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				// 排序类型
				"SortType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				// 登陆Token
				"Token": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},

			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var (
					token    string
					isOK     bool
					page     int
					limit    int
					sortBy   string
					sortType string
				)

				// check login
				token, isOK = params.Args["Token"].(string)
				if !isOK {
					return nil, errors.New("Error: param Token")
				}
				println(token)
				_, err := lib.ParseUserToken(token)
				if err != nil {
					return nil, err
				}

				page, isOK = params.Args["Page"].(int)
				if !isOK || page < 1 {
					page = 1
				}
				limit, isOK = params.Args["Limit"].(int)
				if !isOK || limit < 5 {
					limit = 20
				}
				sortBy, isOK = params.Args["SortBy"].(string)
				if !isOK {
					sortBy = "id"
				}
				sortType, isOK = params.Args["SortType"].(string)
				if !isOK || sortType != "asc" {
					sortType = "desc"
				}

				// 计算offset
				offset := (page - 1) * limit

				db := database.GetDB()
				var users []database.User
				if db.Order(sortBy + " " + sortType).Offset(offset).Limit(limit).Find(&users).RecordNotFound() {
					return nil, errors.New("Error: not found at page = " + strconv.Itoa(page))
				}

				return users, nil
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
				//println(user.Password)

				_, err := passlib.Verify(passwordQuery, user.Password)
				if err != nil {
					return nil, err
				}

				// 生成token
				expTime := time.Now().Add(time.Hour * 72).Unix()
				token, err := lib.GetUserToken(user.ID, expTime)
				if err != nil {
					return nil, err
				}
				return loginUser{ID: 1, Token: token}, nil
			},
		},
	},
})
