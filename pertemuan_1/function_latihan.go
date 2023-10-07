package main

import (
	"fmt"
)

func main() {
	car := make(map[string]string)
	car["name"] = "BMW"
	car["color"] = "Black"

	message := generateMessage(car)
	displayMessage(message)
}

func generateMessage(car map[string]string) string {
	return "Mobil " + car["name"] + " berwarna " + car["color"]
}

func displayMessage(message string) {
	fmt.Println(message)
}
