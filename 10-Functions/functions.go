package main

import "fmt"

func main() {
	sayMessage("Hello Go!")
	for i := 0; i < 5; i++ {
		sayMessage2("Hello Go!", i)
	}

	greeting := "Hello"
	name := "Stacey"
	sayGreeting(greeting, name)
	fmt.Println(greeting, name)

	greeting1 := "Hello1"
	name1 := "Stacey1"
	sayGreeting2(&greeting1, &name1)
	fmt.Println(greeting1, name1)

	sum("The sum is", 1, 2, 3, 4, 5)

	s1 := sum1(1, 2, 3, 4, 5)
	fmt.Println("The sum is ", s1)

	s2 := sum2(1, 2, 3, 4, 5)
	fmt.Println("The sum is ", *s2)

	s3 := sum3(1, 2, 3, 4, 5)
	fmt.Println("The sum is ", s3)

	d := divide(5.0, 3.0)
	fmt.Println(d)

	d2, err := divide2(5.0, 1.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d2)

	func() {
		fmt.Println("Hello Go!")
	}()

	func() {
		msg := "Hello Go!"
		fmt.Println(msg)
	}()

	for i := 0; i < 5; i++ {
		func() {
			fmt.Println(i)
		}()
	}

	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println(j)
		}(i)
	}

	f := func() {
		fmt.Println("Hello Go!")
	}
	f()

	var divide3 func(float64, float64) (float64, error)
	divide3 = func(a float64, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("cannot divide by zero")
		} else {
			return a / b, nil
		}
	}
	d3, err := divide3(5, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d3)

	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
	fmt.Println("The new name is:", g.name)

	g2 := greeter2{
		greeting: "Hello",
		name:     "Go",
	}
	g2.greet2()
	fmt.Println("The new name is:", g2.name)
}

func sayMessage(msg string) {
	fmt.Println(msg)
}

func sayMessage2(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is", idx)
}

func sayGreeting(greeting string, name string) {
	fmt.Println(greeting, name)
	name = "Ted"
	fmt.Println(name)
}

func sayGreeting2(greeting1 *string, name1 *string) {
	fmt.Println(*greeting1, *name1)
	*name1 = "Ted1"
	fmt.Println(*name1)
}

func sum(msg string, values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
}

func sum1(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

func sum2(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

func sum3(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}

func divide(a, b float64) float64 {
	return a / b
}

func divide2(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = ""
}

type greeter2 struct {
	greeting string
	name     string
}

func (g *greeter2) greet2() {
	fmt.Println(g.greeting, g.name)
	g.name = ""
}
