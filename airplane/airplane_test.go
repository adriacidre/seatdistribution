package airplane

import "testing"

func TestAssignSeat(t *testing.T) {
	expectations := []string{
		// First aisle seats
		"1C", "1D", "1G", "1H",
		"2C", "2D", "2G", "2H",
		// Window seats
		"1A", "1J",
		"2A", "2J",
		// Middle seats
		"1B", "1E", "1F", "1I",
		"2B", "2E", "2F", "2I",
	}
	a := New()
	sec := a.AddSection("id", 2, []int{3, 4, 3})

	for _, e := range expectations {
		seat, ok := sec.Assign()
		if !ok {
			t.Errorf("Not ok")
		}
		if seat != e {
			t.Errorf("Got %s but was expecting %s", seat, e)
		}
	}

	// Once all seats are full, if I try to assign a new slot, i should get a
	if _, ok := sec.Assign(); ok {
		t.Error("Expecting seat assignement to be unsuccessful but it was")
	}
}

func TestSections(t *testing.T) {
	id := "id"
	a := New()
	_, ok := a.GetSection(id)
	if ok {
		t.Fatalf("Expecting non success when getting a non existing section")
	}
	a.AddSection(id, 3, []int{3, 4, 3})
	sec, ok := a.GetSection(id)
	if !ok {
		t.Fatalf("Expecting success when getting an existing section")
	}
	if sec == nil {
		t.Fatal("Expecting to retrieve a valid section")
	}
}

func TestGet(t *testing.T) {
	expectations := map[int]string{
		15: "1F",
		16: "1I",
		17: "2B",
		7:  "2G",
		8:  "2H",
		9:  "1A",
		1:  "1C",
		2:  "1D",
		3:  "1G",
		4:  "1H",
		12: "2J",
		13: "1B",
		5:  "2C",
		6:  "2D",
		10: "1J",
		11: "2A",
		14: "1E",
		18: "2E",
		19: "2F",
		20: "2I",
	}
	a := New()
	sec := a.AddSection("id", 2, []int{3, 4, 3})

	for in, out := range expectations {
		seat, ok := sec.Get(in)
		if !ok {
			t.Errorf("Not ok")
		}
		if seat != out {
			t.Errorf("Got %s but was expecting %s", seat, out)
		}
	}

}
