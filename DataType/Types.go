package main

import "fmt"

func main(){

	//boolean data type
	var bb bool = false
	fmt.Printf("%v %T\n",bb,bb)

	// boolean operations
	n := 1 
	fmt.Println(n==1)
	m := 1
	fmt.Println(m==2)
	// n:=1 == 1

	//integer data type
	var i int
	fmt.Printf("%v %T\n",i,i)

	// int operations
	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	var c int = 10    //1010
	var d int8 = 3    //0011
	//fmt.Println(c + d)    -- mismatched types
	fmt.Println(c + int(d))

	//Bitwise operators
	fmt.Println(a & b)     // 0010
	fmt.Println(a | b)     // 1011
	fmt.Println(a ^ b)     // 1001
	fmt.Println(a &^ b)    // 0100

	// Bit shift
	e := 8    // 2^3
	fmt.Println(e << 3)   // 2^3 * 2^3 = 2^6
	fmt.Println(e >> 3)   // 2^3 / 2^3 = 2^0
	
	//float data types
	var nn float32 = 3.14
	nn = 2.1E14
	fmt.Printf("%v %T\n",nn,nn)

	// complex data types
	var x complex64 = 1 + 2i
	fmt.Printf("%v %T\n",x,x)
    var y complex64 = 2 + 1i
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	
	// real() imag() functions
	fmt.Printf("%v %T\n",real(x),real(y))
	fmt.Printf("%v %T\n",imag(x),imag(y))

    // complex() function
	var z complex64 = complex(2,4)
	fmt.Printf("%v %T\n",z,z)

	//string data type
	var str2 string = "This is a string"
	fmt.Printf("%v %T\n", str2[2], str2[2])
	fmt.Printf("%v %T\n", string(str2[2]), str2[2])

	var str3 string = " and bring"
	fmt.Printf("%v %T\n", str2 + str3, str2+str3)
	w := []byte(str3)
	fmt.Printf("%v %T\n", w, w)
	
}