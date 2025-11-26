package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

type Demo[S, T any] struct{}

func (d Demo[S, T]) foo(a ...string) {

}
