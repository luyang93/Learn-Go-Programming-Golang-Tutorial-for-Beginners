package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			i = i / 2
		} else {
			i = 2*i + 1
		}
	}

	i := 2
	for ; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println(i)

	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	k := 0
	for {
		fmt.Println(k)
		k++
		if k == 10 {
			break
		}
	}

	for i1 := 0; i1 < 10; i1++ {
		if i1%2 == 0 {
			continue
		}
		fmt.Println(i1)
	}

	for i2 := 0; i2 <= 3; i2++ {
		for j2 := 0; j2 <= 3; j2++ {
			fmt.Println(i2 * j2)
		}
	}

	fmt.Println("---------")
	for i3 := 0; i3 <= 3; i3++ {
		for j3 := 0; j3 <= 3; j3++ {
			fmt.Println(i3 * j3)
			if i3*j3 >= 3 {
				break
			}
		}
	}

	fmt.Println("---------")
Loop:
	for i3 := 0; i3 <= 3; i3++ {
		for j3 := 0; j3 <= 3; j3++ {
			fmt.Println(i3 * j3)
			if i3*j3 >= 3 {
				break Loop
			}
		}
	}

	s := []int{1, 2, 3}
	fmt.Println(s)
	for k, v := range s {
		fmt.Println(k, v)
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
	for k, v := range statePopulations {
		fmt.Println(k, v)
	}
	for _, v := range statePopulations {
		fmt.Println(v)
	}
	for k := range statePopulations {
		fmt.Println(k)
	}

	s1 := "hello GO!"
	fmt.Println(s1)
	for k, v := range s1 {
		fmt.Println(k, string(v))
	}
}
