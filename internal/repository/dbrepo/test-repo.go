package dbrepo

import (
	"errors"
	"time"

	"github.com/andres15mol/bookings/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}


//InsertReservation insert a reservation inn the database
func (m *testDBRepo) InsertResevation(res models.Reservation) (int, error) {
	
	return 1, nil
}

// InsertRoomRestriction inserts a room restriction in to the database
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	
	return nil
}
//SearchAvailabilityByDates returns true if availability exist for roomID, and false if no availability exist
func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error){
	

	return false, nil
}

//SearchAvailabilityForAllRooms return a slice of available rooms, if any, for given date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error){
	
	var rooms []models.Room

	
	return rooms, nil

}

//GetRoomByID gets a room by id
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error){
	var room models.Room
	if id > 2 {
		return room, errors.New("Some error")
	}
	


	return room, nil
}