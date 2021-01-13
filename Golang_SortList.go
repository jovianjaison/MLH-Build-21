package main

import (
	"fmt"
	"sort"
)

func main() {
	//Using a map collection
	//Using Hackathon names as keys and Days as values

	m := map[string]int{"Local Hack Day": 27, "HackDavis": 2, "PyJac": 14}
	keys := make([]string, 0 , len(m))

	for i := range m {
  		keys = append(keys,i)
	}

	sort.Strings(keys)

	for _, i := range keys {
    		fmt.Println(i, m[i])
	}
}
