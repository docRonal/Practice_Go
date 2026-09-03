package main

import(
"fmt"
"image"
"image/color"
"image/png"
"os"
)

func main(){

file, err := os.Open("Image.png")
	if err!=nil{
		fmt.Println("Cant find image")
		return 
	}
defer file.Close()

img, err := png.Decode(file)
	if err!=nil{
		fmt.Println("Cant decode")
		return
	}
	bounds := img.Bounds()
	newImg :=image.NewRGBA(bounds)
	
	message := "hack"
	var bits []byte 
	
	for i := 0; i < len(message); i++ {
		for j := 7; j >= 0; j-- {
			bit := (message[i] >> j) & 1
			bits = append(bits, bit)
		}
	}
	
	bitIndex := 0
	
	for y :=bounds.Min.Y; y< bounds.Max.Y; y++{
		for x:=bounds.Min.X; x<bounds.Max.X; x++{
		pixelColor := img.At(x,y)
		
		c:=color.RGBAModel.Convert(pixelColor).(color.RGBA)
		if bitIndex < len(bits){
			secretBit := bits[bitIndex]
			c.R = (c.R & 254) | secretBit
			bitIndex++
		} 
		
		newImg.Set(x,y,c)
		}
	}
	
outFile, err := os.Create("output.png")

	if err!=nil{
		fmt.Println("Cant save file", err)
		return
	}
defer outFile.Close()

png.Encode(outFile,newImg)
fmt.Println("File saved")
}
