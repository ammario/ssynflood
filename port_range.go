package main

import ()

type portRange struct {
	minPort int
	maxPort int
}

func NewPortRange(s string) *portRange {
	this := portRange{}
	this.Parse(s)
	return &this
}

func (this *portRange) Parse(s string) {

}
