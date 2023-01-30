package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("start")
	fmt.Println("middle")
	fmt.Println("end")

	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")

	res, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

	a := "start1111"
	defer fmt.Println(a)
	a = "end1111"

	//a1, b1 := 1, 0
	//ans := a1 / b1
	//fmt.Println(ans)

	fmt.Println("start")
	//panic("something bad happened")
	fmt.Println("end")

	fmt.Println("start4444")
	panicker()
	fmt.Println("end4444")

	fmt.Println("start3333")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("something bad happened")
	fmt.Println("end3333")

	fmt.Println("start2222")
	defer fmt.Println("this was deferred")
	panic("something bad happened")
	fmt.Println("end2222")

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello Go!"))
	})
	err2 := http.ListenAndServe(":8080", nil)
	if err2 != nil {
		panic(err2.Error())
	}
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			panic(err)
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")

}
