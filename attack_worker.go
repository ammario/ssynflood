package main

type attackWorker struct {
	source CIDR
	target CIDR
}

func NewAttackWorker(source CIDR, target CIDR) *attackWorker {
	this := attackWorker{}
	this.source = source
	this.target = target
	return &this
}

func (this *attackWorker) Start() {

}
