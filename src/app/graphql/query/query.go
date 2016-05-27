package query

import (
	"github.com/graphql-go/graphql"

	"app/database"
	"app/graphql/gtype"
)

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
					println("is OK")
					return database.User{}, nil
				}

				println("is not OK")
				return database.User{}, nil
			},
		},
	},
})
