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

func errorCheck(err error) {
	if err != nil {
		fmt.Print(err.Error())
	}
}

func save_image(file_name, img_url string) error {

	// first parse the end of the url to get the extention type of the file
	file_url, err := url.Parse(img_url)
	errorCheck(err)

	path := file_url.Path
	splitted := strings.Split(path, ".")

	// modify the file_name
	file_name += "."
	file_name += splitted[len(splitted)-1]

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

	wg.Done()
	return nil
}

func Download_Images(houses map[string]string) {

	for file_name, img_url := range houses {
		wg.Add(1)
		go save_image(file_name, img_url)
	}
	wg.Wait()
}
