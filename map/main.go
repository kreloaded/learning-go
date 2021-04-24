package main

import "fmt"

func main() {
	// var colors map[string]string
	// colors := map[string]string{}
	colors := make(map[string]string)

	colors["white"] = "#FFFFFF"
	fmt.Println(colors)

	colors["red"] = "FF0000"
	fmt.Println(colors)

	delete(colors, "red")
	fmt.Println(colors)
}
