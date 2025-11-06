package main

import (
	"fmt"
	"strings"
)

func main(){
	conferenceName:= "GO CONFERENCE"
	const conferenceTickets int =50
	var remainingTickets uint =50
	bookings:=[]string{}

	for{
		fmt.Printf("Welcome to %v booking application!\n",conferenceName)
		fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
		fmt.Println("Get your tickets here to attend")

		var (firstName string
			 lastName string
			 userTickets int
			 email string)

		fmt.Print("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name:")
		fmt.Scan(&lastName)

		isValidName:=len(firstName)>=2 && len(lastName)>=2
		if(!isValidName){
			fmt.Println("First name or last name you entered is too short")
			continue
		}

		fmt.Print("Enter your email:")
		fmt.Scan(&email)

		isValidEmail:=strings.Contains(email,"@") && strings.Contains(email,".")

		if(!isValidEmail){
			fmt.Println("Email address you entered is invalid")
			continue
		}

		fmt.Print("Enter the number of tickets:")
		fmt.Scan(&userTickets)

		isValidTicketNumber:=userTickets>0 && userTickets<=int(remainingTickets)

		if(!isValidTicketNumber){
			fmt.Printf("The number of tickets you entered is invalid. We have %v tickets remaining\n",remainingTickets)
			continue
		}

		remainingTickets-=uint(userTickets)
		bookings=append(bookings,firstName+" "+lastName)

		firstNames:=[]string{}
		for _,booking:=range bookings{
			var firstName=strings.Fields(booking)[0]
			firstNames=append(firstNames,firstName)
		}


		fmt.Println("***************************************************************************************************")
		fmt.Printf("User %v booked %v tickets\n",firstName+" "+lastName,userTickets)
		fmt.Printf("Thankyou %v for booking %v tickets. You will receive a confirmation email at %v shortly\n",firstName,userTickets,email)
		fmt.Printf("The remaining tickets are %v\n",remainingTickets)
		fmt.Println("The first names of bookings are:",firstNames)
		fmt.Println("***************************************************************************************************")

		var noTicketsLeft bool= (remainingTickets==0)
		if noTicketsLeft{
			fmt.Println("Our conference is booked out. Come back next year")
			break
		}
	}
}
