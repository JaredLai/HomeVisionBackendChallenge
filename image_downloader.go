// image_downloader.go

package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

// helper funciton to check for errors. 
func errorCheck(err error) {
	if err != nil {
		fmt.Print(err.Error())
	}
}

// saves image to the output folder
// takes two inputs, file_name and the url for the image
// downloads iamge to the output folder, creates one if folder doesn't already exist
func save_image(file_name, img_url string) error {

	// first parse the end of the url to get the extention type of the file
	file_url, err := url.Parse(img_url)
	errorCheck(err)

	path := file_url.Path
	splitted := strings.Split(path, ".")

	// modify the file_name
	file_name += "."
	file_name += splitted[len(splitted)-1]

	// get the image from the given url
	response, err := http.Get(img_url)
	errorCheck(err)

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Error when downloading file" + file_name)
	}

	// create output directory
	if _, err := os.Stat("/output/"); os.IsNotExist(err) {
		os.Mkdir("output", 0755) // Create your file
	}

	// d := []byte("")
	// errorCheck(os.WriteFile("/output/"+file_name, response.Body, 0644))

	// first create an empty file
	file, err := os.Create("output/"+file_name)
	if err != nil {
		return err
	}

	defer file.Close()

	// write data that was downloaded
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	wg.Done()	// tell wait group this goroutine is finished
	return nil
}

// downloads images concurrently 
// given a map of houses, loops througha and downloads images
// does not return a value 
func Download_Images(houses map[string]string) {

	// looping through the elements in the map
	for file_name, img_url := range houses {
		wg.Add(1)	// adds to the wait group
		go save_image(file_name, img_url)	// calls save_image concurrently 
	}
	wg.Wait()	// waits for all goroutines to finish 
}
