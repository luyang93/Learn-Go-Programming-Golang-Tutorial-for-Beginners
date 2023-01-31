package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var writer Writer = ConsoleWriter{}
	writer.Write([]byte("Hello Go!"))

	myIntCounter := IntCounter(0)
	var incrementer Incrementer = &myIntCounter
	for i := 0; i < 10; i++ {
		fmt.Println(incrementer.Increment())
	}

	var wc2 Writer2Closer2 = NewBufferedWriter2Closer2()
	n, err := wc2.Write2([]byte("Hello YouTube listeners, this is a test"))
	fmt.Println(n, err)
	err = wc2.Close2()
	fmt.Println(err)

	bwc := wc2.(*BufferedWriter2Closer2)
	fmt.Println(bwc)

	r, ok := wc2.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}

	r2, ok := wc2.(*BufferedWriter2Closer2)
	if ok {
		fmt.Println(r2)
	} else {
		fmt.Println("Conversion failed")
	}

	var myObj interface{} = NewBufferedWriter2Closer2()
	if wc3, ok := myObj.(Writer2Closer2); ok {
		wc3.Write2([]byte("Hello YouTube listeners, this is a test"))
		wc3.Close2()
	}

	var i interface{} = true
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("I don't know what i is")
	}

	var mwc4 Writer4Closer4 = &myWriter4Closer4{}
	fmt.Println(mwc4)

	var mwc5 Writer5Closer5 = myWriter5Closer5{}
	fmt.Println(mwc5)
}

type Writer interface {
	Write([]byte) (int, error)
}
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

type Writer2 interface {
	Write2([]byte) (int, error)
}

type Closer2 interface {
	Close2() error
}

type Writer2Closer2 interface {
	Writer2
	Closer2
}

type BufferedWriter2Closer2 struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriter2Closer2) Write2(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, nil
		}
	}
	return n, nil
}

func (bwc *BufferedWriter2Closer2) Close2() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriter2Closer2() *BufferedWriter2Closer2 {
	return &BufferedWriter2Closer2{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

type Writer4 interface {
	Write4([]byte) (int, error)
}

type Closer4 interface {
	Close4() error
}

type Writer4Closer4 interface {
	Writer4
	Closer4
}

type myWriter4Closer4 struct{}

func (mwc *myWriter4Closer4) Write4(data []byte) (int, error) {
	return 0, nil
}

func (mwc *myWriter4Closer4) Close4() error {
	return nil
}

type Writer5 interface {
	Write5([]byte) (int, error)
}

type Closer5 interface {
	Close5() error
}

type Writer5Closer5 interface {
	Writer5
	Closer5
}

type myWriter5Closer5 struct{}

func (mwc myWriter5Closer5) Write5(data []byte) (int, error) {
	return 0, nil
}

func (mwc myWriter5Closer5) Close5() error {
	return nil
}
