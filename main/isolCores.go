package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func GetIsolCPUs(str1 string) []int {

	// rm last \n if is pressent
	if str1[len(str1)-1] == '\n' {
		str1 = str1[0 : len(str1)-1]
	}

	// split between "isolcpus=" and " "
	res1 := strings.SplitAfter(str1, "isolcpus=")
	str2 := res1[len(res1)-1]
	res2 := strings.Split(str2, " ")

	// substrings that contains ","
	res3 := strings.SplitAfter(res2[0], ",")

	var isolCores []int

	for index, element := range res3 {
		if index != len(res3)-1 {
			element = element[0 : len(element)-1]
		}

		// is not a range
		if !strings.Contains(element, "-") {
			y, e := strconv.Atoi(element)
			if e != nil {
				return []int{}
			}

			isolCores = append(isolCores, y)

			// is a range
		} else {
			coreStart, e := strconv.Atoi(strings.Split(element, "-")[0])

			if e != nil {
				return []int{}
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
	path := "/proc/cmdline"
	// path := "input2"

	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	str1 := string(content)

	fmt.Println(GetIsolCPUs(str1))
}
