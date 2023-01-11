package main
import "fmt"

var firstName string = "Akanksha"
var lastName string = "Verma"
var age = 12

// above code can be replaced with block code, as shown below
var(
	fname string = "Ak"
	lname string = "V"
	a int = 12
)

func main(){
   fmt.Println(firstName)
   fmt.Println(lastName)
   fmt.Println(fname)
   fmt.Println(lname)
}