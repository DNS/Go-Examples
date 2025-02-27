package main

import "fmt"
import "strconv"


func main () {
	i := 5
	
	f := float64(i)
	u := uint(i)
	
	s1 := strconv.Itoa(i)
	s2 := fmt.Sprintf("%d", i)
	
	fmt.Println(i, f, u, s1, s2)
}

