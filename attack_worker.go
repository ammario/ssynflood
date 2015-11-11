package main

import (
//	"github.com/google/gopacket/layers"
)

type attackWorker struct {
	source CIDR
	target CIDR
	ports  portRange
}

func NewAttackWorker(source CIDR, target CIDR, ports portRange) *attackWorker {
	this := attackWorker{}
	this.source = source
	this.target = target
	this.ports = ports
	return &this
}

func (this *attackWorker) Send() {
	//var dstport layers.TCPPort
}

func (this *attackWorker) Start() {
	this.Send()
}
