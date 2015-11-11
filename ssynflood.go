package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	sourceCIDR  = kingpin.Flag("source", "Source CIDR.").Short('s').Default("0.0.0.0/0").String()
	targetCIDR  = kingpin.Arg("target", "Target CIDR").Required().String()
	targetPorts = kingpin.Arg("port", "Target port range (0:65535)").Default("0:65535").String()
)

func main() {
	kingpin.Parse()
	source := NewCIDR(*sourceCIDR)
	target := NewCIDR(*targetCIDR)
	ports := NewPortRange(*targetPorts)
	worker := NewAttackWorker(*source, *target, *ports)
	worker.Start()
}
