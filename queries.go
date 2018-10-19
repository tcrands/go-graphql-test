package main

import (
	"github.com/graphql-go/graphql"
)

var queries = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"songs": songsQuery(),
		"album": albumQuery(),
	},
})

func songsQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(songType),
		Args: graphql.FieldConfigArgument{
			"album": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: songResolver,
	}
}

func albumQuery() *graphql.Field {
	return &graphql.Field{
		Type: albumType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: albumResolver,
	}
}

// SONGS: curl -g 'http://localhost:12345/graphql?query={songs(album:"ts-fearless"){title,duration}}'
// ALBUMS: curl -g 'http://localhost:12345/graphql?query={album(id:"0"){title}}'
