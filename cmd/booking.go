package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

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
	PressEnterToContinue()

	// Implement the logic for making a booking
	booking.RoomType = AC
	for {
		clearTerminal()
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

	clearTerminal()
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
	PressEnterToContinue()
}

func (app *HotelApp) ShowBookingStatus() {
	reader := bufio.NewReader(os.Stdin)

	// Implement the logic for showing status of a booking
	fmt.Println("You selected to show status of a booking.")
	PressEnterToContinue()

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
			PressEnterToContinue()
			return
		}
	}
	fmt.Println("Booking not found!!")
	PressEnterToContinue()
}

// / TODO: complete this
func (app *HotelApp) CancelBooking() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("You selected to cancel a booking.")
	PressEnterToContinue()

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

	found := false
	for id, booking := range app.BookingData {
		if booking.RoomNumber == rn && app.Rooms[rn] {
			// room number is found
			delete(app.BookingData, id)
			found = true
			fmt.Println("Booking cancelled successfully.")
			PressEnterToContinue()
			break
		}
	}

	if !found {
		fmt.Println("Booking not found!!")
	}
	PressEnterToContinue()
}

// / TODO: complete this
func (app *HotelApp) ShowAllBookings() {
	// Implement the logic for showing all bookings
	fmt.Println("You selected to show all bookings.")
	PressEnterToContinue()

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
	PressEnterToContinue()
}
