package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"time"
)

// filter применяет преобразование для каждого пикселя (например, оттенки серого)
func filter(img draw.Image) {
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y).(color.RGBA64)
			gray := uint16((originalColor.R + originalColor.G + originalColor.B) / 3)
			newColor := color.RGBA64{R: gray, G: gray, B: gray, A: originalColor.A}
			img.Set(x, y, newColor)
		}
	}
}

func mains() {
	// Открытие изображения
	inputFile, err := os.Open("C:/Users/Lenovo/Downloads/R.png")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	img, _, err := image.Decode(inputFile)
	if err != nil {
		panic(err)
	}

	drawImg, ok := img.(draw.Image)
	if !ok {
		panic("Image conversion to draw.Image failed")
	}

	// Применение фильтра и замер времени
	start := time.Now()
	filter(drawImg)
	duration := time.Since(start)
	fmt.Printf("Sequential processing took %v\n", duration)

	// Сохранение изображения
	outputFile, err := os.Create("C:/Users/Lenovo/Pictures/Screenshots")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, drawImg)
	if err != nil {
		panic(err)
	}
}
