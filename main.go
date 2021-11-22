package main

import (
	"flag"
	"fmt"
	"gifhelper"
	"os"
)

func main() {
	fmt.Println("Welcome to the Worm Race!")

	datafiles := os.Args[1:]

	worms := ReadWormsFromFiles(datafiles)
	fmt.Println("Data successfully read worms from file")

	trackLength := 50.0

	race := BuildRaceTrack(worms, trackLength)

	fmt.Println("Congratulations winning genotype: ", WinnerIs(race).genotype)

	WormRace := SimulateRace(worms, race)
	fmt.Println(len(WormRace))

	var imgWidth int
	var outputFilename string
	var animOutputFile string
	flag.IntVar(&imgWidth, "width", 500, "Width (and height) of the image to create.")
	flag.StringVar(&outputFilename, "o", "out.png", "Name of PNG to output.")
	flag.StringVar(&animOutputFile, "a", "anim.gif", "Animated GIF to write.")
	flag.Parse()

	img := WormRace[0].DrawWormsToImage(imgWidth)
	WriteImageAsPNG(img, outputFilename)

	frames := AnimateRace(WormRace, imgWidth)
	fmt.Println("len frames:", len(frames))
	gifhelper.ImagesToGIF(frames, animOutputFile)
}
