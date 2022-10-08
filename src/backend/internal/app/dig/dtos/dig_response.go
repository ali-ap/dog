package dtos

import (
	"github.com/miekg/dns"
)

// swagger:model DigResponse
type DigResponse struct {

	// DNS records
	// in: []string
	Records []string `json:"records"`

	// raw query response
	// in: dns.Msg
	Raw dns.Msg `json:"raw"`
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
