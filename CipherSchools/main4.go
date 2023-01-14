package main
import "fmt"

type Vertex struct{
	x int
	y int
}

func main(){
	// struct
	v :=Vertex{x:1, y:2}

	n := Vertex{1,2}

	q := Vertex{}  //x,y are 0
	fmt.Println(v,n,q)


	newVar := newVertex()
	fmt.Println(newVar)
	newVar.SetX(3)
	fmt.Println(newVar)


	X := newVar.x
	Y := newVar.y 
	fmt.Print(X," ",Y)

}

func (v Vertex) GetX() int{
	return v.x
}

//setter
func (v *Vertex)SetX(i int){
	v.x = i
}

//constructor method
func newVertex() Vertex{
	var x int
	var y int
	if x==0{
		x=3
	}
	if y==0{
		y=3
	}
	return Vertex{x,y}
}