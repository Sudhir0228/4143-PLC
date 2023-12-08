//Sudhir Ray
//PLC - 4143
//program 4

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	// Create folders for sequential and concurrent downloads
	sequentialFolder := "sequential_images"
	concurrentFolder := "concurrent_images"
	createFolder(sequentialFolder)
	createFolder(concurrentFolder)

	// Image urls for this assignment.
	urls := []string{
		"https://stocksnap.io/photo/sandpiper-bird-H0YXLE9EQP",
		"https://stocksnap.io/photo/home-flowers-RTEYB2HRH0",
		"https://stocksnap.io/photo/flowers-vase-DE2HZJ2YVK",
		"https://stocksnap.io/photo/reading-glasses-2RHWIACZP0",
		"https://stocksnap.io/photo/office-work-BCLRC8HNEO",
	}

	// Sequential download
	startTime := time.Now()
	for _, url := range urls {
		err := downloadImageSequential(url, sequentialFolder)
		if err != nil {
			fmt.Printf("Sequential download error: %v\n", err)
		}
	}
	sequentialDuration := time.Since(startTime)
	fmt.Println("")
	fmt.Printf("Sequential download time: %s\n", sequentialDuration)
	fmt.Println("")

	// Concurrent download
	startTime = time.Now()
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			err := downloadImageConcurrent(u, concurrentFolder)
			if err != nil {
				fmt.Printf("Concurrent download error: %v\n", err)
			}
		}(url)
	}
	wg.Wait()
	concurrentDuration := time.Since(startTime)
	fmt.Println("")
	fmt.Printf("Concurrent download time: %s\n", concurrentDuration)
	fmt.Println("")
}

// Function to download the image sequentially
func downloadImageSequential(url string, folder string) error {
	// Create a new `http.Request` object.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// Create a new `http.Client` object.
	client := &http.Client{}

	// Do the request and get the response.
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// Check the response status code.
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("response status code: %d", resp.StatusCode)
	}

	// Create a unique filename based on the URL with a .jpg extension.
	filename := filepath.Join(folder, "sequential_image_"+extractFilename(url)+".jpg")
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	// Copy the image from the response body to the file.
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		f.Close()
		return err
	}

	// Close the file.
	f.Close()

	// Print a success message.
	fmt.Println("Sequential image saved to", filename)
	return nil
}

// Function to download the image concurrently
func downloadImageConcurrent(url string, folder string) error {
	// Create a new `http.Request` object.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// Create a new `http.Client` object.
	client := &http.Client{}

	// Do the request and get the response.
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// Check the response status code.
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("response status code: %d", resp.StatusCode)
	}

	// Create a unique filename based on the URL with a .jpg extension.
	filename := filepath.Join(folder, "concurrent_image_"+extractFilename(url)+".jpg")
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	// Copy the image from the response body to the file.
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		f.Close()
		return err
	}

	// Close the file.
	f.Close()

	// Print a success message.
	fmt.Println("Concurrent image saved to", filename)
	return nil
}

// Helper function to create a folder if it doesn't exist
func createFolder(folder string) {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.Mkdir(folder, os.ModePerm)
	}
}

// Helper function to extract filename from URL
func extractFilename(url string) string {
	// Use filepath.Base to extract the filename
	return filepath.Base(url)
}
