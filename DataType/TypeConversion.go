package main
import "fmt"
import "strconv"
func main(){
	var i int = 43
	fmt.Printf("%v %T\n",i,i)

	// after type conversion
	var f float32
	f = float32(i)   // conversion operator
	fmt.Printf("%v %T\n",f,f)

	// converting int to float is okay but when we do the vice versa there are chances of data loss
	var ff float32 = 34.5
	fmt.Printf("%v %T\n",ff,ff)

	var ii int
	ii = int(ff)
	fmt.Printf("%v %T\n",ii,ii)  // 34.5 is converted to int but 0.5 data is lost during conversion

	// int -- string conversion
	var str string 
	str = string(i)        // here, the int 43 is converted to its unicode and hence, its equivalent char value is printed
	fmt.Printf("%v %T\n",str,str) 

	// string -- int conversion
    var s string
	var s2 string = "100"
	fmt.Printf("%v %T\n",i,i)
	s = strconv.Itoa(i)
	fmt.Printf("%v %T\n",s,s)
	i , _ = strconv.Atoi(s2)
	fmt.Printf("%v %T\n",i,i)
}