package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func NewHotelApp() *HotelApp {
	return &HotelApp{
		Menu:        []MenuItem{},
		BookingData: make(map[UniqueBookingID]Booking),
		Rooms:       make(map[int]bool),
		BookedAC:    0,
		BookedNonAC: 0,
	}
}

func (app *HotelApp) Welcome() {
	fmt.Println("Hotel Management System CLI App")
	fmt.Println("Copyright@2024 vickvey")
	PressEnterToContinue()
}

func (app *HotelApp) SetMenu() {
	app.Menu = []MenuItem{
		{1, "to make a booking."},
		{2, "to show status of a booking."},
		{3, "to cancel a booking."},
		{4, "to show all bookings."},
		{0, "exit menu."},
	}
}

func (app *HotelApp) InitializeRooms() {
	for i := FIRST_AC_ROOM; i <= LAST_AC_ROOM; i++ {
		app.Rooms[i] = false
	}
	for i := FIRST_NON_AC_ROOM; i <= LAST_NON_AC_ROOM; i++ {
		app.Rooms[i] = false
	}
}

func (app *HotelApp) CheckoutMessage() {
	fmt.Println("Thank you for using this App!")
	fmt.Println("GoodBye.")
}

func (app *HotelApp) PrintMenu() {
	fmt.Println("Menu Page")
	fmt.Println("--------------------------------------")
	for _, item := range app.Menu {
		fmt.Printf("Press %d: %s\n", item.Key, item.Label)
	}
	fmt.Println("--------------------------------------")
}

func (app *HotelApp) Run() {
	app.Welcome()
	reader := bufio.NewReader(os.Stdin)

	for {
		app.PrintMenu()

		fmt.Print("Enter your choice: ")
		choiceStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// Extract the integer choice from the input string
		choice, err := strconv.Atoi(strings.TrimSpace(choiceStr))
		if err != nil {
			fmt.Println("Invalid choice:", choiceStr)
			PressEnterToContinue()
			continue
		}

		// Handle the user's choice
		switch choice {
		case 1:
			app.MakeBooking()
		case 2:
			app.ShowBookingStatus()
		case 3:
			app.CancelBooking()
		case 4:
			app.ShowAllBookings()
		case 0:
			fmt.Println("Exiting the menu.")
			PressEnterToContinue()
			return
		default:
			fmt.Println("Invalid choice:", choice)
		}
	}
}
