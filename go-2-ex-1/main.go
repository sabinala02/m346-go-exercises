package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

type BirthDate struct {
	DayOfBirth   byte
	MonthOfBirth byte
	YearOfBirth  int16
}

type Profile struct {
	Name             FullName
	Birthday         BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		Name: FullName{
			FirstName: "",
			LastName:  "",
		},
		Birthday: BirthDate{
			DayOfBirth:   5,
			MonthOfBirth: 3,
			YearOfBirth:  2002,
		},
		NumberOfSiblings: 2,
		ZodiacSign:       '\u2653',
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	me.NumberOfSiblings = 3
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
