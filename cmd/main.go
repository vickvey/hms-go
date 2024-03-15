package main

func main() {
	// Initialize the hotel app and load the required things
	myHotel := NewHotelApp()
	myHotel.SetMenu()

	// Run the hotel app
	myHotel.Run()

	// Exiting message
	myHotel.CheckoutMessage()
}
