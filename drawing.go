package main

import (
	"bufio"
	"canvas"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

// AnimateRace creates an image for every step in wormRace.
// The image size is 3 times the track length
func AnimateRace(wormRace []*RaceBoard, imgWidth int) []image.Image {
	images := make([]image.Image, 0, len(wormRace)*5)

	for _, r := range wormRace {
		images = append(images, r.DrawWormsToImage(imgWidth))
		images = append(images, r.DrawWormsToImage(imgWidth))
		images = append(images, r.DrawWormsToImage(imgWidth))
		images = append(images, r.DrawWormsToImage(imgWidth))
		images = append(images, r.DrawWormsToImage(imgWidth))
	}
	return images
}

func (r *RaceBoard) DrawWormsToImage(canvasWidth int) image.Image {

	aspectRatio := r.height / r.width

	canvasHeight := int(float64(canvasWidth) * aspectRatio)
	c := canvas.CreateNewCanvas(canvasWidth, canvasHeight)

	// create a black background
	c.SetFillColor(canvas.MakeColor(0, 0, 0))
	c.ClearRect(0, 0, canvasWidth, canvasHeight)
	c.Fill()

	// DrawStartLine(c, cnvsSize, float64(cnvsSize)/3)

	for _, w := range r.athletes {

		scale := float64(canvasHeight) / r.height // 3.33

		for _, seg := range w.body {
			cx := seg.position.y * scale
			cy := seg.position.x * scale
			ry := scale / 2
			if seg.contracted {
				ry = ry * 1.5
			}
			rx := (seg.length / 2) * scale
			c.SetFillColor(canvas.MakeColor(255, 255, 255))
			c.Ellipse(cx, cy, rx, ry)
			c.Fill()
		}
		// w.DrawWormToCanvas(c, r.boardSize.x, canvasHeight)
	}

	return canvas.GetImage(c)
}

// func DrawStartLine(c canvas.Canvas, cnvsSize int, y float64) {

// 	c.MoveTo(0, y)
// 	c.LineTo(float64(cnvsSize), y)
// }

// func (w *Worm) DrawWormToCanvas(c canvas.Canvas, bSize float64, cSize int) {

// 	scale := float64(cSize) / bSize

// 	for _, seg := range w.body {
// 		cx := seg.position.y * scale
// 		cy := seg.position.x * scale
// 		r := (seg.length / 2) * scale
// 		c.SetFillColor(canvas.MakeColor(255, 255, 255))
// 		c.Circle(cx, cy, r)
// 		c.Fill()
// 	}

// }

func WriteImageAsPNG(i image.Image, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	b := bufio.NewWriter(f)
	err = png.Encode(f, i)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Wrote %s OK.\n", filename)
}
