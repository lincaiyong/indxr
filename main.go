package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	d := Demo{
		T: [10]string{"x"},
	}
	fmt.Println(d)
}

const C = 10

type Demo struct {
	X, T [C]string `json:"t,omitempty"`
}

func (d Demo) foo(a ...string) {

}
