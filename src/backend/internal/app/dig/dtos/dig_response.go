package dtos

import (
	"github.com/miekg/dns"
)

// swagger:model CommonResponse
type DigResponse struct {

	// Message of the response
	// in: string
	Records []string `json:"records"`
	Raw     dns.Msg  `json:"raw"`
}

func MapDNSMessageToDigResponse(msg *dns.Msg) *DigResponse {
	res := DigResponse{
		Raw: *msg,
	}
	for _, record := range msg.Answer {
		res.Records = append(res.Records, record.String())
	}
	return &res
}
