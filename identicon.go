package main

import (
	"crypto/sha512"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	str := "test"
	length := 60
	fmt.Println(str)
	generateIdenticon(str, length)
}

func hash(str string) []byte {
	//init the hash
	sha512 := sha512.New()

	//pass the string
	sha512.Write([]byte(str))

	//64bit hash code
	return sha512.Sum(nil)
}

func generateIdenticon(str string, length int) {

	hashCode := hash(str)

	imgColor := fixColor(hashCode)

	m := image.NewRGBA(image.Rect(0, 0, length, length))

	oddColor := color.RGBA{imgColor[0], imgColor[1], imgColor[2], 255}
	evenColor := color.RGBA{imgColor[3], imgColor[4], imgColor[5], 255}

	draw.Draw(m, m.Bounds(), &image.Uniform{oddColor}, image.ZP, draw.Src)

	posX, posY, index := 0, 0, 0

	for x := 0; x < 6; x++ {
		for y := 0; y < 6; y++ {

			if x+y == 0 || (x == 0 && y == 5) || (x == 5 && y == 0) || (x == 5 && y == 5) {
				go drawRect(m, oddColor, posX, posY, posX+10, posY+10)
			} else if hashCode[index]%2 == 0 {
				go drawRect(m, evenColor, posX, posY, posX+10, posY+10)
				index++
			} else {
				go drawRect(m, oddColor, posX, posY, posX+10, posY+10)
				index++
			}
			posX += 10
		}
		posX = 0
		posY += 10
	}

	f, _ := os.OpenFile(str+".png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, m)
}

func fixColor(hashCode []byte) []byte {
	return hashCode[0:6]
}

func drawRect(mainImage draw.Image, colorObj color.RGBA, x1 int, y1 int, x2 int, y2 int) {
	temp := image.NewRGBA(image.Rect(x1, y1, x2, y2))
	draw.Draw(mainImage, temp.Bounds(), &image.Uniform{colorObj}, image.ZP, draw.Src)
}
