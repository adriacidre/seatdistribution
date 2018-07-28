package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// POST post method.
const POST = "POST"

// GET post method.
const GET = "GET"

// AssignInput input data to be received on AssignHandler.
type AssignInput struct {
	ID string `json:"ID"`
}

// AssignHandler assigns a seat to the current request.
func AssignHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != POST {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}

	var input AssignInput
	if err := json.Unmarshal(body, &input); err != nil {
		output := fmt.Sprintf("Invalid input %v", err)
		http.Error(w, output, http.StatusBadRequest)
		return
	}

	if len(input.ID) == 0 {
		http.Error(w, "Invalid input id", http.StatusBadRequest)
		return
	}

	sec, ok := plane.GetSection(input.ID)
	if !ok {
		http.Error(w, "the section does not exist", http.StatusBadRequest)
		return
	}

	if seat, ok := sec.Assign(); ok {
		output := fmt.Sprintf(`"%s"`, seat)
		fmt.Fprint(w, output)
		return
	}

	http.Error(w, "this section is full", http.StatusBadRequest)
}
