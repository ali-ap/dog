package controllers

import (
	internal "dog/internal/app/common"
	"dog/internal/app/dig/dtos"
	"dog/internal/app/dig/queries"
	"github.com/gin-gonic/gin"
	"net/http"
)

// swagger:route POST /api/v1/dig/ Dig
// consumes:
//     - app/json
// produces:
//     - app/json
// parameters:
//       + name: input
//         in: body
//         required: true
//         type: DigQuery
// responses:
//  400: CommonError
//  200: DigResponse
func Dig(context *gin.Context) {

	var query queries.DigQuery
	if err := context.ShouldBindJSON(&query); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": internal.NewCommonError(internal.PARAMETER_VALIDATION_ERROR, err.Error())})
		return
	}
	res, err := query.Handle()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": internal.NewCommonError(internal.GENERIC_ERROR, err.Error())})
		return
	}
	context.JSON(http.StatusOK, dtos.MapDNSMessageToDigResponse(res))
}
