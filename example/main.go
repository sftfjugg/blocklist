package main

import (
	_ "github.com/coredns/coredns/core/plugin"
	_ "github.com/sftfjugg/blocklist"

	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/coremain"
)

func init() {
	dnsserver.Directives = append(
		[]string{"log", "blocklist"},
		dnsserver.Directives...,
	)
}

func main() {
	coremain.Run()
}
