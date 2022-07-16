package main

// aggregate types (pointer, slices,maps)

func main() {

	intMap := make(map[string]int)

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	for key, value := range intMap {
		println(key, value)
	}

	delete(intMap, "three")

	for key, value := range intMap {
		println(key, value)
	}

	element, ok := intMap["four"] // element is zero value if key is not found in map and ok is false if key is not found in map and ok is true if key is found in map
	if ok {
		println(element)
	} else {
		println("key", element, "not found")
	}

	intMap["six"] = 6

	for key, value := range intMap {
		println(key, value)
	}
}
