package main

import "fmt"

func hello() string {
	return "Hello, World!"
}

func PrintHelloWithName(name string) {
	fmt.Println(fmt.Sprintf("hello, %s", name))
}
