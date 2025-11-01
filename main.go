package main

import "fmt" //format package

func main(){
	conferenceName:= "GO CONFERENCE"
	const conferenceTickets int =50
	var remainingTickets uint =50
	//fmt.Println("Welcome to", conferenceName,"booking application!")
	//fmt.Printf("ConferenceName is of %T type\n",conferenceName)//string
	// fmt.Printf("conferenceTickets is of %T type\n",conferenceTickets)//int
	// fmt.Printf("remainingTickets is of %T type\n",remainingTickets)//uint
	fmt.Printf("Welcome to %v booking application!\n",conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")
	var (userName string; userTickets int)
	userName="Ruhi"
	userTickets=2
	fmt.Printf("User %v booked %v tickets\n",userName,userTickets)
}
