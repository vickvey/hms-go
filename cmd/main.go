package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

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
			app.PressEnterToContinue()
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
			app.PressEnterToContinue()
			return
		default:
			fmt.Println("Invalid choice:", choice)
		}
	}
}

func (app *HotelApp) getFreshACRoomNumber() (int, error) {
	for i := FIRST_AC_ROOM; i <= LAST_AC_ROOM; i++ {
		if !app.Rooms[i] {
			app.Rooms[i] = true
			return i, nil
		}
	}
	return 0, fmt.Errorf("all ac rooms are full")
}

func (app *HotelApp) getFreshNONACRoomNumber() (int, error) {
	for i := FIRST_NON_AC_ROOM; i <= LAST_NON_AC_ROOM; i++ {
		if !app.Rooms[i] {
			app.Rooms[i] = true
			return i, nil
		}
	}
	return 0, fmt.Errorf("all non-ac rooms are full")
}

func (app *HotelApp) MakeBooking() {
	booking := Booking{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("You selected to make a booking.")
	app.PressEnterToContinue()

	// Implement the logic for making a booking
	booking.RoomType = AC
	for {
		app.clearTerminal()
		fmt.Println("Booking Page")
		fmt.Println("--------------------------------------")
		fmt.Println("What type of room do you want?")
		fmt.Println("Press 0: for AC, Cost: Rs ", AC_PRICE)
		fmt.Println("Press 1: for Non-AC, Cost: Rs ", NON_AC_PRICE)
		fmt.Println("--------------------------------------")

		fmt.Print("Your choice: ")
		choiceStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		// Extract the integer choice from the input string
		choice, err := strconv.Atoi(strings.TrimSpace(choiceStr))
		if err != nil {
			fmt.Println("Invalid choice:", choiceStr)
			continue
		}

		if choice == 0 {
			booking.RoomType = AC
			break
		} else if choice == 1 {
			booking.RoomType = NONAC
			break
		}
	}

	if booking.RoomType == AC {
		if app.BookedAC >= TOTAL_AC_ROOMS {
			fmt.Println("Sorry! All AC rooms are full!")
			return
		}
		fmt.Printf("Your room type choice is: %s\n", "AC")
	} else {
		if app.BookedNonAC >= TOTAL_NON_AC_ROOMS {
			fmt.Println("Sorry! All NON-AC rooms are full!")
			return
		}
		fmt.Printf("Your room type choice is: %s\n", "NON-AC")
	}

	// Get Unique Booking ID
	booking.ID = UniqueBookingID(uuid.New())

	fmt.Printf("Enter your name: ")
	name_str, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	booking.Name = strings.TrimRight(name_str, "\n")

	fmt.Printf("Enter your age: ")
	age_str, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	age, err := strconv.Atoi(strings.TrimSpace(age_str))
	if err != nil {
		panic(err)
	}
	if age < 0 || age >= 150 {
		fmt.Println("Age cannot be less than 0 or more than 150 years.")
		panic("Invalid age")
	}
	booking.Age = age

	fmt.Printf("Enter your aadhar number: ")
	aadhar_str, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	booking.AadharNumber = strings.TrimRight(aadhar_str, "\n")

	app.clearTerminal()
	fmt.Printf("Your booking details are: \n\n")
	fmt.Printf("Name: %v\n", booking.Name)
	fmt.Printf("Age: %v\n", booking.Age)
	fmt.Printf("Aadhar Number: %v\n", booking.AadharNumber)
	if booking.RoomType == AC {
		fmt.Println("Room Type: AC")
	} else {
		fmt.Println("Room Type: NON-AC")
	}

	/// TODO: yet to add the "are these details correct section.."

	fmt.Println("\nBooking done.")

	if booking.RoomType == AC {
		booking.RoomNumber, err = app.getFreshACRoomNumber()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Your alloted room number is: ", booking.RoomNumber)
		fmt.Println("Please deposit amount Rs ", AC_PRICE)
		app.BookedAC++
		booking.AmountPaid = float32(AC_PRICE)
	} else {
		booking.RoomNumber, err = app.getFreshNONACRoomNumber()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Your alloted room number is: ", booking.RoomNumber)
		fmt.Println("Please deposit amount Rs ", NON_AC_PRICE)
		app.BookedNonAC++
		booking.AmountPaid = float32(NON_AC_PRICE)
	}

	// saving the booking details in the app
	app.BookingData[booking.ID] = booking
	app.PressEnterToContinue()
}

func (app *HotelApp) ShowBookingStatus() {
	reader := bufio.NewReader(os.Stdin)

	// Implement the logic for showing status of a booking
	fmt.Println("You selected to show status of a booking.")
	app.PressEnterToContinue()

	fmt.Println("Booking Status Page")
	fmt.Println("---------------------------------------")
	fmt.Printf("Enter your room number: ")
	rn_str, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	rn, err := strconv.Atoi(strings.TrimSpace(rn_str))
	if err != nil {
		panic(err)
	}

	for _, booking := range app.BookingData {
		if booking.RoomNumber == rn && app.Rooms[rn] {
			fmt.Println("Name: ", booking.Name)
			fmt.Println("Aadhar Number: ", booking.AadharNumber)
			if booking.RoomType == AC {
				fmt.Println("Room Type: AC")
			} else {
				fmt.Println("Room Type: NON-AC")
			}
			fmt.Println("Booking amount paid: Rs ", booking.AmountPaid)
			fmt.Println("---------------------------------------")
			app.PressEnterToContinue()
			return
		}
	}
	fmt.Println("Booking not found!!")
	app.PressEnterToContinue()
}

// / TODO: complete this
func (app *HotelApp) CancelBooking() {
	reader := bufio.NewReader(os.Stdin)

	// Implement the logic for showing status of a booking
	fmt.Println("You selected to cancel a booking.")
	app.PressEnterToContinue()

	fmt.Println("Booking Cancellation Page")
	fmt.Println("---------------------------------------")
	fmt.Printf("Enter your room number: ")
	rn_str, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	rn, err := strconv.Atoi(strings.TrimSpace(rn_str))
	if err != nil {
		panic(err)
	}

	for _, booking := range app.BookingData {
		if booking.RoomNumber == rn && app.Rooms[rn] {
			// room number is found
			delete(app.BookingData, booking.ID)

			fmt.Println("Booking cancelled successfully.")
			app.PressEnterToContinue()
			return
		}
	}
	fmt.Println("Booking not found!!")
	app.PressEnterToContinue()
}

// / TODO: complete this
func (app *HotelApp) ShowAllBookings() {
	// Implement the logic for showing all bookings
	fmt.Println("You selected to show all bookings.")
	app.PressEnterToContinue()

	fmt.Println("Booking List Page")
	fmt.Println("-----------------------------------------------")

	if len(app.BookingData) == 0 {
		fmt.Println("No Booking data to show!!")
		return
	}

	for _, booking := range app.BookingData {
		fmt.Println("-----------------------------------------------")
		fmt.Println("Room number: ", booking.RoomNumber)
		fmt.Println("Name of guest: ", booking.Name)
		fmt.Println("Aadhar number of guest: ", booking.AadharNumber)
		if booking.RoomType == AC {
			fmt.Println("Room Type: AC")
		} else {
			fmt.Println("Room Type: NON-AC")
		}
		fmt.Println("Booking amount paid: Rs ", booking.AmountPaid)
	}
	fmt.Println("-----------------------------------------------")
	app.PressEnterToContinue()
}

func main() {
	// Initialize the hotel app
	myHotel := NewHotelApp()
	myHotel.SetMenu()

	myHotel.Run()

	// Exiting message
	myHotel.CheckoutMessage()
}

func (app *HotelApp) clearTerminal() {
	// Check the operating system
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// For Unix-like systems, use ANSI escape codes
		fmt.Print("\033[H\033[2J")
	}
}

// Also clears the terminal buffer
func (app *HotelApp) PressEnterToContinue() {
	fmt.Print("Press Enter to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	app.clearTerminal()
}

func (app *HotelApp) CheckoutMessage() {
	fmt.Println("Thank you for using this App!")
	fmt.Println("GoodBye.")
}

func (app *HotelApp) Welcome() {
	fmt.Println("Hotel Management System CLI App")
	fmt.Println("Copyright@2024 vickvey")
	app.PressEnterToContinue()
}

func (app *HotelApp) SetMenu() {
	// Initialize the Menu field with menu items
	app.Menu = []MenuItem{
		{1, "to make a booking."},
		{2, "to show status of a booking."},
		{3, "to cancel a booking."},
		{4, "to show all bookings."},
		{0, "exit menu."},
	}
}

func (app *HotelApp) PrintMenu() {
	fmt.Println("Menu Page")
	fmt.Println("--------------------------------------")
	for _, item := range app.Menu {
		fmt.Printf("Press %d: %s\n", item.Key, item.Label)
	}
	fmt.Println("--------------------------------------")
}
