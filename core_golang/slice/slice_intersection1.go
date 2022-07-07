package main

import "fmt"

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func difference(slice1 []string, slice2 []string) []string {
	var diff []string
	var inter []string

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			} else {
				inter = append(inter, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}
	fmt.Println("intersected : ", removeDuplicateStr(inter))

	return diff
}

func main() {
	slice1 := []string{"foo", "bar", "hello"}
	slice2 := []string{"foo", "world", "bar", "foo"}

	fmt.Printf("non intersected %+v\n", difference(slice1, slice2))
}
