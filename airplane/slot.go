package airplane

// slot is a phisical space on a specific section of an airplane. It can be
// either an aisle or a seat.
type slot struct {
	isSeat  bool
	occuped bool
}

func (s *slot) assign() bool {
	if !s.isAssignable() {
		return false
	}
	s.occuped = true
	return true
}

func (s *slot) isAssignable() bool {
	if !s.isSeat {
		return false
	}
	if s.occuped {
		return false
	}
	return true
}
