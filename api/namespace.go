package api

import (
	vm "github.com/robertkrimen/otto"
)

type module struct {
	Dependencies Set
	Exports      interface{}
	ModuleType   string
}

//short hand for a map of methods and values
type namespace map[string]module

//return true if a method or value key is contained in the map
func (n *namespace) contains(k string) bool {
	_, ok := n[k]
	return ok

}

//add a moadule key to the  map
func (n *namespace) add(k string) {
	n[k] = make(map[string]module)
}

//set a key to a value with dependencies
func (n *namespace) set(deps []string, k string, v interface{}) {
	n[k] = &module{mapset.NewSetFromSlice(deps), typedef(v)}

}

//set a key to a value without dependencies
func (n *namespace) set(k string, v interface{}) {
	n[k] = &module{mapset.NewSet(), v, typedef(v)}
}

// define a module without dependencies
func (n *namespace) define(k string, v interface{}) {
	if !n.contains(k) {
		n.add(k)
	}
	n[k].set(k, v)
}

// define a module with dependencies
func (n *namespace) define(deps []string, k string, v interface{}) {
	if !n.contains(k) {
		n.add(k)
	}
	n[k].set(deps, k, v)
}
func (n *namespace) require(k string) module {
	return n[k]

}
