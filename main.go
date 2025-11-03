package main

import "fmt" //format package

func main(){
	conferenceName:= "GO CONFERENCE"
	const conferenceTickets int =50
	var remainingTickets uint =50
	var bookings [50]string

	fmt.Printf("Welcome to %v booking application!\n",conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")

	var (userName string; userTickets int; email string)

	fmt.Print("Enter your name:")
	fmt.Scan(&userName)

	fmt.Print("Enter your email:")
	fmt.Scan(&email)

	fmt.Print("Enter the number of tickets:")
	fmt.Scan(&userTickets)

	remainingTickets-=uint(userTickets)
	bookings[0]=userName

	fmt.Printf("The whole array is %v\n",bookings)
	fmt.Printf("The type of array is %T\n",bookings)
	fmt.Printf("The first value is %v\n",bookings[0])
	fmt.Printf("The type of first value is %T\n",bookings[0])
	fmt.Printf("The second value is %v\n",bookings[1])
	fmt.Printf("The type of second value is %T\n",bookings[1])

	fmt.Println("***************************************************************************************************")
	fmt.Printf("User %v booked %v tickets\n",userName,userTickets)
	fmt.Printf("Thankyou %v for booking %v tickets. You will receive a confirmation email at %v shortly\n",userName,userTickets,email)
	fmt.Printf("The remaining tickets are %v\n",remainingTickets)
}
