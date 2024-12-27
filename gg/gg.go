package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"sync"
	"time"
)

// filterParallel применяет преобразование к одной строке пикселей
func filterParallel(img draw.Image, y int, wg *sync.WaitGroup) {
	defer wg.Done()

	bounds := img.Bounds()
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		originalColor := img.At(x, y).(color.RGBA64)
		gray := uint16((originalColor.R + originalColor.G + originalColor.B) / 3)
		newColor := color.RGBA64{R: gray, G: gray, B: gray, A: originalColor.A}
		img.Set(x, y, newColor)
	}
}

func maing() {
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

	// Параллельная обработка строк изображения
	bounds := drawImg.Bounds()
	var wg sync.WaitGroup

	start := time.Now()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		wg.Add(1)
		go filterParallel(drawImg, y, &wg)
	}
	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("Parallel processing took %v\n", duration)

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
