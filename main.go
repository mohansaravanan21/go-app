package main

import (
	"fmt"
	"go-app/helper"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

// conferenceName := "Go Conference" -> only for variables not const or explicit var definition like uint
// var conferenceName string= "Go Conference"
const conferenceTickets = 50

var remainingTickets uint = conferenceTickets

type UserData struct { // holds diffrent data types of data in key value pair.
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// var bookings []string // only one data type
var bookings = make([]UserData, 0)

// key value pairs
// var data  map[string]string // only one data type
var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	// var remainingTickets uint = conferenceTickets - 10
	// remainingTickets = -1
	// fmt.Println(remainingTickets)
	// fmt.Println(&remainingTickets)

	// for true {
	for {
		// get user inputs
		firstName, lastName, email, userTickets := getUserInputs()
		// for remainingTickets > 0 && len(bookings) > 0 {

		// User Input Validations
		// var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
		isValidName, isValidEmail, isValidTicketCount := helper.ValidateInputs(firstName, lastName, email, userTickets, remainingTickets)
		// isValidCity :=  city == "Bangalore" || city == "Bengaluru"

		if isValidEmail && isValidName && isValidTicketCount {

			bookTickets(firstName, lastName, userTickets, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			// userName = "Don"
			// userTickets = 2
			// fmt.Printf("The whole bookings array: %v\n", bookings)
			// fmt.Printf("The First value: %v\n", bookings[0])
			// fmt.Printf("The Slice Type: %T\n", bookings)
			// fmt.Printf("The Slice Length: %v\n", len(bookings))
			// fmt.Println(bookings)

			// call func to print first names
			// var firstNames []string = getFirstNames(bookings)
			firstNames := getFirstNames()
			fmt.Printf("These are firstnames of all the bookings done: %v\n", firstNames)

			// fmt.Printf("These are all the bookings done: %v\n", bookings)

			if remainingTickets == 0 {
				// end process
				fmt.Println("House Full!!!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your Name is too short")
			}
			if !isValidEmail {
				fmt.Println("Your email is not valid")
			}
			if !isValidTicketCount {
				fmt.Println("Your Ticket number is invalid")
			}
			// fmt.Println("Your Input is Invalid")
			// fmt.Printf("We only have %v tickets remaining\n", remainingTickets)
		}

	}
	wg.Wait()

	// city := "Bangalore"

	// switch city {
	// case "Bangalore", "Bengaluru": // execute code
	// case "Chennai", "Madras":
	// case "Coimbatore":
	// default: // nothing matched
	// }
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking app!\n", conferenceName)
	fmt.Println("Get your tickets soon")
	fmt.Printf("conferenceTickets is a type of: %T, conferenceName is type of: %T\n", conferenceTickets, conferenceName)
	fmt.Printf("We have %v available tickets and %v tickets left for the event!\n", conferenceTickets, remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, value := range bookings {
		var names = value.firstName // explore this!!!!
		firstNames = append(firstNames, names)
	}
	return firstNames
	// fmt.Printf("These are firstnames of all the bookings done: %v\n", firstNames)
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name

	// get user input
	fmt.Println("Enter Your First Name: ")
	fmt.Scan(&firstName) // pointers for addres ref. also called special variable sin Go

	fmt.Println("Enter Your Last Name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter email: ")
	fmt.Scan(&email)

	fmt.Println("Enter No. of tickets you want: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(firstName string, lastName string, userTickets uint, email string) {
	// var userData = make(map[string]string)
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) // explore

	bookings = append(bookings, userData) // slice - a dynamic array

	remainingTickets = remainingTickets - userTickets
	fmt.Printf("List of Bookimgs is %v\n", bookings)
	fmt.Printf("The user %v booked %v tickets today.\n", firstName, userTickets)
	fmt.Printf("Thank you %v for your purchase. Only %v tickets are remaining now! Hurry Up!!!\n", firstName, remainingTickets)
	// fmt.Printf("Thank you %v for your purchase. Only %v tickets are remaining now! Hurry Up!!!\n", firstName, remainingTickets)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v Tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##########################################################")
	fmt.Printf("Sending Ticket:\n %v \nto email:\n %v\n", ticket, email)
	fmt.Println("##########################################################")
	wg.Done()
}
