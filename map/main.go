package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"white": "#FFFFFF",
	}
	fmt.Println(colors)
	delete(colors,"red")
	fmt.Println(colors)
}