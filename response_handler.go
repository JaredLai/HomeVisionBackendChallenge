// response_handler.go

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var api_key = "http://app-homevision-staging.herokuapp.com/api_project/houses/"

func GetResponse(page int, per_page string) {

	url := api_key + "?page=" + strconv.FormatInt(int64(page), 10) + "&per_page=" + per_page

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
			fmt.Println(string(responseData))
			break
		}
	}

}
