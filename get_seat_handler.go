package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// GetSeatHandler assigns a seat to the current request.
func GetSeatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != GET {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	number, ok := r.URL.Query()["number"]
	if !ok {
		http.Error(w, "Invalid input number", http.StatusBadRequest)
		return
	}
	section, ok := r.URL.Query()["section"]
	if !ok {
		output := fmt.Sprintf("Invalid input section")
		http.Error(w, output, http.StatusBadRequest)
		return
	}

	sec, ok := plane.GetSection(section[0])
	if !ok {
		http.Error(w, "the section does not exist", http.StatusBadRequest)
		return
	}

	n, err := strconv.Atoi(number[0])
	if err != nil {
		http.Error(w, "invalid number format", http.StatusBadRequest)
		return
	}

	if seat, ok := sec.Get(n); ok {
		output := fmt.Sprintf(`"%s"`, seat)
		fmt.Fprint(w, output)
		return
	}

	http.Error(w, "this section is full", http.StatusBadRequest)
}
