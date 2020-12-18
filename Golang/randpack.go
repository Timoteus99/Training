package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomInteger() int {
	//returns a random integer
	return rand.Int()
}

func randomInteger0(n int) int {
	//returns a random integer between 0 and n
	return rand.Intn(n)
}

func randomInRange(min int, max int) int {
	//returns a random integer in a given range
	return rand.Intn(max-min) + min
}

func randomFloat() float64 {
	//returns a random float between 0.0 and 1.0
	return rand.Float64()
}

func randomSlice(n int, m int) []int {
	//return an array of lenghts n with random integers between 0 and m
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = rand.Intn(m)
	}
	return slice
}

func randomPerm(n int) []int {
	//returns a pseudo-random perm of ints from 0 to n
	return rand.Perm(n)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(randomInteger())
	fmt.Println(randomInteger0(100))
	fmt.Println(randomInRange(20, 50))
	fmt.Println(randomFloat())
	fmt.Println(randomSlice(5, 20))
	fmt.Println(randomPerm(5))
}
