package main

import "fmt"

func getName()(string, string){
	return "ACM", "UTSA"
}

func main(){
	first, last := getName()
	fmt.Println("Hello, "+first+" "+last+".")
}
