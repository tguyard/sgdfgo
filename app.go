package main

import (
	"fmt"
	"os"

	"github.com/sgdfgo/scraper"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Rentrer le login/motDePasse")
		os.Exit(-1)
	}
	s := scraper.New()
	s.Connect(os.Args[1], os.Args[2])
	// s.ScrapStructures()
	s.ScrapExport()
	// data := s.Export()
	// fmt.Println(data)
}
