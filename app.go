package main

import (
	"crypto/rand"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/sgdfgo/serv"
)

func main() {
	port := 5163
	if len(os.Args) == 2 {
		var err error
		port, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Printf("Invalid port value %s", os.Args[1])
		}
	}

	keyLength := 32
	if _, err := os.Stat("secret.hmac.bin"); os.IsNotExist(err) {
		f, err := os.Create("secret.hmac.bin")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		serv.Secret = make([]byte, keyLength)
		_, err = rand.Read(serv.Secret)
		if err != nil {
			panic(err)
		}
		_, err = f.Write(serv.Secret)
		if err != nil {
			panic(err)
		}
	} else {
		f, err := os.Open("secret.hmac.bin")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		serv.Secret, err = ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
	}

	server := serv.New(port)
	server.Start()

	//

	// s := scraper.New()
	// s.Connect(os.Args[1], os.Args[2])
	// s.ScrapStructures()
	// s.ScrapExport()
	// data := s.Export()
	// fmt.Println(data)
}
