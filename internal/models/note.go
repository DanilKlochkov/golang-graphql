package models

import (
	"github.com/graphql-go/graphql"
)

type Note struct {
	ID      int    `json:"id,omitempty"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var NoteType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Tutorial",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"content": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
