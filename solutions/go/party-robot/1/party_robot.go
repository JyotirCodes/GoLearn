package partyrobot
import "fmt"
// Welcome greets a person by name.
func Welcome(name string) string {
    welcomeMessage:= fmt.Sprintf("Welcome to my party, %s!", name)
	return welcomeMessage
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	birthdayMessage:= fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
    return birthdayMessage
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
    welcomeMessage:= Welcome(name)
    tableDirections:= fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.", table, direction, distance)
    seatPartner:= fmt.Sprintf("You will be sitting next to %s.", neighbor)
	finalMessage:= fmt.Sprintf("%s\n%s\n%s", welcomeMessage, tableDirections, seatPartner)
    return finalMessage
}
