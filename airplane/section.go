package airplane

import (
	"strconv"
)

// Section is a division of an airplane. Each section can have its own seaat
// structure divided in blocks.
type Section struct {
	rows   []row
	blocks []int
}

// newSection creates a new section with given seat structure.
func newSection(rows int, blocks []int) *Section {
	s := Section{blocks: blocks}
	for j := 0; j < rows; j++ {
		r := row{}
		count := 0
		for i, n := range blocks {
			for j := 0; j < n; j++ {
				r.slots = append(r.slots, &slot{isSeat: true})
				r.seats = append(r.seats, intToChar(count))
				count++
			}
			if i < len(blocks)-1 {
				r.slots = append(r.slots, &slot{})
				r.seats = append(r.seats, "")
			}
		}
		s.rows = append(s.rows, r)
	}

	return &s
}

// Assign automatically assigns a seat on the current section based on the
// priority rules of first aisle seats, second window seats and last middle
// seats.
func (a *Section) Assign() (string, bool) {
	return assign(a.rows)
}

func assign(rows []row) (string, bool) {
	// Assign aisled seats
	for i, r := range rows {
		if s, ok := r.assignAisleSeats(); ok {
			return strconv.Itoa(i+1) + s, true
		}
	}
	// Assign window seats
	for i, r := range rows {
		if s, ok := r.assignWindowSeats(); ok {
			return strconv.Itoa(i+1) + s, true
		}
	}
	// Assign middleSeats
	for i, r := range rows {
		if s, ok := r.assignMiddleSeats(); ok {
			return strconv.Itoa(i+1) + s, true
		}
	}

	return "", false
}

// Get returns the assigned seat number (e.g., 32C, 44F, etc) given the index
// of the seat (i.e., what is the seat number of the 29th assigned seat?).
func (a *Section) Get(number int) (seat string, ok bool) {
	ss := newSection(len(a.rows), a.blocks)
	for i := 0; i < number; i++ {
		seat, ok = assign(ss.rows)
	}

	return
}
