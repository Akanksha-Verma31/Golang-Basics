package main
import "fmt"
func main(){
	var i int =27
	fmt.Printf("%v %T\n",i,i)

	k := 99
	fmt.Printf("%v %T\n",k,k)

	j := 9.0
	fmt.Printf("%v %T\n",j,j)
// here it would always float64 as its type never float32
	
}