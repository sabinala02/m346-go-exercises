package main

import "fmt"

func main() {
	// TODO: create a map called "modules"
	modules := map[int16]string{
		104: "Modul 104",
		117: "Filius",
		346: "Go",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// TODO: delete one
	delete(modules, 104)

	// TODO: add one
	modules[162] = "Daten"

	// TODO: replace one
	modules[162] = modules[254]
	modules[254] = "Prozesse"

	fmt.Println(modules)
}
