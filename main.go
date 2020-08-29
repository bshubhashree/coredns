package main

import (
	_ "github.com/coredns/coredns/plugin/cache"
	_ "github.com/coredns/coredns/plugin/debug"
	_ "github.com/coredns/coredns/plugin/errors"
	_ "github.com/coredns/coredns/plugin/file"
	_ "github.com/coredns/coredns/plugin/forward"
	_ "github.com/coredns/coredns/plugin/health"
	_ "github.com/coredns/coredns/plugin/log"
	_ "github.com/coredns/coredns/plugin/metrics"
	_ "github.com/coredns/coredns/plugin/nsid"
	_ "github.com/coredns/coredns/plugin/pprof"

	_ "gitlab.com/sitenet/coredns/parser"

	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/coremain"
)

var directives = []string{
	"nsid",
	"prometheus",
	"pprof",
	"errors",
	"log",
	"file",
	"parser",
	"forward",
}

func init() {
	dnsserver.Directives = directives
}

func main() {
	coremain.Run()
}
