package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

/*
	Do this function to get data job list  from api dev.dansmultipro.

	Args:
		page: pagination
		location: job location
		description: description exp (python, ruby)
		full_time: true/false

	Returns:
		Json data job according to the filled filter

*
*/
func Jobs(w http.ResponseWriter, r *http.Request) {
	var client = &http.Client{}
	var data []Job

	query := r.URL.Query()

	queryParams := ""
	for k, v := range query {
		if queryParams == "" {
			queryParams = queryParams + k + "=" + v[0]
		} else {
			queryParams = queryParams + "&" + k + "=" + v[0]
		}
	}

	if queryParams != "" {
		queryParams = "?" + queryParams
	}

	request, err := http.NewRequest("GET", "http://dev3.dansmultipro.co.id/api/recruitment/positions.json"+queryParams, nil)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	response, err := client.Do(request)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(data)
}

/*
	Do this function to get job detail from api dev.dansmultipro.

	Args:
		id: id job

	Returns:
		Json job data detail according to id

*
*/
func JobByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var client = &http.Client{}
	var data Job

	request, err := http.NewRequest("GET", "http://dev3.dansmultipro.co.id/api/recruitment/positions/"+id, nil)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	response, err := client.Do(request)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(data)
}
