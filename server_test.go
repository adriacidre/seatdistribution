package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adriacidre/seatdistribution/airplane"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/sections/assign", AssignHandler).Methods("POST")
	router.HandleFunc("/sections", AddSectionHandler).Methods("POST")
	router.HandleFunc("/sections/seat", GetSeatHandler).Methods("GET")
	return router
}

func TestAddSectionEndpoint(t *testing.T) {
	plane = airplane.New()
	var flagtests = []struct {
		name   string
		input  *AddSectionInput
		output string
		code   int
	}{
		{
			name: "success",
			input: &AddSectionInput{
				ID:     "id",
				Rows:   1,
				Blocks: []int{3, 4, 3},
			},
			output: `{"success"}`,
			code:   200,
		},
		{
			name: "id not provided",
			input: &AddSectionInput{
				Rows:   1,
				Blocks: []int{3, 4, 3},
			},
			output: "You must provide an ID for the section\n",
			code:   400,
		},
		{
			name: "empty rows",
			input: &AddSectionInput{
				ID:     "id",
				Blocks: []int{3, 4, 3},
			},
			output: "You must provide a number of rows for the section\n",
			code:   400,
		},
		{
			name: "empty blocks",
			input: &AddSectionInput{
				ID:   "id",
				Rows: 1,
			},
			output: "You must define the blocks structure\n",
			code:   400,
		},
	}
	for _, tt := range flagtests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.input)
			request, _ := http.NewRequest("POST", "/sections", bytes.NewBuffer(body))
			response := httptest.NewRecorder()
			Router().ServeHTTP(response, request)
			assert.Equal(t, tt.output, response.Body.String())
			assert.Equal(t, tt.code, response.Code)
		})
	}
}

func TestAssignEndpoint(t *testing.T) {
	plane = airplane.New()
	plane.AddSection("id", 1, []int{3, 4, 3})
	var flagtests = []struct {
		name   string
		input  *AssignInput
		output string
		code   int
	}{
		{
			name: "success",
			input: &AssignInput{
				ID: "id",
			},
			output: `"1C"`,
			code:   200,
		},
		{
			name:   "not section provided",
			input:  &AssignInput{},
			output: "Invalid input id\n",
			code:   400,
		},
	}
	for _, tt := range flagtests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.input)
			request, _ := http.NewRequest("POST", "/sections/assign", bytes.NewBuffer(body))
			response := httptest.NewRecorder()
			Router().ServeHTTP(response, request)
			assert.Equal(t, tt.output, response.Body.String())
			assert.Equal(t, tt.code, response.Code)
		})
	}
}

func TestAssignUnconfiguredSection(t *testing.T) {
	plane = airplane.New()
	body, _ := json.Marshal(&AssignInput{
		ID: "id",
	})
	request, _ := http.NewRequest("POST", "/sections/assign", bytes.NewBuffer(body))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, "the section does not exist\n", response.Body.String())
	assert.Equal(t, 400, response.Code)
}

func TestGetSeatEndpoint(t *testing.T) {
	plane = airplane.New()
	plane.AddSection("id", 2, []int{3, 4, 3})
	var flagtests = []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "success",
			input:  "10",
			output: `"1J"`,
		},
		{
			name:   "not section provided",
			input:  "11",
			output: `"2A"`,
		},
	}
	for _, tt := range flagtests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/sections/seat", nil)
			q := req.URL.Query()
			q.Add("section", "id")
			q.Add("number", tt.input)
			req.URL.RawQuery = q.Encode()

			response := httptest.NewRecorder()
			Router().ServeHTTP(response, req)
			assert.Equal(t, tt.output, response.Body.String())
		})
	}
}
