package main

import "github.com/google/uuid"

type RoomType int
type RoomPrice float32
type UniqueBookingID uuid.UUID

type MenuItem struct {
	Key   int
	Label string
}

type Booking struct {
	ID           UniqueBookingID
	Name         string
	Age          int
	AadharNumber string
	RoomType     RoomType
	RoomNumber   int
	AmountPaid   float32
}

type HotelApp struct {
	Menu        []MenuItem
	BookingData map[UniqueBookingID]Booking
	Rooms       map[int]bool // Rooms: room_number -> booked/not-booked
	BookedAC    int
	BookedNonAC int
}

func (app *HotelApp) InitializeRooms() {
	for i := FIRST_AC_ROOM; i <= LAST_AC_ROOM; i++ {
		app.Rooms[i] = false
	}
	for i := FIRST_NON_AC_ROOM; i <= LAST_NON_AC_ROOM; i++ {
		app.Rooms[i] = false
	}
}

func NewHotelApp() *HotelApp {
	// Initialize HotelApp with an empty map for BookingData and RoomsData
	return &HotelApp{
		Menu:        []MenuItem{},
		BookingData: make(map[UniqueBookingID]Booking),
		Rooms:       make(map[int]bool),
		BookedAC:    0,
		BookedNonAC: 0,
	}
}
