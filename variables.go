package main

import "fmt"


var a string
var b = "Hello, world!"
var c string = "Hello, world!"
var d, e int
var f, g int = 1, 2
var h, i = 1, "Hello, world!"
var j float64 = 0.12345
const k int64 = 123456789
const l, m = 1, 2
const (
	Z1 = 3.14
	Z2 = "abc"
	Z3
)

func main () {
	var msg = "Hello, world!"
	fmt.Println(msg)
	
	msg1 := "Hello"
	msg2 := 123
	fmt.Printf("%s %d\n", msg1, msg2)
}

