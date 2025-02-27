package main

import "fmt"

func main () {
	message := call1("world")
	fmt.Println(message)
	
	call2(123, "ok")
}

func call1 (name string) string {
	return "call1: Hello, " + name + "!"
}


func call2 (a int, b string) {
	fmt.Printf("call2: %d %s\n", a, b)
}

