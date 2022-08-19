package main

import "fmt"

func main() {
	mymap := map[string]string{"a": "1", "container.apparmor.security.beta.kubernetes.io/" + "foo": "default", "b": "2"}

	for a, b := range mymap {
		fmt.Println(a, "---", b)
	}
	fmt.Println(mymap)
}
