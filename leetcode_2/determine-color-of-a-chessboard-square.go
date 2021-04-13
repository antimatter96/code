package main

import "fmt"

func squareIsWhite(coordinates string) bool {
	return (coordinates[0]-97+coordinates[1]-48)%2 == 0
}

func driver__squareIsWhite() {
	fmt.Println(squareIsWhite("a1") == false)
	fmt.Println(squareIsWhite("h3") == true)
	fmt.Println(squareIsWhite("c7") == false)
}
