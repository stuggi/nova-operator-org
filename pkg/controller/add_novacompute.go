package controller

import (
	"github.com/nova-operator/pkg/controller/novacompute"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, novacompute.Add)
}
