package main

import "fmt"

func main() {
	// var colors1 map[string]string
	colors1 := make(map[string]string)

	colors1["yellow"] = "wait"
	colors := map[string]string{
		"red":   "danger",
		"green": "go",
	}
	fmt.Print(colors1)
	// fmt.Printf("%+v", colors)
	// delete(colors, "green")
	fmt.Printf("%+v\n", colors)
	printMap(colors)
	printMap(colors1)

}

func printMap(mp map[string]string) {
	for key, val := range mp {
		fmt.Println(key, val)
	}
}
