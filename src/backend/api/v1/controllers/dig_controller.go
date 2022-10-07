package controllers

import (
	internal "dog/internal/app/common"
	"dog/internal/app/dig/dtos"
	"dog/internal/app/dig/queries"
	"github.com/gin-gonic/gin"
	"net/http"
)

// swagger:route POST /api/v1/notification/topic/ Notification SendPushNotificationToTopic
// send notification to all subscribed devices to a specific topic
//
// consumes:
//     - app/json
// produces:
//     - app/json
// parameters:
//       + name: input
//         in: body
//         required: true
//         type: SendPushNotificationToTopicCommand
// responses:
//  400: CommonError
//  200: CommonResponse
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
