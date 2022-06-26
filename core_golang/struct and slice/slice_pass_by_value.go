package main

import "fmt"

func main() {
	mySlice := []string{"Hii","How","are","you"}
	fmt.Println(mySlice)
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(slice []string){
	slice[0] = "Hello"
}