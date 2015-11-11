package main

import (
	"log"
	"net"
)

type CIDR struct {
	ip  net.IP
	net *net.IPNet
}

func NewCIDR(s string) *CIDR {
	this := CIDR{}
	this.Parse(s)
	return &this
}

func (this *CIDR) Parse(s string) {
	ip, net, err := net.ParseCIDR(s)
	if err != nil {
		log.Fatal(err)
	}
	this.ip = ip
	this.net = net
}
