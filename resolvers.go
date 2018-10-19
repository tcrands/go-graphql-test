package main

import (
	"strings"

	"github.com/graphql-go/graphql"
)

func songResolver(params graphql.ResolveParams) (interface{}, error) {
	a := params.Args["album"]
	if a != nil {
		return Filter(songs, func(song Song) bool {
			return strings.Contains(song.Album, a.(string))
		}), nil
	}
	return songs, nil
}

func albumResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	if id != nil {
		for _, album := range albums {
			if album.ID == id.(string) {
				return album, nil
			}
		}
	}
	return albums, nil
}
