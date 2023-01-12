package main
import "fmt"
func main(){
	var name string = "Akanksha"  //var declared in lowercase has scope in the package only
	var Surname string = "Verma"  //var declared in camelcase is globally available 
	fmt.Println(name)
	fmt.Println(Surname)
}
// can't use the below print st as 'name' var has scope limited to the main func only but 'Surname' is globally visible
//fmt.Println(name)
