package main

import (
	"fmt"
	"math"
)

const myConst int16 = 27

const (
	x = iota
	y = iota
	z = iota
)

const (
	x2 = iota
	y2
	z2
)

const x3 = iota

const (
	_ = iota
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

const (
	_ = iota + 10
	catSpecialist1
	dogSpecialist1
	snakeSpecialist1
)

const (
	_  = iota // ignore first value by signing to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	ZB
	YB
)
const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeddFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	fmt.Printf("%v, %T\n", myConst, myConst)
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)

	const myConst1 float64 = math.Pi
	fmt.Printf("%v, %T\n", myConst1, myConst1)

	const a int = 14
	const b string = "foo"
	const c float32 = 3.14
	const d bool = true
	fmt.Printf("%v, %v, %v, %v\n", a, b, c, d)

	var b2 int = 27
	fmt.Printf("%v, %T\n", a+b2, a+b2)

	const a3 = 14
	var b3 float64 = 14
	fmt.Printf("%v, %T\n", a3+b3, a3+b3)

	const a4 = iota
	fmt.Printf("%v, %T\n", a4, a4)

	fmt.Printf("%v, %v, %v\n", x, y, z)

	fmt.Printf("%v, %v, %v, %v\n", x2, y2, z2, x3)

	var specialistType int = catSpecialist
	fmt.Printf("%v\n", specialistType == catSpecialist)
	fmt.Printf("%v\n", specialistType == dogSpecialist)

	var specialistType2 int
	fmt.Printf("%v\n", specialistType2 == catSpecialist)

	fmt.Printf("%v\n", catSpecialist1)

	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)

	var roles byte = isAdmin | canSeddFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHeadquarters&roles == isHeadquarters)

}
