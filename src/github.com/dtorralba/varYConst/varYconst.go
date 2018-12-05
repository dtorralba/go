package main

import "fmt"

var c, python, java bool

const pi float32 = 3.14

var i, j int = 1, 2

func Sumar(x int, y int) int {
	return x + y
}

func main() {
	//Asignacion multiple, infiere tipos var c, python, java = true, false, "no!"
	k := 3 // var k = 3
	var i int
	fmt.Println(Sumar(i, j))
	fmt.Println(i, c, python, java, pi, k)
}
