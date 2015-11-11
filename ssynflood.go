package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	sourceCIDR = kingpin.Flag("source", "Source CIDR.").Short('s').Default("0.0.0.0/0").String()
	targetCIDR = kingpin.Arg("target", "Target CIDR").Required().String()
	targetPort = kingpin.Arg("port", "Target port").Default("0:65535").String()
)

func main() {
	kingpin.Parse()
	source := NewCIDR(*sourceCIDR)
	target := NewCIDR(*targetCIDR)
	worker := NewAttackWorker(*source, *target)
	worker.Start()
}
