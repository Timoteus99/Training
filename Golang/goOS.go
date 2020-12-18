package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// reading entire file into memory
	data, err := ioutil.ReadFile("testGO.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))

	// using an absolute file path
	data2, err2 := ioutil.ReadFile("/home/legion/code/trinitrotoluen/Training/testGO.txt")
	if err2 != nil {
		fmt.Println("File reading error", err2)
		return
	}
	fmt.Println("Contents of file:", string(data2))

	// passing the file path as a command-line flag

	fptr := flag.String("fpath", "testGO.txt", "file path to read from")
	flag.Parse()
	data3, err3 := ioutil.ReadFile(*fptr)
	if err3 != nil {
		fmt.Println("File reading error", err3)
		return
	}
	fmt.Println("Contents of file:", string(data3))

	// reading a file in small chuncks

	path := "testGO.txt"

	buf, err4 := os.Open(path)
	if err4 != nil {
		log.Fatal(err4)
	}

	defer func() {
		if err4 = buf.Close(); err4 != nil {
			log.Fatal(err4)
		}
	}()

	r := bufio.NewReader(buf)
	b := make([]byte, 8)
	for {
		n, err4 := r.Read(b)
		if err4 != nil {
			fmt.Println("Error reading file:", err4)
			break
		}
		fmt.Println(string(b[0:n]))
	}

	// read a file line by line

	// path already defined

	buf2, err5 := os.Open(path)
	if err5 != nil {
		log.Fatal(err5)
	}

	defer func() {
		if err5 = buf2.Close(); err5 != nil {
			log.Fatal(err5)
		}
	}()

	snl := bufio.NewScanner(buf2)
	for snl.Scan() {
		fmt.Println(snl.Text())
	}
	err5 = snl.Err()
	if err5 != nil {
		log.Fatal(err5)
	}
}
