// response_handler.go

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var api_key = "http://app-homevision-staging.herokuapp.com/api_project/houses/"

// structs for the json response
type EntireResponse struct {
	House []House 	`json:"houses"`
	Ok bool `json:"ok"`
}

// struct for each house, extracting info needed 
type House struct {
	Id int `json:"id"`
	Address string `json:"address"`
	URL string `json:"photoURL"`
}

// function that sends a get request and returns the json data
// takes two parameters, page and per_page, which are the url params for the api
func GetResponse(page int, per_page string) ([]byte) {

	// format the url
	url := api_key + "?page=" + strconv.FormatInt(int64(page), 10) + "&per_page=" + per_page
	var data []byte

	// loop to keep sending get requests to the api in case of falilure (since it's flaky)
	for {
		response, err := http.Get(url)

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		// check if the response is valid(aka statuscode = 200)
		// if true break and return 
		if response.StatusCode == 200 {

			responseData, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}
			// fmt.Println(string(responseData))
			data = responseData
			break
		}
	}

	return data
}

// funciton that parses the json
// takes one parameter data which should be the json data
// returns a map with the house id-address as the key and the image url as the value
func Parse_json(data []byte) (map[string]string) {
	// declare variables
	house_map := make (map[string]string)
	
	var responseObject EntireResponse
	json.Unmarshal(data, &responseObject)	// parse the json

	// loop through all the houses 
	for _, house := range responseObject.House {
		// fmt.Println(house.Id, house.Address, house.URL)
		house_map[fmt.Sprint(house.Id) + "-" + house.Address] = house.URL	// store data into the map 
	}
	return house_map
}