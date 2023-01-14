package main
import "fmt"

func main(){
	awake := true
	fmt.Print("Can a Fast Attack be made : ")
	fmt.Println(CanFastAttack(awake))
	var knightIsAwake = false
	var archerIsAwake = false
	var prisonerIsAwake = true
	var petDog = true
	fmt.Print("The group can be spied : ")
	fmt.Println(CanSpy(knightIsAwake,archerIsAwake,prisonerIsAwake))
	fmt.Println(CanSignalPrisoner(archerIsAwake,prisonerIsAwake))
    fmt.Print("Can the Prisoner be freed: ", CanFreePrisoner(knightIsAwake,archerIsAwake,prisonerIsAwake,petDog))
}
func CanFastAttack(awake bool)bool{
	if awake == true{
		return false
	}else{
		return false
	}
}
func CanSpy(knight, archer, prisoner bool)bool{
	return knight || archer || prisoner
}
func CanSignalPrisoner(archer, prisoner bool)bool{
	if archer==false && prisoner==true{
		return true
	}else{
		return false
	}
}
func CanFreePrisoner(knight,archer,prisoner,pet bool)bool{
	if prisoner==true{
		if knight==false && archer==false && pet==true{
			return true
		}
	}
	return false
}