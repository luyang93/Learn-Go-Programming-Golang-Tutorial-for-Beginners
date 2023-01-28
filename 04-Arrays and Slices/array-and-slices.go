package main

import "fmt"

func main() {
	grade1 := 97
	grade2 := 85
	grade3 := 93
	fmt.Printf("Grades: %v, %v, %v\n", grade1, grade2, grade3)

	grades := [3]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)

	grades4 := [...]int{97, 85, 93}
	fmt.Printf("Grades4: %v\n", grades4)

	var students [5]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	students[1] = "Ahmed"
	students[2] = "Arnold"
	fmt.Printf("Student #1: %v\n", students[1])
	fmt.Printf("Number of Students: %v\n", len(students))

	var identityMatrix [3][3]int = [3][3]int{
		[3]int{1, 0, 0},
		[3]int{0, 1, 0},
		[3]int{0, 0, 1},
	}
	var identityMatrix2 [3][3]int = [3][3]int{}
	identityMatrix2[0] = [3]int{1, 0, 0}
	identityMatrix2[1] = [3]int{0, 1, 0}
	identityMatrix2[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix, identityMatrix2)

	a := [...]int{1, 2, 3}
	b := a
	fmt.Println(a)
	b[1] = 5
	fmt.Println(a, b)

	a1 := []int{1, 2, 3}
	b1 := a1
	fmt.Println(a1)
	b1[1] = 5
	fmt.Println(a1, b1)

	a2 := [...]int{1, 2, 3}
	b2 := &a2
	fmt.Println(a2)
	b2[1] = 5
	fmt.Println(a2, *b2)

	a3 := []int{1, 2, 3}
	fmt.Println(a3)
	fmt.Printf("Length: %v, Capacity: %v\n", len(a), cap(a))

	a4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b4 := a4[:]   // slice of all elements
	c4 := a4[3:]  // slice from 4th element to end
	d4 := a4[:6]  // slice first 6 elements
	e4 := a4[3:6] // slice the 4th, 5th, and 6th elements
	fmt.Println(a4, b4, c4, d4, e4)
	a4[5] = 42
	fmt.Println(a4, b4, c4, d4, e4)

	a5 := make([]int, 3)
	fmt.Println(a5)
	fmt.Printf("Length: %v, Capacity: %v\n", len(a5), cap(a5))

	a6 := make([]int, 3, 100)
	fmt.Println(a6)
	fmt.Printf("Length: %v, Capacity: %v\n", len(a6), cap(a6))
	a6 = append(a6, 1)
	fmt.Println(a6)
	fmt.Printf("Length: %v, Capacity: %v\n", len(a6), cap(a6))

	a7 := []int{}
	fmt.Println(a7)
	fmt.Printf("Length: %v, Capacity: %v\n", len(a7), cap(a7))
	a7 = append(a7, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(a7)
	fmt.Printf("Length: %v, Capacity: %v\n", len(a7), cap(a7))
	a7 = append(a7, []int{1, 2, 3, 4, 5, 6, 7}...)
	fmt.Println(a7)
	fmt.Printf("Length: %v, Capacity: %v\n", len(a7), cap(a7))

	a8 := []int{1, 2, 3, 4, 5}
	fmt.Println(a8)
	b8 := a8[1:]
	b81 := append(a8[:2], a8[3:]...)
	fmt.Println(a8, b8, b81)
}
