package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"

	"app/graphql/mutation"
	"app/graphql/query"
)

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    query.RootQuery,
	Mutation: mutation.RootMutation,
})

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query()["query"][0], schema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Now server is running on port 8080")
	fmt.Println("Get single user: curl -g 'http://localhost:8080/graphql?query={user(id:1){id,name}}'")
	fmt.Println("Login user: curl -g 'http://localhost:8080/graphql?query={login(phone:\"13701370137\",password:\"123456\"){ID,Token}}'")
	fmt.Println("Create new user: curl -g 'http://localhost:8080/graphql?query=mutation+_{createUser(Name:\"alex\",Phone:\"1371300\",Password:\"1234\"){ID}}'")
	fmt.Println("Load todo list: curl -g 'http://localhost:8080/graphql?query={todoList{id,text,done}}'")
	http.ListenAndServe(":8080", nil)
}
