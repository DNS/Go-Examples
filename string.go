package main

import "fmt"



func main () {
	str1 := "Hello"

	str2 := `Multiline
	strings`
	
	fmt.Println(str1)
	fmt.Println(str2)
	
	// string concatenation
	str3 := str1 + " concat !!!"
	fmt.Println(str3)
}

