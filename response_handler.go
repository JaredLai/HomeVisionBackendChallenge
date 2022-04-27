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

type EntireResponse struct {
	House []House 	`json:"houses"`
	Ok bool `json:"ok"`
}

type House struct {
	Id int `json:"id"`
	Address string `json:"address"`
	URL string `json:"photoURL"`
}

func GetResponse(page int, per_page string) ([]byte) {

	url := api_key + "?page=" + strconv.FormatInt(int64(page), 10) + "&per_page=" + per_page
	var data []byte

	for {
		response, err := http.Get(url)

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
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

func Parse_json(data []byte) (map[string]string) {
	house_map := make (map[string]string)
	
	var responseObject EntireResponse
	json.Unmarshal(data, &responseObject)

	for _, house := range responseObject.House {
		// fmt.Println(house.Id, house.Address, house.URL)
		house_map[fmt.Sprint(house.Id) + "-" + house.Address] = house.URL
	}
	return house_map
}