package airplane

type row struct {
	slots []*slot
	seats []string
}

// assignAisleSeats tries to assign the seats close to the aisle.
func (r *row) assignAisleSeats() (string, bool) {
	var prev *slot
	for i, s := range r.slots {
		if !s.isSeat {
			if prev.assign() {
				return r.seats[i-1], true
			}
		} else if prev != nil {
			if !prev.isSeat {
				if s.assign() {
					return r.seats[i], true
				}
			}
		}
		prev = s
	}
	return "", false
}

// assignWindowSeats assigns seats close to the window.
func (r *row) assignWindowSeats() (string, bool) {
	if r.slots[0].assign() {
		return r.seats[0], true
	}
	last := len(r.slots) - 1
	if r.slots[last].assign() {
		return r.seats[last], true
	}
	return "", false
}

// assignMiddleSeats assigns middle seats.
func (r *row) assignMiddleSeats() (string, bool) {
	for i := 0; i < len(r.slots); i++ {
		if r.slots[i].assign() {
			return r.seats[i], true
		}
	}
	return "", false
}
