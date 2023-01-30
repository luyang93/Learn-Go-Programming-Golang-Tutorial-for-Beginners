package main

import "fmt"

func main() {
	a := 43
	b := a
	fmt.Println(a, b)
	a = 27
	fmt.Println(a, b)

	var a1 int = 42
	var b1 *int = &a1
	fmt.Println(&a1, b1)
	fmt.Println(a1, *b1)
	a1 = 27
	fmt.Println(a1, *b1)
	*b1 = 14
	fmt.Println(a1, *b1)

	a2 := [3]int{1, 2, 3}
	b2 := &a2[0]
	c2 := &a2[1]
	fmt.Printf("%v %p %p\n", a2, b2, c2)

	var ms *myStruct
	var ms2 *myStruct
	fmt.Println(ms, ms2)
	ms = &myStruct{foo: 42}
	ms2 = new(myStruct)
	fmt.Println(ms, ms2)
	(*ms2).foo = 44
	fmt.Println(ms, ms2)
	ms2.foo = 100
	fmt.Println(ms2.foo)

	a3 := []int{1, 2, 3}
	b3 := a3
	fmt.Println(a3, b3)
	a3[1] = 42
	fmt.Println(a3, b3)

	a4 := map[string]string{"foo": "bar", "baz": "buz"}
	b4 := a4
	fmt.Println(a4, b4)
	a4["foo"] = "qux"
	fmt.Println(a4, b4)
}

type myStruct struct {
	foo int
}
