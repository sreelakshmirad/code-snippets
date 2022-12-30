package main

import (
	"github.com/sreelakshmirad/code-snippets/sort"
)

func main() {

	//
	allLogs := map[string]int{
		"Err1": 5,
		"Err2": 2,
		"Err3": 1,
		"Err4": 1,
	}
	sort.SortMap(allLogs)

}
