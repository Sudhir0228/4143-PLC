// Sudhir Ray
// PLC-4143
// Program 3
package main

import (
	"fmt"

	"github.com/Sudhir0228/modules" // Adjust the import path based on your actual module path
)

func main() {
	imageURL := "https://s3e8p5g8.rocketcdn.me/wp-content/uploads/2020/11/midwestern-state-university2.jpg" // Replace with your actual image URL
	fileName := "downloaded_image.jpg"                                                                      // Specify the desired file name

	//will download a image file using its url
	err := modules.GetImage(imageURL, fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	text := "Midwestern State University"
	outputPath := "output.png"

	//will print out colored text to a output file
	err = modules.PrintColor(text, outputPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Image saved as", outputPath)

	err = modules.Pixels(outputPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	output2 := "graymidwestern.png"

	err = modules.Grayscale(fileName, output2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

}
