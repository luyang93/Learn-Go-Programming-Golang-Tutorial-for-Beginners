package main

import (
	"fmt"
	"math"
)

func main() {
	if true {
		fmt.Println("The test is true")
	}

	if false {
		fmt.Println("The test is false")
	}

	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}
	if pop, ok := statePopulations["Not"]; ok {
		fmt.Println(pop)
	} else {
		fmt.Println(ok)
	}

	number := 50
	guess := 30
	if guess < number {
		fmt.Println("Too low")
	} else if guess > number {
		fmt.Println("Too high")
	} else if guess == number {
		fmt.Println("You got it!")
	}
	fmt.Println(number <= guess, number >= guess, number != guess)

	number1 := 50
	guess1 := 101
	if guess1 < 1 || returnTrue() || guess1 > 100 {
		fmt.Println("The guess must be between 1 and 100!")
	}
	if guess1 >= 1 && guess1 <= 100 {
		if guess1 < number1 {
			fmt.Println("Too low")
		}
		if guess1 > number1 {
			fmt.Println("Too high")
		}
		if guess1 == number1 {
			fmt.Println("You got it!")
		}
		fmt.Println(number1 <= guess1, number1 >= guess1, number1 != guess1)
	}
	fmt.Println(false)
	fmt.Println(!false)

	number2 := 50
	guess2 := 1
	if guess2 < 1 {
		fmt.Println("The guess must be greater than 1!")
	} else if guess2 > 100 {
		fmt.Println("The guess must be less than 100!")
	} else {
		if guess2 < number2 {
			fmt.Println("Too low")
		}
		if guess2 > number2 {
			fmt.Println("Too high")
		}
		if guess2 == number2 {
			fmt.Println("You got it!")
		}
		fmt.Println(number2 <= guess2, number2 >= guess2, number2 != guess2)
	}

	myNum := 0.1
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}

	myNum2 := 0.12312123123
	if math.Abs(myNum2/math.Pow(math.Sqrt(myNum2), 2)-1) < 0.001 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}

	switch 4 {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	}

	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	}

	i1 := 9
	switch {
	case i1 <= 10:
		fmt.Println("less than or equal to ten")
	case i1 <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

	i2 := 9
	switch {
	case i2 <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough
	case i2 >= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

	var i interface{} = [23]int{}
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is string")
	case [3]int:
		fmt.Println("i is [3]int")
		break
		fmt.Println("This will not print too")
	default:
		fmt.Println("i is another type")
	}
}

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}
