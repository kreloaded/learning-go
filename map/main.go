package main

import "fmt"

func mapCreateMethods() {
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

func main() {
	// mapCreateMethods()
	colors := map[string]string{
		"red":   "#FF0000",
		"white": "#FFFFFF",
		"green": "#4BF745",
	}

	fmt.Println(colors)

	printMap(colors)
}

func printMap(mapToPrint map[string]string) {
	for color, hexCode := range mapToPrint {
		fmt.Println("Hex code for", color, "is", hexCode)
	}
}
