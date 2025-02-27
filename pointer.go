package main

import "fmt"


func main () {
	b := *getPointer()
	fmt.Println("b:", b)
	
	x := 222
	passPointer(&x)
	fmt.Println("x:", x)
}

func getPointer () (myPointer *int) {
	a := 888
	return &a
}


func passPointer (myPointer *int) {
	*myPointer = 777
}

