package main

import "fmt"

func main() {
	// friends := map[string]int{
	// 	"Pratiksha": 24,
	// 	"Gaurav":    25,
	// 	"Ritika" : "26"
	// 	"Aditya":    24,
	// 	"Aachal":    25,
	// }

	//to initialize empty map
	//var friends map[string]string
	friends := make(map[string]int)  //make is inbuilt function
	friends["Pratiksha"] = 24 
	friends["Gaurav"] = 25

	//to delete by inbuilt delete function
	//delete(friends, "Gaurav")

	fmt.Println(friends)

	printMap(friends)
}

func printMap(f map[string]int){
	//iteration order is not fixed
	for name, age := range f {
		fmt.Println(name, "is",age,"yrs old")
	}
}
