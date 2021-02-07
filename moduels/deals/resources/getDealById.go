package resources

import "github.com/gin-gonic/gin"

// swagger:model Deal
type Deal struct {
	id int
	name string
}

// swagger:parameters getDealByID
type DealByIDRequest struct {
	// in: path
	id int
}

// swagger:response dealResponse
type DealByIDResponse struct {
	// in: body
	deal Deal
}

// swagger:route GET /deals/{id} wognnai-deal getDealByID
//
// Gets deal by id.
//
// Responses:
//    default: genericError
//        200: dealResponse
func GetDealByID(context *gin.Context) {

	context.JSON(200, gin.H{
		"message": "pong",
	})
}