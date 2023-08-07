package concierge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage_Store(t *testing.T) {
	concierge := NewStorage(100)

	// Test storing and retrieving carry-on bags
	ticket, err := concierge.Store(Bag{Size: CarryOn})
	assert.NoError(t, err)

	//
	retrievedBag, err := concierge.Retrieve(ticket)
	assert.NoError(t, err)
	assert.Equal(t, CarryOn, retrievedBag.Size)

	//

	// Test storing and retrieving checked bags
	_, err = concierge.Store(Bag{Size: CheckedBag})
	assert.NoError(t, err)

	// Test exceeding storage capacity for checked bags
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic as expected when exceeding checked bag capacity")
		}
	}()
	for i := 0; i < 99; i++ {
		_, err = concierge.Store(Bag{Size: CheckedBag})
		if err != nil {
			panic(err)
		}
	}
}

func TestStorage_Retrieve(t *testing.T) {
	concierge := NewStorage(100)

	// Test storing and retrieving carry-on bags
	ticket, err := concierge.Store(Bag{Size: CarryOn})
	assert.NoError(t, err)

	//
	retrievedBag, err := concierge.Retrieve(ticket)
	assert.NoError(t, err)
	assert.Equal(t, CarryOn, retrievedBag.Size)

	// Test storing and retrieving checked bags
	ticket, err = concierge.Store(Bag{Size: CheckedBag})
	assert.NoError(t, err)

	retrievedBag, err = concierge.Retrieve(ticket)
	assert.NoError(t, err)
	assert.Equal(t, CheckedBag, retrievedBag.Size)

	// Wrong retrieval
	retrievedBag, err = concierge.Retrieve(Ticket{ID: 10000})
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidTicketID, err)
}
