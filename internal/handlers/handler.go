package handlers

import (
	"net/http"

	"github.com/DanilKlochkov/golang-graphql/internal/models"
	"github.com/DanilKlochkov/golang-graphql/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func GraphqlHandler(r *repository.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body map[string]interface{}
		c.BindJSON(&body)
		requestString := body["query"].(string)

		rootQuery := graphql.NewObject(graphql.ObjectConfig{
			Name: "RootQuery",
			Fields: graphql.Fields{
				"note": getNote(r),
			},
		})

		rootMutation := graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"create": insertNoteMutation(r),
			},
		})

		schema, err := graphql.NewSchema(graphql.SchemaConfig{
			Query:    rootQuery,
			Mutation: rootMutation,
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		res := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: requestString,
		})

		c.JSON(http.StatusOK, res)
	}
}

func getNote(r *repository.Repository) *graphql.Field {
	return &graphql.Field{
		Type:        models.NoteType,
		Description: "get note item by ID",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(int)
			i64 := int64(id)

			model, err := r.Get(i64)
			if err != nil {
				return nil, err
			}

			return model, nil
		},
	}
}

func insertNoteMutation(r *repository.Repository) *graphql.Field {
	return &graphql.Field{
		Type:        models.NoteType,
		Description: "insert note item",
		Args: graphql.FieldConfigArgument{
			"title": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"content": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			note := models.Note{
				Title:   p.Args["title"].(string),
				Content: p.Args["content"].(string),
			}
			model, err := r.Create(note)
			if err != nil {
				return nil, err
			}
			return model, nil
		},
	}
}
