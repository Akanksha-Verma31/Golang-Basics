package main
import "fmt"

var i int = 5

func main(){
	fmt.Println(i)
	var i int = 3 //shadowing variable
	fmt.Println(i)
	i = 44
	fmt.Println(i)
}