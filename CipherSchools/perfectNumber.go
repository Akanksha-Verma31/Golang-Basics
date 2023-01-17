package main
import "fmt"
func main(){
	var num int64
	fmt.Println("Enter number : ")
	fmt.Scan(&num)
	TypeOfNumber(num)
	var error = errors.New("Pls enter a positive number")
}
var ErrOnlyPositive = errors.New("Plz enter a positive number")
type Classification int
const (
	ClassificationDeficient Classification = 0

	ClassificationPerfect   Classification = 1

	ClassificationAbundant  Classification = 2

	ClassificationInvalid   Classification = 3
)
func TypeOfNymber(input int64) (Classification, error) {
	var sum_of_the_factor int64
	if input <= 0 {
		return ClassificationInvalid, ErrOnlyPositive
	}
	for i := int64(1); i < input; i++ {
		if input%i == 0 {
			sum_of_the_factor += i
		}
	}
	if sum_of_the_factor == input {
		return ClassificationPerfect, nil
	} else if sum_of_the_factor > input {
		return ClassificationAbundant, nil
	} else {
		return ClassificationDeficient, nil
	}
}