package parser

import (
	"github.com/caddyserver/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

func init() {
	caddy.RegisterPlugin("parser", caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {
	c.Next()

	dnsserver.
		GetConfig(c).
		AddPlugin(func(next plugin.Handler) plugin.Handler {
			return Parser{Next: next}
		})

	return nil
}
