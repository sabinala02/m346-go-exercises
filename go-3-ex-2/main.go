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

func outputDateRange(zodiacSign rune) {
	fmt.Printf("%c: ", zodiacSign)

	switch zodiacSign {
	case Aries:
		fmt.Println("21.03. - 20.04")
	case Taurus:
		fmt.Println("21.04. - 20.05")
	case Gemini:
		fmt.Println("21.05. - 20.06")
	case Cancer:
		fmt.Println("21.06. - 20.07")
	case Leo:
		fmt.Println("21.07. - 20.08")
	case Virgo:
		fmt.Println("21.08. - 20.09")
	case Libra:
		fmt.Println("21.09. - 20.10")
	case Scorpius:
		fmt.Println("21.10. - 20.11")
	case Sagittarius:
		fmt.Println("21.11. - 20.12")
	case Capricornus:
		fmt.Println("21.12. - 20.01")
	case Aquarius:
		fmt.Println("21.01. - 20.02")
	case Pisces:
		fmt.Println("21.02. - 20.03")
	default:
		fmt.Println("invalid star sign")
	}
}

func main() {
	for zodiacSign := Aries; zodiacSign <= Pisces; zodiacSign++ {
		outputDateRange(zodiacSign)
	}
}
