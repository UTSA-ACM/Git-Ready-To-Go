package main
import "fmt"
func getName()(string, string){
	return "Andrew", "Hutton"
}
func main(){
	first, last := getName()
	fmt.Println("Hello, "+first+" "+last+".")
}
