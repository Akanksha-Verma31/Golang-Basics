package main
import "fmt"
func Welcome(name string) (message string) {
	message = fmt.Sprintf("welcome to my party, %v!", name)
	return
}
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy Birthday %v, You are %v years old", name, age)
}
func AssignTable(
	name string,
	tableNumber int,
	seatingPartner string,
	direction string,
	distance float64,
) string {
	return fmt.Sprintf(
		"Welcome to my party, %v! \nYou have been assigned to table %03d. Your table is on the %v, exactly %.1f meters from here \nYou will be sitting next to %v",
		name,
		tableNumber,
		direction,
		distance,
		seatingPartner,
	)
	// %d int
	// %s string
	// %f float
	// %v generic

}