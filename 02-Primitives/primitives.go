package main

import "fmt"

func main() {
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	n1 := 1 == 1
	n2 := 1 == 2
	fmt.Printf("%v, %T\n", n1, n1)
	fmt.Printf("%v, %T\n", n2, n2)

	var n3 bool
	fmt.Printf("%v, %T\n", n3, n3)

	m := 42
	fmt.Printf("%v, %T\n", m, m)

	var m1 uint16 = 42
	fmt.Printf("%v, %T\n", m1, m1)

	a := 10 // 1010
	b := 3  // 0011
	fmt.Println(a+b, a-b, a*b, a/b, a%b)
	fmt.Println(a & b)  // 0010
	fmt.Println(a | b)  // 1011
	fmt.Println(a ^ b)  // 1001
	fmt.Println(a &^ b) // 1000

	var a1 int = 10
	var b1 int8 = 3
	fmt.Println(a1 + int(b1))

	a2 := 8                   // 00001000
	fmt.Println(a2<<3, a2>>3) // 01000000, 00000001

	var n4 float64 = 3.14
	fmt.Printf("%v, %T\n", n4, n4)
	n4 = 13.7e72
	fmt.Printf("%v, %T\n", n4, n4)
	n4 = 2.1e14
	fmt.Printf("%v, %T\n", n4, n4)

	a3 := 10.2
	b3 := 3.7
	fmt.Println(a3+b3, a3-b3, a3*b3, a3/b3)

	var n5 complex64 = 1.1 + 2.1i
	fmt.Printf("%v, %T\n", n5, n5)
	fmt.Printf("%v, %T\n", real(n5), real(n5))
	fmt.Printf("%v, %T\n", imag(n5), imag(n5))

	var n6 complex128 = 1.1 + 2.1i
	fmt.Printf("%v, %T\n", n6, n6)
	fmt.Printf("%v, %T\n", real(n6), real(n6))
	fmt.Printf("%v, %T\n", imag(n6), imag(n6))

	a4 := 1 + 2i
	b4 := 2 + 5.2i
	fmt.Println(a4+b4, a4-b4, a4*b4, a4/b4)

	var n7 complex128 = complex(1.1, 1.1)
	fmt.Printf("%v, %T\n", n7, n7)
	fmt.Printf("%v, %T\n", real(n7), real(n7))
	fmt.Printf("%v, %T\n", imag(n7), imag(n7))

	s := "this is a string"
	s3 := "sjfalsdjflaskjdl"
	fmt.Printf("%v, %T\n", len(s), s)
	fmt.Printf("%v, %T\n", s[2], s[2])
	fmt.Printf("%v, %T\n", s+s3, s+s3)

	s2 := "中文is is a string"
	fmt.Printf("%v, %T\n", len(s2), s2)
	fmt.Printf("%v, %T\n", s2[0:3], s2[0:3])

	bytes := []byte(s2)
	fmt.Printf("%v, %T\n", bytes, bytes)

	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)

}
