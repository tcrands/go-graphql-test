package main

import (
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
)

func main() {
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: queries,
	})

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: r.URL.Query().Get("query"),
		})
		json.NewEncoder(w).Encode(result)
	})
	http.ListenAndServe(":12345", nil)
}
