//Sudhir Ray
//PLC - 4143
//program 4

package main

import (
	"fmt"
	"github.com/Sudhir0228/mymodules"
	"time"
)

func downloadImagesSequential(urls []string) {
	for i, imageURL := range urls {
		//save image under the given name and print it out
		fileName := fmt.Sprintf("downloaded_image_%d.jpg", i)
		err := mymodules.GetImage(imageURL, fileName)
		//print out result of download with its url
		if err != nil {
			fmt.Printf("Error downloading %s: %v\n", imageURL, err)
		} else {
			fmt.Printf("Downloaded: %s\n", imageURL)
		}
	}
}

// struct to hold the result of each download
type downloadResult struct {
	url string
	err error
}

// Concurrent version of the image downloader.
func downloadImagesConcurrent(urls []string) {
	//channels are used to communicate between goroutines
	resultCh := make(chan downloadResult)
	//save image under the given name and print it out
	for i, imageURL := range urls {
		fileName := fmt.Sprintf("downloaded_image_%d.jpg", i)
		go func(url, file string) {
			err := mymodules.GetImage(url, file)
			resultCh <- downloadResult{url, err}
		}(imageURL, fileName)

	}
	//print out result of download with its url
	for range urls {
		result := <-resultCh
		if result.err != nil {
			fmt.Printf("Error downloading %s: %v\n", result.url, result.err)
		} else {
			fmt.Printf("Downloaded: %s\n", result.url)
		}
	}
}

func main() {
	// store the link for each image file that we will download later
	urls := []string{
		"https://unsplash.com/photos/hvdnff_bieQ/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://unsplash.com/photos/HQaZKCDaax0/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://images.unsplash.com/photo-1698778573682-346d219402b5?ixlib=rb-4.0.3&q=85&fm=jpg&crop=entropy&cs=srgb&w=640",
		"https://unsplash.com/photos/Bs2jGUWu4f8/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://cdn.stocksnap.io/img-thumbs/960w/aerial-beach_DBWLJ9UUWB.jpg",
		"https://plus.unsplash.com/premium_photo-1699537318938-7af01500c359?q=80&w=3087&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		"https://images.unsplash.com/photo-1699475554452-f24c6a035a41?q=80&w=3164&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		"https://cdn.stocksnap.io/img-thumbs/960w/deer-animal_NERJJRVKBO.jpg",

		// Add more image URLs
	}

	// Sequential download
	start := time.Now()
	downloadImagesSequential(urls)
	fmt.Printf("\n \n Sequential download took: %v\n\n", time.Since(start))

	//Concurrent download
	start = time.Now()
	downloadImagesConcurrent(urls)
	fmt.Printf("\n \n Concurrent download took: %v\n", time.Since(start))
}

// Helper function to download and save a single image.
func downloadImage(url, filename string) error {
	// TODO: Implement download logic
	return nil
}
