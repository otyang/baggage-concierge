package concierge

import (
	"errors"
	"sync"
)

// BagSize represents the size of a bag which is either (carry-on or checked bag)
type BagSize string

const (
	CarryOn    BagSize = "carry-on"
	CheckedBag BagSize = "checked"
)

// Bag Struct represents a baggage with its size
type Bag struct {
	Size BagSize
}

// Ticket represents a ticket given to a customer when customer's bag is stored
// Same ticket is collected when Bag Retrieval is needed
type Ticket struct {
	ID int
}

// BaggageConcierge represents the interface for the Baggage Concierge (From Question)
type BaggageConcierge interface {
	Store(bag Bag) Ticket
	Retrieve(ticket Ticket) Bag
}

var (
	ErrInvalidTicketID        = errors.New("invalid ticket ID")
	ErrStorageFull            = errors.New("no available storage bins")
	ErrStorageFullCheckedBags = errors.New("no available storage bins for checked bags")
)

// Storage represents the actual implementation of BaggageConcierge Interface
type Storage struct {
	mutex       sync.Mutex
	bins        []BagSize
	nextID      int
	limitOfBins int
}

// NewStorage creates a new Storage instance
func NewStorage(binsLimit int) *Storage {
	return &Storage{
		bins:        make([]BagSize, 0),
		nextID:      1,
		limitOfBins: binsLimit,
	}
}

// Store stores a bag in the Storage and returns a ticket
func (storage *Storage) Store(bag Bag) (Ticket, error) {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()

	if len(storage.bins) >= storage.limitOfBins /*100*/ {
		return Ticket{}, ErrStorageFull
	}

	if bag.Size == CarryOn {
		if len(storage.bins) < storage.limitOfBins-2 /*98*/ { // Check: can slots for 2 X carry-ons bagsize scale?
			storage.bins = append(storage.bins, CarryOn, CarryOn)
		} else {
			storage.bins = append(storage.bins, CarryOn) // Add 1 X carry-on bag-size to fill up the last slot
		}
	} else {
		if len(storage.bins) < storage.limitOfBins-1 /*99*/ { // Check: 1 slot for a checked bag
			storage.bins = append(storage.bins, CheckedBag)
		} else {
			return Ticket{}, ErrStorageFullCheckedBags
		}
	}

	ticket := Ticket{ID: storage.nextID}
	storage.nextID++
	return ticket, nil
}

// Retrieve retrieves a bag from the Storage using the provided ticket
func (storage *Storage) Retrieve(ticket Ticket) (Bag, error) {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()

	if ticket.ID <= 0 || ticket.ID >= storage.nextID {
		return Bag{}, ErrInvalidTicketID
	}

	index := ticket.ID - 1
	if index < 0 || index >= len(storage.bins) {
		return Bag{}, ErrInvalidTicketID
	}

	bagSize := storage.bins[index]
	storage.bins = append(storage.bins[:index], storage.bins[index+1:]...)

	return Bag{Size: bagSize}, nil
}
