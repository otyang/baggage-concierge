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

	// Test exceeding storage capacity for checked bags (not more than 100 as stated in task. i.e 100 finite bins)
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

// BenchmarkBaggageStorage benchmarks the BaggageStorage implementation
func BenchmarkBaggageStorage(b *testing.B) {
	concierge := NewStorage(100)

	// Benchmark storing bags
	b.Run("BenchmarkStore", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bag := Bag{Size: CarryOn}
			_, _ = concierge.Store(bag)
		}
	})

	// Benchmark retrieving bags
	b.Run("BenchmarkRetrieve", func(b *testing.B) {
		// Store bags first to retrieve later
		for i := 0; i < b.N; i++ {
			bag := Bag{Size: CarryOn}
			ticket, _ := concierge.Store(bag)
			_ = ticket
		}

		// Now benchmark the retrieval
		b.ResetTimer() // Reset the timer before the retrieval benchmark
		for i := 0; i < b.N; i++ {
			ticket := Ticket{ID: i + 1}
			concierge.Retrieve(ticket)
		}
	})
}
