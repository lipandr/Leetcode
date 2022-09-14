package main

import (
	"sort"
	"strconv"
)

func FizzBuzz(n int) []string {
	var result []string
	for i := 1; i <= n; i++ {
		div3 := i%3 == 0
		div5 := i%5 == 0
		if div3 && div5 {
			result = append(result, "FizzBuzz")
		} else if div5 {
			result = append(result, "Buzz")
		} else if div3 {
			result = append(result, "Fizz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}

func FizzBuzzApproach2(n int) []string {
	var result []string
	for i := 1; i <= n; i++ {
		tmp := ""
		if i%3 == 0 {
			tmp += "Fizz"
		}
		if i%5 == 0 {
			tmp += "Buzz"
		}
		if tmp == "" {
			tmp = strconv.Itoa(i)
		}
		result = append(result, tmp)
	}
	return result
}

func FizzBuzzApproach3(n int) []string {
	var result []string
	fizzbuzzMap := map[int]string{
		3: "Fizz",
		5: "Buzz",
	}
	// As a Golang map is an unordered collection, it does not preserve the order of keys.
	// We can use additional data structures to iterate over these maps in sorted order.
	// We need to perform the following:
	//Create a slice.
	//Store keys to the slice.
	//Sort the slice by keys.
	//Iterate over the map by the sorted slice.

	keys := make([]int, 0, len(fizzbuzzMap))
	for k := range fizzbuzzMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for i := 1; i <= n; i++ {
		tmp := ""
		for _, v := range keys {
			if i%v == 0 {
				tmp += fizzbuzzMap[v]
			}
		}
		if tmp == "" {
			tmp = strconv.Itoa(i)
		}
		result = append(result, tmp)
	}
	return result
}

// Time Complexity = O(n)
// Space Complexity = O(1)
