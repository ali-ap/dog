package services

import (
	"dog/configs"
	"github.com/lixiangzhong/dnsutil"
	"github.com/miekg/dns"
)

type DigService struct {
	Groper *dnsutil.Dig
}

func NewDigService() *DigService {
	var dig dnsutil.Dig
	dig.At(configs.AppConfig.Dig.Host)
	return &DigService{Groper: &dig}
}

func (d *DigService) GetMsg(Type uint16, domain string) (*dns.Msg, error) {
	return d.Groper.GetMsg(Type, domain)
}
