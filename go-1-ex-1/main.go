package main

import "fmt"

func main() {
	// TODO: Declare and initialize the variables being used in the output!

	firstName := "Sabina"
	lastName := "Lazri"
	dayOfBirth := 5
	monthOfBirth := 3
	yearOfBirth := 2002
	numberOfSiblings := 2
	heightInMeters := 1.65
	zodiacSign := '\u2653'

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
