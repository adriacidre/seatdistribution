package airplane

// Airplane airplane seat map.
type Airplane struct {
	sections map[string]*Section
}

// New creates a new airplane
func New() *Airplane {
	return &Airplane{
		sections: make(map[string]*Section),
	}
}

// AddSection adds a section to an airplane with a given number of rows, and
// seat distribution in blocks.
// To provide seat distribution in blocks we can use a format like [3,4,3] which
// means the rows will be divided on 3 blocks respectively with 3, 4 and 3 seats
// each.
func (a *Airplane) AddSection(id string, rows int, blocks []int) *Section {
	a.sections[id] = newSection(rows, blocks)
	return a.sections[id]
}

// GetSection gets a previously stored section.
func (a *Airplane) GetSection(id string) (*Section, bool) {
	sec, ok := a.sections[id]
	return sec, ok
}
