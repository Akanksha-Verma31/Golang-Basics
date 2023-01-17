package main
import "fmt"
func ShareWith(name interface{})string{
	default_msg := "one for %v, one for me"
	switch v:=name.(type){
	case string:
		return fmt.Sprintf(default_msg,v)
	default:
		return fmt.Sprintf(default_msg,"me")
	}
}
func ShareWith2(name string)string{

	if name==""{
  
	  name="You"
  
	}
  
	return "One for "+name+", One for me"
  
  }