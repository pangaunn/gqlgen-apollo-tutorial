package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pangaunn/gqlgen-apollo-tutorial/app/handler"
)

func main() {
	port := "7777"
	g := gin.Default()

	g.GET("/pg", handler.PlaygroundHandler("/graph"))
	g.POST("/graph", handler.GraphqlHandler())
	log.Fatalln(g.Run(":" + port))
}
