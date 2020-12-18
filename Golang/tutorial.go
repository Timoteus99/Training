package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// UNSIGNED INTEGERS(non negative)

// uint8 / byte (0 - 255)
// uint16, uint32, uint64

// SIGNED INTEGERS

// int8 (-128 to 127)
// int16, int32, int64

// 3 MACHINE DEPENDENT TYPES

// uint (32 or 64 bits)
// int (same size as uint)
// uintptr (unsigned integer to store the uninterpreted bits of a pointer value)

// FLOATING POINT NUMBERS

// float32 (IEE-754 32-bit)
// float64

// complex64 or complex128

// STRINGS

// BOOLEANS

//All types have assigned starting values

func test(x int) {
	fmt.Println(x)
}

func add(x int, y int) int {
	return (x + y)
}

func changeValue(str *string) {
	*str = "changed!"
}

func changeValue2(str string) {
	str = "changed!"
}

func main() {

	var number uint16 = 260
	number = number + 5
	var myname string
	myname = "Timo"
	var myname2 string = "tej"
	fmt.Println("Hello World!", number)
	fmt.Println(myname + myname2)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type a number, and I will substract 10 from it!: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64) // _ --> value is error (if any)
	fmt.Printf("10 less than your number is %d,", input-10)
	fmt.Printf(" abs values of that is: %g", math.Abs(float64(input-10)))

	x := 5
	y := 6
	val := x < y
	fmt.Printf(" %t", val)

	// <, >, ==, !=, <=, >=
	// We can also compare strings and characters
	// Chained conditi --> || or, ! not, && and

	if x < y {
		fmt.Println(", y is greater than x")
	} else if x == y {
		fmt.Println("x and y are equal")
	} else {
		fmt.Println("x is greater than y")
	}

	z := 0
	for z <= 5 {
		fmt.Println(z)
		z++
	}
	for u := 0; u <= 5; u++ { // We can also go u += 2
		fmt.Println(u)
	}

	ans := 10

	switch ans {
	case 1:
		fmt.Println("1. one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("Not a case!")
	}

	switch {
	case ans > 0:
		fmt.Printf("%d is greater than zero!", ans)
	}

	var arr [5]int
	arr[0] = 100
	arr[4] = 900 //pretty much same as py
	fmt.Println(arr)

	arr2 := [3]int{4, 5, 6}
	fmt.Println(arr2)
	// to tell lenght its also len

	sum := 0

	for i := 0; i < len(arr2); i++ {
		sum += arr2[i]
	}

	fmt.Println(sum)

	//multi-dim arrays

	arr2D := [2][3]int{{1, 2, 3}, {3, 3, 4}} //2 arrays inside, with 3 elements
	fmt.Println(arr2D[0][2])

	//slices

	var rezani [5]int = [5]int{1, 2, 3, 4, 5}
	var izrezanec []int = rezani[1:3]
	fmt.Println(izrezanec)
	fmt.Println(cap(izrezanec)) //capacity

	var a []int = []int{5, 6, 7, 8, 9, 5}
	b := append(a, 10) //returns a new slice
	fmt.Println(b)

	//slices with make()

	ab := make([]int, 5)
	fmt.Println(ab)

	for i, element := range a {
		fmt.Printf("%d: %d\n", i, element) //replace i with _ for just elt
	}
	//check for duplicates in array
	for i, element := range a {
		for j, element2 := range a {
			if element == element2 && i < j {
				fmt.Println(element)
			}
		}
	}

	//maping

	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}
	fmt.Println(mp)

	maap := make(map[string]int)
	fmt.Println(maap)

	//acces add change maps

	mp["apple"] = 900
	fmt.Println(mp["apple"])

	delete(mp, "orange")

	mp["trinitrotoluen"] = 10
	fmt.Println(mp)

	value, ok := mp["apple"] //if exists store in val, else val is default for type
	fmt.Println(value, ok)

	//functions

	test(5)

	answer := add(3, 4)
	fmt.Println(answer)

	v := test
	v(134)

	test2 := func(x int) {
		fmt.Println(x)
	}

	test2(5000)

	test3 := func(x int) int {
		return x * -1
	}(8)
	fmt.Println(test3)

	//memory location

	m := 7
	fmt.Println(m, &m)
	n := &m
	*n = 8
	fmt.Println(m, n)

	toChange := "hello!"
	fmt.Println(toChange)
	changeValue(&toChange) //with pointer
	fmt.Println(toChange)

}
