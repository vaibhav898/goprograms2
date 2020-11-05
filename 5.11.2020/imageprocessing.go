package main

import (
	"image"
	"image/color"
	"log"

	"github.com/disintegration/imaging"
)

func main() {
	
	src, err := imaging.Open("test.png")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	
	src = imaging.CropAnchor(src, 300, 300, imaging.Center)

	
	src = imaging.Resize(src, 200, 0, imaging.Lanczos)

	
	img1 := imaging.Blur(src, 5)

	
	img2 := imaging.Grayscale(src)
	img2 = imaging.AdjustContrast(img2, 20)
	img2 = imaging.Sharpen(img2, 2)

	
	img3 := imaging.Invert(src)

	
	img4 := imaging.Convolve3x3(
		src,
		[9]float64{
			-1, -1, 0,
			-1, 1, 1,
			0, 1, 1,
		},
		nil,
	)

	
	dst := imaging.New(400, 400, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, img1, image.Pt(0, 0))
	dst = imaging.Paste(dst, img2, image.Pt(0, 200))
	dst = imaging.Paste(dst, img3, image.Pt(200, 0))
	dst = imaging.Paste(dst, img4, image.Pt(200, 200))

	err = imaging.Save(dst, "test.jpg")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}