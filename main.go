package main

import (
	"flag"
	"fmt"
	"gifhelper"
	"log"
	"os"
)

func main() {
	fmt.Println("Welcome to the Worm Race!")

	var filetype string
	var imgWidth int
	var outputFilename string
	var animOutputFile string
	flag.StringVar(&filetype, "ftype", "", "Filetype containing the worm data")
	flag.IntVar(&imgWidth, "width", 500, "Width (and height) of the image to create.")
	flag.StringVar(&outputFilename, "o", "out.png", "Name of PNG to output.")
	flag.StringVar(&animOutputFile, "a", "anim.gif", "Animated GIF to write.")
	flag.Parse()

	if filetype == "" {
		log.Fatal("Must supply a filetype with -ftype.")
	}

	datafiles := os.Args[3:]

	worms := ReadWormsFromFiles(filetype, datafiles)
	fmt.Println("Data successfully read worms from file")

	trackLength := 100.0

	race := BuildRaceTrack(worms, trackLength)

	WormRace := SimulateRace(worms, race)
	fmt.Println(len(WormRace))

	fmt.Println("Congratulations winning genotype: ", WinnerIs(WormRace[len(WormRace)-1]).genotype)

	img := WormRace[0].DrawWormsToImage(imgWidth)
	WriteImageAsPNG(img, outputFilename)

	frames := AnimateRace(WormRace, imgWidth)
	fmt.Println("len frames:", len(frames))
	gifhelper.ImagesToGIF(frames, animOutputFile)
}
