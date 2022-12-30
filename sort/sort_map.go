package sort

import (
	"fmt"
	"sort"
)

// SortMap sorts a map in the order of values.
func SortMap(unOrderedMap map[string]int) {

	keys := make([]string, 0, len(unOrderedMap))
	for k := range unOrderedMap {
		keys = append(keys, k)
	}

	// sorting the keys
	sort.SliceStable(keys, func(i, j int) bool {
		return unOrderedMap[keys[i]] < unOrderedMap[keys[j]]
	})

	// printing in the order of values
	for _, k := range keys {
		fmt.Println(k, unOrderedMap[k])
	}

}
