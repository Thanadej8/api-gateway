// Package classification api-gateway service.
//
// the purpose of this application is learning
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Title: api-gateway
//     Host: localhost:8080
//     BasePath: /
//     Version: 0.0.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Thanadej Phadtong<authanadej@gmail.com>
// swagger:meta

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/Thanadej8/api-gateway/graph"
	"github.com/Thanadej8/api-gateway/graph/generated"
	"github.com/Thanadej8/api-gateway/modules/deals/resources"
)

func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	

	fmt.Printf("test")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}


func setupGlobalMiddleware(handler http.Handler) http.Handler {
    return uiMiddleware(handler)
}

func uiMiddleware(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Shortcut helpers for swagger-ui
        if r.URL.Path == "/swagger-ui" || r.URL.Path == "/api/help" {
            http.Redirect(w, r, "/swagger-ui/", http.StatusFound)
            return
        }
        // Serving ./swagger-ui/
        if strings.Index(r.URL.Path, "/swagger-ui/") == 0 {
            http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("swagger-ui"))).ServeHTTP(w, r)
            return
        }
        handler.ServeHTTP(w, r)
    })
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	url := ginSwagger.URL("/swagger.json") // The url pointing to API definition
	fmt.Println("url: ")
	fmt.Println(url)
	r.GET("/swagger", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.GET("/deals/:id", resources.GetDealByID)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}