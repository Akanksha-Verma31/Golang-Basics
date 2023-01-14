package main
import "fmt"

func main(){
	n := "Akanksha"
	nn := HelloName(n)
	fmt.Println("Inside main")
	fmt.Println(nn)

	var name string
    fmt.Println("Enter Name: ")
    fmt.Scan(&name)
	fmt.Println("Your name is ",name)

	greeting := hello("Akanksha")
	fmt.Println(greeting)
}

func HelloName(name string) string{
	fmt.Println("Inside Func")
	fmt.Println("Hello, ",name)
	return "Hello, " + name
}

func hello(name string)(greeting string){
	greeting = fmt.Sprintf("hello %v", name)
	return
}