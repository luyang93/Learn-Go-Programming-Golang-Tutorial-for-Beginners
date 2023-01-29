package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int
	actorName  string
	episodes   []string
	companions []string
}

type Doctor2 struct {
	Number     int
	ActorName  string
	Episodes   []string
	Companions []string
}

type Animal struct {
	Name   string `required max: "100"`
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	fmt.Println(statePopulations)
	statePopulations["Georgia"] = 10310371
	fmt.Println(statePopulations)
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)
	pop, ok := statePopulations["Oho"]
	fmt.Println(pop, ok)
	_, ok = statePopulations["Ohio"]
	fmt.Println(ok)
	sp := statePopulations
	fmt.Println(sp)
	fmt.Println(statePopulations)
	delete(sp, "Ohio")
	fmt.Println(sp)
	fmt.Println(statePopulations)

	m := map[[3]int]string{}
	fmt.Println(m)

	statePopulations2 := make(map[string]int)
	fmt.Println(statePopulations2)

	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[1])

	aDoctor1 := Doctor{
		3,
		"Jon Pertwee",
		[]string{"1", "2"},
		[]string{
			"Liz Shaw",
			"Jo Grant",
		},
	}
	fmt.Println(aDoctor1)
	fmt.Println(aDoctor1.actorName)
	fmt.Println(aDoctor1.companions[1])

	aDoctor2 := Doctor2{
		Number:    3,
		ActorName: "Jon Pertwee",
		Companions: []string{
			"Liz Shaw",
			"Jo Grant",
		},
	}
	fmt.Println(aDoctor2)
	fmt.Println(aDoctor2.ActorName)
	fmt.Println(aDoctor2.Companions[1])

	aDoctor3 := struct {
		name string
	}{name: "John Pertwee"}
	anotheraDoctor3 := aDoctor3
	fmt.Println(aDoctor3)
	anotheraDoctor3.name = "Tom Baker"
	fmt.Println(aDoctor3)
	fmt.Println(anotheraDoctor3)

	aDoctor4 := struct {
		name string
	}{name: "John Pertwee4"}
	anotheraDoctor4 := &aDoctor4
	fmt.Println(aDoctor4)
	anotheraDoctor4.name = "Tom Bake4"
	fmt.Println(aDoctor4)
	fmt.Println(*anotheraDoctor4)

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b.Name)

	b2 := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 30,
		CanFly:   false,
	}
	fmt.Println(b2.Name)

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
