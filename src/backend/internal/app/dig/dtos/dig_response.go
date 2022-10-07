package dtos

import (
	"github.com/miekg/dns"
)

// swagger:model CommonResponse
type DigResponse struct {

	// Message of the response
	// in: string
	Records []string `json:"message"`
}

func MapDNSMessageToDigResponse(msg *dns.Msg) *DigResponse {
	res := DigResponse{}

	for _, record := range msg.Answer {
		res.Records = append(res.Records, record.String())
	}
	return &res
}
