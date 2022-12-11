package domain

import "lj/app/execute"

type Domain struct {
	execute *execute.Execute
}

func NewDomain(execute *execute.Execute) *Domain {
	return &Domain{execute: execute}
}
