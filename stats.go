package main

import (
	"fmt"
	"os"
	"strconv"
)

// AvgNumOfSeg calculates the average number of segments
// of the worms in the group
func AvgNumOfSeg(worms []int) int {
	total := 0
	for _, w := range worms {
		total += w
	}
	// fmt.Println("avg num segs:", total/len(worms))
	return total / len(worms)
}

// AvgNumOfSeg calculates the average og a list of float64
func Avg(worms []float64) float64 {
	total := 0.0
	for _, w := range worms {
		total += w
	}
	return total / float64(len(worms))
}

func Average(sampleData [][]string) []string {
	segments := make([]int, len(sampleData))
	for i, data := range sampleData {
		segments[i], _ = strconv.Atoi(data[1])
	}
	avgSegNum := strconv.Itoa(AvgNumOfSeg(segments))

	segLengths := make([]float64, len(sampleData))
	for i, data := range sampleData {
		segLengths[i], _ = strconv.ParseFloat(data[2], 64)
	}
	avgSegLen := strconv.FormatFloat(Avg(segLengths), 'E', -1, 64)

	cntrctnFctrs := make([]float64, len(sampleData))
	for i, data := range sampleData {
		cntrctnFctrs[i], _ = strconv.ParseFloat(data[3], 64)
	}
	avgContraction := strconv.FormatFloat(Avg(cntrctnFctrs), 'E', -1, 64)

	return []string{sampleData[0][0], avgSegNum, avgSegLen, avgContraction}
}

// SampleAverages takes a list of data lists from all trials
// from multiple individuals from a group. It calculates the
// average of each individual in the group.
func SampleAverages(groupData [][]string) map[int][]string {

	allSamples := make(map[int][][]string)

	// map each individual sample to its trials
	for _, row := range groupData {
		i, err := strconv.Atoi(row[1])
		if err != nil {
			fmt.Println("Could not read sample number in xlsx file")
			os.Exit(1)
		}
		allSamples[i] = append(allSamples[i], row)
	}

	// Find average of sample
	avgs := make(map[int][]string)

	for i, sample := range allSamples {
		for j := range sample {
			sample[j] = append(sample[j][:1], sample[j][2:]...)
		}
		avgs[i] = Average(sample)
	}

	return avgs
}

func GroupAverage(gtype string, samples map[int][]string) *Worm {
	averages := make([][]string, 0)
	for _, data := range samples {
		averages = append(averages, data)
	}

	return StoreData(Average(averages))
}

func Statistics(data [][]string) []*Worm {

	// headings := data[1]

	// var correct = []string{"genotype", "sample", "numSegments", "avgSegLength", "contractionFactor"}
	// for i, val := range headings {
	// 	if val != correct[i] {
	// 		fmt.Println("Incorrect order of data matrix")
	// 		os.Exit(1)
	// 	}
	// }

	metaData := data[1:]

	// map genotypes to all associated data
	groups := make(map[string][][]string)
	for _, row := range metaData {
		g := row[0]
		groups[g] = append(groups[g], row)
	}

	// map sample data to genotypes
	replicates := make(map[string]map[int][]string)
	for gtype, samples := range groups {
		replicates[gtype] = SampleAverages(samples)
	}

	// list average worm from multiple samples from each genotype
	reps := make([]*Worm, 0)
	for gtype, samples := range replicates {
		avg := GroupAverage(gtype, samples)
		reps = append(reps, avg)
	}

	return reps
}
