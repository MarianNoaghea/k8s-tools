package main

import (
	"errors"
	"fmt"
	"log"
	"plugin"

	"github.com/MarianNoaghea/plugin-go/common"
)

func init() {
	log.Println("pkg1 init")
}

type MyInterface interface {
	M1()
}

func changeSet(comm.NewCPUSetInt64)

func DoMagicWithPlugin(pluginPath string) error {
	p, err := plugin.Open(pluginPath)
	if err != nil {
		return err
	}

	v, err := p.Lookup("V")
	if err != nil {
		return err
	}

	*v.(*common.CPUSet) = common.NewCPUSetInt64(1, 2, 3)
	fmt.Println("main program: we have", *v.(*common.CPUSet))

	f, err := p.Lookup("F_print_V")
	if err != nil {
		return err
	}

	// call the function from plugin
	*v.(*common.CPUSet) = f.(func() common.CPUSet)()

	fmt.Println("main program: we have", *v.(*common.CPUSet))

	f1, err := p.Lookup("Foo")
	if err != nil {
		return err
	}
	i, ok := f1.(MyInterface)
	if !ok {
		return errors.New("f1 does not implement MyInterface")
	}
	i.M1()

	fM, err := p.Lookup("M")
	if err != nil {
		return err
	}
	fM.(func())()

	return nil
}

func main() {
	err := DoMagicWithPlugin("../demo1-plugins/plugin1.so")
	if err != nil {
		fmt.Println("LoadAndInvokeSomethingFromPlugin error:", err)
		return
	}

}

/* de scris un plugin care afiseaza "hello world" pt of functie myfunc, si sa apelez din kubelet myfunc pt a afisa "hello world" */
