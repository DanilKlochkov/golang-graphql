package app

import (
	"github.com/DanilKlochkov/golang-graphql/internal/handlers"
	"github.com/DanilKlochkov/golang-graphql/internal/repository"
	"github.com/gin-gonic/gin"
)

func Run() error {
	r, err := repository.New()
	if err != nil {
		return err
	}

	s := gin.Default()

	s.POST("/graphql", handlers.GraphqlHandler(r))

	return s.Run(":8080")
}
