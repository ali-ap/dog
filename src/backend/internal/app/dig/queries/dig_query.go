package queries

import (
	"dog/internal/services"
	"github.com/go-playground/validator/v10"
	"github.com/miekg/dns"
)

// swagger:model SendPushNotificationToTokenCommand
type DigQuery struct {

	// title of the notification
	// in: body
	// required: true
	Type uint16 `json:"type" validate:"required,gte=0,lte=5258"`

	// body of the notification
	// in: body
	// required: true
	Domain string `json:"domain" validate:"required,hostname"`
}

func (q *DigQuery) Handle() (*dns.Msg, error) {
	err := q.Validate()
	if err != nil {
		return nil, err
	}
	service := services.NewDigService()
	return service.GetMsg(q.Type, q.Domain)
}
func (q *DigQuery) Validate() error {
	validate := validator.New()
	err := validate.Struct(q)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return validationErrors
	}

	return nil
}
