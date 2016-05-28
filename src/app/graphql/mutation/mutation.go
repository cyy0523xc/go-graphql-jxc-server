package mutation

import (
	"errors"
	"github.com/graphql-go/graphql"

	"app/database"
	"app/graphql/gtype"
)

type createUserParams struct {
	Name     string
	Phone    string
	Password string
}

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createUser": &graphql.Field{
			Type: gtype.UserType,
			Args: graphql.FieldConfigArgument{
				"Name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"Phone": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"Password": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},

			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var isOK bool
				user := new(database.User)

				user.Name, isOK = params.Args["Name"].(string)
				if !isOK {
					return database.User{}, errors.New("Error: Name param")
				}
				user.Phone, isOK = params.Args["Phone"].(string)
				if !isOK {
					return database.User{}, errors.New("Error: Phone param")
				}
				user.Password, isOK = params.Args["Password"].(string)
				if !isOK {
					return database.User{}, errors.New("Error: Password param")
				}

				database.GetDB().Create(&user)
				println(user.ID)
				return user, nil
			},
		},
	},
})
