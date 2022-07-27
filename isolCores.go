package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func getIsolCPUs() []int {
	content, err := ioutil.ReadFile("/proc/cmdline")

	if err != nil {
		log.Fatal(err)
	}

	str1 := string(content)

	// rm last \n if is pressent
	if str1[len(str1)-1] == '\n' {
		str1 = str1[0 : len(str1)-1]
	}

	// strings that contains "="
	res1 := strings.SplitAfter(str1, "=")
	isolStr := res1[len(res1)-1]

	// strings that contains ","
	res2 := strings.SplitAfter(isolStr, ",")

	var isolCores []int

	for index, element := range res2 {
		if index != len(res2)-1 {
			element = element[0 : len(element)-1]
		}

		// is not a range
		if !strings.Contains(element, "-") {

			// fmt.Println(element, "at index ", index, " a single")
			y, e := strconv.Atoi(element)
			if e != nil {
				fmt.Println("Something went wrong1")
			}

			isolCores = append(isolCores, y)

			// is a range
		} else {
			coreStart, e := strconv.Atoi(strings.Split(element, "-")[0])

			if e != nil {
				fmt.Println("Something went wrong2")
			}

			coreStop, e := strconv.Atoi(strings.SplitAfter(element, "-")[1])

			if e != nil {
				fmt.Println("Something went wrong3")
			}

			for i := coreStart; i <= coreStop; i++ {
				isolCores = append(isolCores, i)
			}
		}
	}

	return isolCores

}

func main() {
	fmt.Println(getIsolCPUs())
}
