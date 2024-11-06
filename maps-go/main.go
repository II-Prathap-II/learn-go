package main

import (
	"fmt"
	"sort"
)

func main() {
	// creates a map of key type - string; value type string.
	// key and value need not be of the same type.
	map1 := make(map[string]string)

	// creating a new key value pair - use the same technique to modify key value pairs
	map1["key 2"] = "value 1-2"
	map1["key 1"] = "value 1-1"
	//map1["key 3"] = "value 1-3"
	map1["key 1"] = map1["key 1"] + " modified"

	// initializing map with values
	map2 := map[string]string{
		"key 1": "value 1",
		"key 2": "value 2",
	}

	// to delete a key value pair of a map
	delete(map2, "key 2")

	// loop over maps using range function
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	// maps are iterated in unspecified manner, so it is a common pattern to order the keys and get values from the map
	fmt.Println("Iterating through the map after sorting the keys of the map and storing it in a slice")

	var map1Keys []string
	for k := range map1 {
		map1Keys = append(map1Keys, k)
	}

	sort.Strings(map1Keys)

	for _, v := range map1Keys {
		fmt.Println(v, map1[v])
	}

	// to check if a key is present in the map use this technique;
	// Subscripting a map yields two values - the value of the subscripting key and a boolean which indicates if the key present in the map.

	var key3Val string
	if val, ok := map1["key 3"]; ok {
		key3Val = val // ok variable reports if the key is present in the map.
	} else {
		key3Val = "no val"
	}
	fmt.Println(key3Val)
}
