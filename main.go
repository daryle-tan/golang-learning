package main

import "fmt"

func printUsernamePassword() {
	var username string ="wagslane"
	var password int = 20947382822

	passwordStr := fmt.Sprintf("%d", password)
	fmt.Println("Authorizationoo: Basic", username+":"+passwordStr)
}

// whole numbers
// int, int8, int16, int32, int64
// positive whole numbers
// uint, uint8, uint16, uint32, uint64, uintptr
// decimal numbers
// float32, float64
// imaginary numbers (rare)
// complex64, complex128

var x uint = 42
var y uint = 2

func add(x, y int) int {
fmt.Println(add(x,y))
	return x + y
	
}

const helloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
