package main
import "fmt"
func main(){
	myArray := [5]int{1,2,3,4,5}
	fmt.Println(myArray)

	mySlice := []int{1,2,3,4}
	fmt.Println(mySlice)

	myslice2:=make([]int,10)
	fmt.Println(myslice2)

	myslice3:=make([]int,0,15)
    fmt.Println(myslice3)

    myslice4 := make([]int,0)
	for i:=0;i<10;i++{
		myslice4 = append(myslice4, 3)
	}
	myslice4[1]=45
	fmt.Println(myslice4)

	//new func returns the pointer to the object but make the func returns the object
	//myslice5 := new([]int)

	slice5:=myslice4[0:10]    //including 0 excluding 10
	slice5[4]=123
	fmt.Println(myslice4)
	fmt.Println(slice5)

	//delete index 4
	slice5=append(slice5[:4], slice5[5:]...)   // ... so that all the remining elements got appedded

	// length of slice
	l:=len(slice5)
	fmt.Println(l)

	for i:=0; i<len(slice5); i++{
		slice5[i]=slice5[i]-5
		fmt.Println(slice5)
	}
	for i:=range slice5{
		slice5[i]=slice5[i]-5
		fmt.Println(slice5[i])
	}
	for index, value :=range slice5{
		slice5[index] = value-5
		fmt.Println("index: ",index," value: ",value)
	}

}