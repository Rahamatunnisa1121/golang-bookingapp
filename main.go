package main

import (
	"fmt"
	"strconv"
	"strings"
)

//in global scope, we cannot use := operator
const conferenceTickets int = 50
const conferenceName string = "GO CONFERENCE"
var remainingTickets uint =50
var bookings =make([]map[string]string,0)//as it is a slice of maps
func main(){
	greetUser()
	for{
		printAvailableTickets()

		firstName,lastName,email,userTickets:=getUserInput()
		isValidName, isValidEmail, isValidTicketNumber:=validateUserInput(firstName,lastName,email,userTickets)

		if(isValidName && isValidEmail && isValidTicketNumber){
			bookTickets(userTickets,firstName,lastName,email)
			firstNames:=getFirstNames()
			fmt.Println("The first names of bookings are:",firstNames)
			if remainingTickets==0{
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
		}else if(!isValidName){
			fmt.Println("First name or last name you entered is too short")
			continue
		}else if(!isValidEmail){
			fmt.Println("Email address you entered is invalid")
			continue	
		}else if(!isValidTicketNumber){
			fmt.Println("Number of tickets you entered is invalid")
			continue
		}
	}
}
func greetUser(){
	fmt.Printf("Welcome to %v booking application!\n",conferenceName)
}
func printAvailableTickets(){
	fmt.Printf("We have a total of %v tickets and %v are still available\n",conferenceTickets,remainingTickets)
}
func getUserInput()(string,string,string,int){
	var firstName string
	var lastName string
	var userTickets int
	var email string	
	fmt.Print("Enter your first name:")
	fmt.Scan(&firstName)		
	fmt.Print("Enter your last name:")
	fmt.Scan(&lastName)		
	fmt.Print("Enter your email:")
	fmt.Scan(&email)		
	fmt.Print("Enter the number of tickets:")
	fmt.Scan(&userTickets)
	return firstName,lastName,email,userTickets
}
func validateUserInput(firstName string,lastName string,email string,userTickets int)(bool,bool,bool){
	isValidName:= len(firstName)>=2 && len(lastName)>=2
	isValidEmail:= strings.Contains(email,"@") && strings.Contains(email,".")
	isValidTicketNumber:= userTickets>0 && userTickets<=int(remainingTickets)
	return isValidName,isValidEmail,isValidTicketNumber
}
func bookTickets(userTickets int,firstName string,lastName string,email string){
	//update remaining tickets
	remainingTickets-=uint(userTickets)
	var userData=make(map[string]string)
	userData["firstName"]=firstName
	userData["lastName"]=lastName
	userData["email"]=email
	userData["userTickets"]=strconv.FormatUint(uint64(userTickets),10)
	bookings=append(bookings,userData)
	fmt.Printf("Thankyou %v for booking %v tickets\n",firstName+" "+lastName,userTickets)
	fmt.Printf("You will receive a confirmation email at %v shortly\n",email)
	fmt.Printf("The remaining tickets are %v\n",remainingTickets)
	fmt.Printf("The bookings are %v\n",bookings)
}

func getFirstNames() []string{
	firstNames:=[]string{}
		for _,booking:=range bookings{
			var firstName=booking["firstName"]
			firstNames=append(firstNames,firstName)
		}
		return firstNames
}