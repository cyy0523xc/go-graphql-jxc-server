package mutation

import (
	"errors"
	"github.com/graphql-go/graphql"
	"gopkg.in/hlandau/passlib.v1"

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
		// curl -g 'http://localhost:8080/graphql?query=mutation+_{createUser(Name:"alex",Phone:"135sd7223",Password:"12321",Password2:"12321"){ID}}'
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
				"Password2": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},

			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var isOK bool
				user := new(database.User)

				user.Name, isOK = params.Args["Name"].(string)
				if !isOK {
					return nil, errors.New("Error: Name param")
				}
				user.Phone, isOK = params.Args["Phone"].(string)
				if !isOK {
					return nil, errors.New("Error: Phone param")
				}
				user.Password, isOK = params.Args["Password"].(string)
				if !isOK {
					return nil, errors.New("Error: Password param")
				}
				var password2 string
				password2, isOK = params.Args["Password2"].(string)
				if !isOK {
					return nil, errors.New("Error: Password2 param")
				}
				if password2 != user.Password {
					return nil, errors.New("Error: Password2 and Password not equal")
				}
				hash, err := passlib.Hash(user.Password)
				if err != nil {
					return nil, err
				}
				user.Password = hash

				if len(database.GetDB().Create(&user).GetErrors()) > 0 {
					return nil, errors.New("Error: db.create")
				}
				return user, nil
			},
		},
	},
})
