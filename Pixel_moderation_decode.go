package main

import (
	"fmt"
	"image/color"
	"image/png"
	"os"
)

func main() {

	file, err := os.Open("output.png")
	if err != nil {
		fmt.Println("Cant find image:", err)
		return
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		fmt.Println("Cant decode:", err)
		return
	}

	bounds := img.Bounds()
	var bits []byte 


PixelLoop:
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			
			pixelColor := img.At(x, y)
			c := color.RGBAModel.Convert(pixelColor).(color.RGBA)


			bit := c.R & 1
			

			bits = append(bits, bit)

			if len(bits) == 32 {
				break PixelLoop 
			}
		}
	}

	var message []byte
		for i := 0; i < len(bits); i += 8 {
		var b byte = 0 
		for j := 0; j < 8; j++ {
			b = (b << 1) | bits[i+j]
		}
		message = append(message, b)
	}

	fmt.Println("Your decoded message:", string(message))
}
