package main

import "fmt"

func Somar(a int, b int) int {
	return a + b
}

func main(){
	fmt.Printf("Resultado: %d", Somar(2,2))	
}