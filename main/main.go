package main

import (
	"fmt"
	"github.com/vijayee/IPVM"
)

func main() {
	tester := new(ipvm.Object)
	tester.Set("god save the queen")
	ipvm.Define("queen", tester)
	fmt.Println(tester)
}
