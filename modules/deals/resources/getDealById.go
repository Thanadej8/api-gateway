package resources

import (
	"context"
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/executor"

	"github.com/Thanadej8/api-gateway/graph"
	"github.com/Thanadej8/api-gateway/graph/generated"
	"github.com/gin-gonic/gin"
)

// swagger:model Deal
type Deal struct {
	// the id of deal
	//
	ID int `json:"id"`
	// name of deal
	//
	Name string `json:"name"`
}

// swagger:parameters getDealByID
type DealByIDRequest struct {

	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:response dealResponse
type DealByIDResponse struct {
	// in: body
	deal Deal
}

// swagger:route GET /deals/{id} deal getDealByID
//
// Gets deal by id.
//
// Responses:
//        200: dealResponse
func GetDealByID(ginContext *gin.Context) {
	fmt.Println("1")
	context, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	context = graphql.StartOperationTrace(context)

	graphQLParam := new(graphql.RawParams)
	fmt.Println("2")
	graphQLParam.Query = `query {
		users {
		  id
		  name
		}
	  }`
	fmt.Println("graphQLParam.Query: " + graphQLParam.Query)
	executorGraphql := executor.New(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	fmt.Println("3")
	response, err :=  executorGraphql.CreateOperationContext(context, graphQLParam)
	fmt.Println("4")
	fmt.Println(`error:`)
	fmt.Println(err)
	fmt.Println("response: ")
	fmt.Println(response)
	if err != nil {
		
	}

	responses, ctx := executorGraphql.DispatchOperation(context, response)

	ginContext.JSON(200, responses(ctx).Data,)
}