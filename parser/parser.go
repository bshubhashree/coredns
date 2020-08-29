package parser

import (
	"context"
	"encoding/json"
	"net"

	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/plugin/pkg/log"
	"github.com/coredns/coredns/request"
	"github.com/miekg/dns"
)

const Name = "parser"

type Parser struct {
	Next plugin.Handler
}

func (g Parser) Name() string { return Name }
func (g Parser) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	log.Info("Parser entering")

	state := request.Request{W: w, Req: r}
	qName := state.QName()
	//qType := state.QType()
	sourceIP := state.IP()
	var clientSubnet net.IP
	opt := r.IsEdns0()
	for _, o := range opt.Option {
		if e, ok := o.(*dns.EDNS0_SUBNET); ok {
			// EDNS0_ECS present
			clientSubnet = e.Address
			// Turn to debug later
			log.Info("qName", qName, "sourceIP", sourceIP, "subnet", clientSubnet)
		}
	}
	msg, _ := json.Marshal(r)
	log.Info("Here is the dns msg", string(msg))
	if state.QName() == "blah.com." {
		a := new(dns.Msg)
		a.SetReply(r)
		a.Compress = true
		a.Authoritative = true

		rr := dns.TypeToRR[dns.TypeTXT]()

		switch rr := rr.(type) {
		case *dns.TXT:
			rr.Hdr = dns.RR_Header{Name: state.QName(), Rrtype: dns.TypeTXT, Class: state.QClass(), Ttl: 60}
			rr.Txt = []string{"Your query is: " + state.QName() + " " + state.Type() + " " + clientSubnet.String()}
		default:
			// do nothing
		}

		a.Answer = append(a.Answer, rr)

		log.Info("Parser returning a TXT record: \n", a)

		err := w.WriteMsg(a)
		if err != nil {
			log.Errorf("Parser failed to write %v", err)
		}

		return dns.RcodeSuccess, nil
	}

	log.Info("Parser not matched. Forward.")
	return plugin.NextOrFailure(g.Name(), g.Next, ctx, w, r)
}
