package main

import "fmt"
	//myOwnShortForm "github.com/stretchr/ssssasfsas"

func main() {
	fmt.Println("Hello world\n")
	Print("hello")

	name := "Akanksha Verma"
	fmt.Printf("hello %v", name)
	fmt.Print("\n")
	//%d- int %s- string %w- error %v-any

	name1 := "aa"
	name2 := "bb"
	name3 := "cc"
	fmt.Printf("hello %v %v %v", name1, name2, name3)

	fmt.Printf("hello %v %v %v", name1, name1, name1)

	fmt.Printf("hello %[1]v %[1]v %[1]v", name1)

	var name5 string = "tarun"
	fmt.Println(name5)

	greetingMessage := Greeting("tarun")
	fmt.Println(greetingMessage)

	v1 := GreetingWithOutputPara("rahul")

	fmt.Println(v1)

}

func Print(input string) {
	fmt.Println(input)
}

func Greeting(personName string) string {

	return "hello " + personName
}

func GreetingWithOutputPara(personName string) (greetingMessage string) {
	if personName == "tarun" {
		greetingMessage = "hello sir"
		return
	}
	greetingMessage = "hello" + personName
	return
}

//GreetingWithVarNumOfPara("tarun","rahul","kanit")

func GreetingWithVarNumOfPara(personNames ...string) (greetingMessages []string) {
	return
}
