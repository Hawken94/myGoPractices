package testPractices

import (
	"image/color"
	"log"
)

func main1() {
	qr, err := qrcode.New("hello hawken", qrcode.Medium)
	if err != nil {
		log.Fatal(err)
	} else {
		qr.BackgroundColor = color.RGBA{50, 205, 50, 255}
		qr.ForegroundColor = color.White
		qr.WriteFile(256, "./qrcode.png")
	}
}
