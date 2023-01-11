package main
import "fmt"

var i int = 33
//j := 34  //we can't use this syntax to declare var at package level

func main(){
	// when we declare any var outside the main func that var has scope throughout the file, it's more like a global var
    fmt.Println(i)
	
	//fmt.Println(j)
}