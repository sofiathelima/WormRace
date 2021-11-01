package main

import (
	"canvas"
	"image"
)

func (w *Worm) DrawWorms(wormRace []*Race, imgWidth int) image.Image {
	// set a new square canvas
	c := canvas.CreateNewPalettedCanvas(1000, 1000, nil)

	// create a black background
	c.SetFillColor(canvas.MakeColor(0, 0, 0))
	c.ClearRect(0, 0, 1000, 1000)
	c.Fill()

	return canvas.GetImage(c)
}
