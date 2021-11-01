package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the Worm Race!")

	datafiles := os.Args[1:]

	worms := ReadWormsFromFiles(datafiles)
	fmt.Println("Data successfully read worms from file")

	startPositions := SetMark(worms)
	for i, worm := range startPositions.athletes {
		fmt.Println("i:", i, "worm:", worm)
	}

	fmt.Println("Congratulations winning genotype: ", WinnerIs(startPositions).genotype)

	// WormRace := SimulateRace(worms, startPositions)
	// fmt.Println(WormRace)

	// var imgWidth int
	// var animOutputFile string
	// flag.IntVar(&imgWidth, "width", 500, "Width (and height) of the image to create.")
	// flag.StringVar(&animOutputFile, "a", "anim.gif", "Animated GIF to write.")
	// flag.Parse()

	// 	frames := DrawWorms(WormRace, imgWidth)
	// 	gifhelper.ImagesToGIF(frames)
}
