package main

import "fmt"

func main() {
	sliceExample1()
	sliceExample2()
	sliceExample3()
	mapExample1()
}

// Allocate a slice of length 3 and capacity of 5, then append item
func sliceExample1() {
	fmt.Println("*********** Slice Example 1 ***********")
	numbers1 := make([]int, 3)
	fmt.Println("numbers1 ==", numbers1)
	fmt.Println("length ==", len(numbers1))

	numbers1 = append(numbers1, 1)
	fmt.Println("after appending:", numbers1)
	fmt.Println()
}

// For loop
func sliceExample2() {
	fmt.Println("*********** Slice Example 2 ***********")
	numbers2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("numbers2 ==", numbers2)
	fmt.Println()

	// For loop with counter
	for i := 0; i < len(numbers2); i++ {
		fmt.Println("index", i, "==", numbers2[i])
	}
	fmt.Println()

	// For each loop
	for i, item := range numbers2 {
		fmt.Println("index", i, "==", item)
	}
	fmt.Println()
}

// Subslice of a slice
func sliceExample3() {
	fmt.Println("*********** Slice Example 3 ***********")
	numbers3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println("numbers3[1] ==", numbers3[1])
	fmt.Println("numbers3[1:4] ==", numbers3[1:4])
	fmt.Println()
}

// Create, read, update, and delete operations in map
func mapExample1() {
	fmt.Println("*********** Map Example 1 ***********")
	countryCapitalMap := map[string]string{
		"France": "Paris",
		"Italy":  "Rome",
		"Japan":  "Tokyo",
		"India":  "New Delhi",
	}
	fmt.Println(countryCapitalMap)
	fmt.Println()

	countryCapitalMap["United States"] = "Washington, D.C."

	for key, value := range countryCapitalMap {
		fmt.Println("Capital of", key, "is", value)
	}

	delete(countryCapitalMap, "France")

	value, exists := countryCapitalMap["France"]

	fmt.Println("\nAfter deletion")
	fmt.Println("value ==", value)
	fmt.Println("exists ==", exists)
}
