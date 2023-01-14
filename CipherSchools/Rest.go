package main
import {
	"log"
	"encoding/json"
	"fmt"
}
type Book struct{
	Title string `json:"booktitle"`
	Author string
	Page int
	Price int
}

func main(){
	book := Book{}
	book.Title="Game of Thrones"
	book.Author="Ashish"
	book.Page=1000
	book.Price=2.5

	//book obj --> book.json ---- marshall
	byteArray, err := json.Marshal(book)
	if err != nil{
		log.Fatal("Got error", err)
	}
	log.Println(string(byteArray))


	// json --> book object --- unmarshal
	book2:=Book{}
	jsondata := `{"booktitle":"game","author":"ashish","page":1000}`
	json.unmarshal([]byte(jsondata), &book2)
	if err!=nil{
		log.Fatal("got error while unmarshal", err)
	}
	fmt.Println(book2.Title)   // some title
}