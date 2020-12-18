package main

import (
	"log"
	"os"
)

func main() {
	testfile, err := os.Create("testfile.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(testfile)
	testfile.Close()
}
