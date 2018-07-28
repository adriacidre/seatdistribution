package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// AddSectionInput input data to be received on AddSectionHandler.
type AddSectionInput struct {
	ID     string `json:"id"`
	Rows   int    `json:"rows"`
	Blocks []int  `json:"blocks"`
}

func (s *AddSectionInput) validate() error {
	if len(s.ID) == 0 {
		return errors.New("You must provide an ID for the section")
	}
	if s.Rows == 0 {
		return errors.New("You must provide a number of rows for the section")
	}
	if len(s.Blocks) == 0 {
		return errors.New("You must define the blocks structure")
	}
	return nil
}

// AddSectionHandler adds a section to the airplane on this server.
func AddSectionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != POST {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}

	var input AddSectionInput
	if err := json.Unmarshal(body, &input); err != nil {
		output := fmt.Sprintf("Invalid input %v", err)
		http.Error(w, output, http.StatusBadRequest)
		return
	}
	if err := input.validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	plane.AddSection(input.ID, input.Rows, input.Blocks)

	fmt.Fprint(w, `{"success"}`)
}
