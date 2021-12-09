package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
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
	contraction, _ := strconv.ParseFloat(data[3], 64)

	w := Worm{
		genotype:          data[0],
		body:              make([]*Segment, numSeg),
		contractionFactor: contraction,
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

func ReadDataFromExcel(filename string) []*Worm {
	file, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println("Error: something went wrong opening the xlsx file.")
	}

	averageWorms := make([]*Worm, 0)

	// Reach each sheet in xlsx file
	// Read rows as [][]strings from xlsx file
	for _, name := range file.GetSheetMap() {
		data, err := file.GetRows(name)
		if err != nil {
			log.Fatal(err)
		}
		averageWorms = append(averageWorms, Statistics(data)...)
	}

	return averageWorms
}

func ReadWormsFromFiles(filetype string, filenames []string) []*Worm {
	worms := make([]*Worm, 0)
	for _, file := range filenames {
		if filetype == "xlsx" {
			worms = ReadDataFromExcel(file)
		} else if filetype == "txt" {
			worms = append(worms, ReadWormFromFile(file))
		}
	}
	return worms
}
