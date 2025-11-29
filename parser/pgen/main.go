package main

import (
	"fmt"
	"github.com/lincaiyong/pgen"
	"os"
	"time"
)

func main() {
	start := time.Now()
	grammar, err := pgen.PreProcess("grammar/go.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_ = os.WriteFile("grammar/go.log", []byte(grammar), 0644)
	output, err := pgen.Run(grammar)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_ = os.Mkdir("goparser", os.ModePerm)
	err = os.WriteFile("goparser/goparser.go", []byte(output), 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("finished in %s\n", time.Since(start))
}
