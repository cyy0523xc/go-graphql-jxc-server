// type是关键词，所以这里包名定义成了gtype
package gtype

import (
	"github.com/graphql-go/graphql"
)

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type: graphql.Int,
		},
		"Name": &graphql.Field{
			Type: graphql.String,
		},
		"Phone": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// 登陆类型
var LoginType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Login",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type: graphql.Int,
		},
		"Token": &graphql.Field{
			Type: graphql.String,
		},
	},
})
