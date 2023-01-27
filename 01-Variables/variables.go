package main

import (
	"fmt"
	"strconv"
)

var j int = 77
var j2 float32 = 77.
var I int = 11

var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

var (
	counter int = 0
)

func main() {
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", j2, j2)

	fmt.Println("Hello, playground")

	fmt.Println(42)

	var i int
	i = 42
	fmt.Println(i)

	var i2 int
	i2 = 42
	i2 = 27
	fmt.Println(i2)

	var i3 int = 43
	fmt.Println(i3)

	i4 := 44
	fmt.Println(i4)
	fmt.Printf("%v, %T\n", i4, i4)

	var i5 float32 = 27
	fmt.Printf("%v, %T\n", i5, i5)

	i6 := 44.
	fmt.Printf("%v, %T\n", i6, i6)

	fmt.Println(j)
	var j int = 100
	fmt.Println(j)

	fmt.Println(I)

	var i7 int = 44
	fmt.Printf("%v, %T\n", i7, i7)
	var j7 float32
	j7 = float32(i7)
	fmt.Printf("%v, %T\n", j7, j7)

	var i8 float32 = 43.5
	fmt.Printf("%v, %T\n", i8, i8)
	var j8 int
	j8 = int(i8)
	fmt.Printf("%v, %T\n", j8, j8)

	var i9 int = 43
	fmt.Printf("%v, %T\n", i9, i9)
	var j9 string
	j9 = strconv.Itoa(i9)
	fmt.Printf("%v, %T\n", j9, j9)
}
