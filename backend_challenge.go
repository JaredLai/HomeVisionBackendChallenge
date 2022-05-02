// backend_challenge.go
package main

import (
	"os"
	"fmt"
)

func main() {

	// delete the current ouput directory 
	err := os.RemoveAll("output")
	if err != nil {
		fmt.Print(err.Error())
	}

	// call the API ten times, for pages 1-10
	// assuming the default per_page as 10; can be easily changed below
	per_page := "10"
	for page :=1 ; page<= 10; page++ {
		result := GetResponse(page, per_page)// parses the response; returns the json result
		houses := Parse_json(result)			// extracts the id, address, url from the json and saves it into a map [id-address]image_url
		Download_Images(houses)					// downloads images into output folder 
	}

	
}
