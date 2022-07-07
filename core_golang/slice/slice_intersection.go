package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func interSection[T constraints.Ordered](pS ...[]T) (result []T) {
	hash := make(map[T]*int) // value, counter
	for _, slice := range pS {
		for _, value := range slice {
			if counter := hash[value]; counter != nil {
				*counter++
			} else {
				i := 1
				hash[value] = &i
			}
		}
	}
	result = make([]T, 0)
	for value, counter := range hash {
		if *counter >= len(pS) {
			result = append(result, value)
		}
	}
	return
}

func main() {
	slice1 := []string{"foo", "bar", "hello"}
	slice2 := []string{"foo", "bar"}
	fmt.Println(interSection(slice1, slice2))

	ints1 := []int{1, 2, 3, 9}
	ints2 := []int{10, 4, 2, 4, 8, 9}
	ints3 := []int{2, 4, 8, 1}
	fmt.Println(interSection(ints1, ints2, ints3))
}
