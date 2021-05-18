package main

import "fmt"

func main() {
	// var colors map[string]string
	// colors := make(map[string]string)

	// delete(colors, "white")

	colors := map[string]string{
		"red":   "0xff0000",
		"green": "0x00ff00",
		"blue":  "0x0000ff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}

/*
MAP
- keys and values must be the same type
- keys are indexed, can be iterated
- reference type!

*/

/*
STRUCT
- keys do not support indexing
- values can of different types
- value type!
- fields are defined at compile time
*/
