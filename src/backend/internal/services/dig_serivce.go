package services

import (
	"dog/configs"
	"github.com/lixiangzhong/dnsutil"
	"github.com/miekg/dns"
	"log"
)

type DigService struct {
	Groper *dnsutil.Dig
}

func NewDigService() *DigService {
	var dig dnsutil.Dig
	config, err := configs.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	dig.At(config.Dig.Host)
	return &DigService{Groper: &dig}
}

func (d *DigService) GetMsg(Type uint16, domain string) (*dns.Msg, error) {
	return d.Groper.GetMsg(Type, domain)
}
