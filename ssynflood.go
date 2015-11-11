package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	targetCIDR = kingpin.Arg("target", "Target CIDR").Required().String()
	targetPort = kingpin.Arg("port", "Target port").Default("0:65535").String()
)

func main() {
	kingpin.Parse()
}
