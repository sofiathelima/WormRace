package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func (w *Worm) BuildWorm(numSeg int, segLen float64) {
	fmt.Println("building worm")
	for i := range w.body {
		w.body[i] = &Segment{
			length: segLen,
		}
		if i != 0 {
			w.body[i].position.y = w.body[i-1].position.y - segLen
		}
	}
}

func StoreData(data []string) *Worm {
	numSeg, _ := strconv.Atoi(data[1])
	segLen, _ := strconv.ParseFloat(data[2], 64)
	// contraction, _ := strconv.Atoi(data[3])

	w := Worm{
		genotype: data[0],
		body:     make([]*Segment, numSeg),
	}

	w.BuildWorm(numSeg, segLen)

	return &w

}

func ReadWormFromFile(filename string) *Worm {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: something went wrong opening the file.")
		fmt.Println("Probably you gave the wrong filename.")
	}
	defer file.Close()

	// Read values as strings from file to worm with one worm per data file
	br := bufio.NewReader(file)
	data := make([]string, 0)
	for {
		var val string
		if _, err := fmt.Fscan(br, &val); err != nil {
			break
		}
		data = append(data, val)
	}

	return StoreData(data)
}

func ReadWormsFromFiles(filenames []string) []*Worm {
	worms := make([]*Worm, len(filenames))
	for i, file := range filenames {
		worms[i] = ReadWormFromFile(file)
	}
	return worms
}
