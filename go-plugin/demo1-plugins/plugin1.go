package main

import (
	"fmt"
	"log"

	"github.com/MarianNoaghea/plugin-go/common"
)

func init() {
	log.Println("plugin1 init")
}

var V common.CPUSet

func F_print_V() common.CPUSet {
	fmt.Println("plugin1: cpuSet variable V= ", V)

	return common.NewCPUSet(9)
}

type foo struct{}

func (foo) M1() {
	fmt.Println("plugin1: invoke foo.M1")
}

func M() {
	fmt.Printf("Hello world, from Marian go plugin\n")
}

var Foo foo
