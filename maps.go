// Maps are key - value store, you can store values based upton key

package main

import "fmt"

func main() {

	m := map[string]int{ //map of String type keys and int type values
		"Sri":  101,
		"Krys": 102,
	}
	fmt.Println(m)

	fmt.Println(m["Sri"])
	fmt.Println(m["Sunny"]) //value is 0 , since sunny is not available in map

	v, ok := m["Sunny"] // checking if a key exists
	fmt.Println(v)
	fmt.Println(ok)

	fmt.Println("***********checking if a key exists**********")

	if v, ok := m["Sunny"]; ok {
		fmt.Println("We are in if print condition", v)
	}

	if v, ok := m["Sri"]; ok {
		fmt.Println("We are in if print condition", v)
	}

	fmt.Println()
	fmt.Println("***********add element to map and range**********")

	fmt.Println(m)
	m["Rinku"] = 103 //adding elements to map
	m["Marley"] = 104

	for k, v := range m { //Iterating key value using range  k=key, v= value
		fmt.Println(k, v)
	}

	fmt.Println("***********Consider an example of slice**********")
	xy := []int{1, 2, 3, 4, 5, 6}
	for i, v := range xy { // It's the same thing for slice and map, here i=index, v=value
		fmt.Println(i, v)
	}

	fmt.Println()
	fmt.Println("***********delete an element from a map**********")
	fmt.Println(m)

	delete(m, "Rinku") // deleting rinku from the map
	fmt.Println(m)

	delete(m, "xyz") //trying to delete a key which is not exist
	fmt.Println(m)
}
