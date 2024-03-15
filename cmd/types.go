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
