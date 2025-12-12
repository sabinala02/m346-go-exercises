package main

import "fmt"

const (
	Aries       = '\u2648' // Widder
	Taurus      = '\u2649' // Stier
	Gemini      = '\u264a' // Zwillinge
	Cancer      = '\u264b' // Krebs
	Leo         = '\u264c' // Löwe
	Virgo       = '\u264d' // Jungfrau
	Libra       = '\u264e' // Waage
	Scorpius    = '\u264f' // Skorpion
	Sagittarius = '\u2650' // Schütze
	Capricornus = '\u2651' // Steinbock
	Aquarius    = '\u2652' // Wassermann
	Pisces      = '\u2653' // Fische
)

func outputWithZodiacSign(p Person) {
	var zodiacSign rune = '?'

	// TODO: Assign proper value to zodiacSign using if/else branching.
	// NOTE: The runes are defined above as constants.
	d := p.Day
	m := p.Month

	if (m == 3 && d >= 21) || (m == 4 && d <= 19) {
		zodiacSign = Aries
	} else if (m == 4 && d >= 20) || (m == 5 && d <= 20) {
		zodiacSign = Taurus
	} else if (m == 5 && d >= 21) || (m == 6 && d <= 20) {
		zodiacSign = Gemini
	} else if (m == 6 && d >= 21) || (m == 7 && d <= 22) {
		zodiacSign = Cancer
	} else if (m == 7 && d >= 23) || (m == 8 && d <= 22) {
		zodiacSign = Leo
	} else if (m == 8 && d >= 23) || (m == 9 && d <= 22) {
		zodiacSign = Virgo
	} else if (m == 9 && d >= 23) || (m == 10 && d <= 22) {
		zodiacSign = Libra
	} else if (m == 10 && d >= 23) || (m == 11 && d <= 21) {
		zodiacSign = Scorpius
	} else if (m == 11 && d >= 22) || (m == 12 && d <= 21) {
		zodiacSign = Sagittarius
	} else if (m == 12 && d >= 22) || (m == 1 && d <= 19) {
		zodiacSign = Capricornus
	} else if (m == 1 && d >= 20) || (m == 2 && d <= 18) {
		zodiacSign = Aquarius
	} else if (m == 2 && d >= 19) || (m == 3 && d <= 20) {
		zodiacSign = Pisces
	}

	fmt.Printf("%s %s, born on %02d.%02d.%04d, has the zodiac sign %c.\n",
		p.FirstName, p.LastName, p.Day, p.Month, p.Year, zodiacSign)
}

type FullName struct {
	FirstName string
	LastName  string
}
type BirthDate struct {
	Day   byte
	Month byte
	Year  uint16
}

type Person struct {
	FullName
	BirthDate
}

func main() {
	grace := Person{FullName{"Grace", "Hopper"}, BirthDate{9, 12, 1906}}
	dennis := Person{FullName{"Dennis", "Ritchie"}, BirthDate{9, 9, 1941}}
	rick := Person{FullName{"Rick", "Astley"}, BirthDate{6, 2, 1966}}
	edsger := Person{FullName{"Edsger", "Dijkstra"}, BirthDate{11, 5, 1930}}
	alan := Person{FullName{"Alan", "Turing"}, BirthDate{23, 6, 1912}}

	outputWithZodiacSign(grace)
	outputWithZodiacSign(dennis)
	outputWithZodiacSign(rick)
	outputWithZodiacSign(edsger)
	outputWithZodiacSign(alan)
}
