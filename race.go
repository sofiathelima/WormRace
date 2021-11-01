package main

import "fmt"

func BuildRaceTrack(w []*Worm) *Race {
	return &Race{
		trackLength: 100.0,
		athletes:    w,
		startLine:   OrderedPair{y: 0},
	}
}

func SetMark(w []*Worm) *Race {
	course := BuildRaceTrack(w)
	for i, competitor := range course.athletes {
		for _, seg := range competitor.body {
			seg.position.x = float64(i + 1)
		}
	}
	return course
}

func WinnerIs(r *Race) *Worm {
	winner := r.athletes[0]
	for i := 1; i < len(r.athletes); i++ {
		if r.athletes[i-1].CalculateSpeed(r) < r.athletes[i].CalculateSpeed(r) {
			winner = r.athletes[i-1]
		} else {
			winner = r.athletes[i]
		}
	}
	return winner
}

func (w *Worm) CalculateSpeed(r *Race) float64 {
	totalLen := float64(len(w.body)) * w.body[0].length
	w.numWavesToFinish = r.trackLength / (totalLen * 0.5)
	fmt.Println("Num Waves to Finish: ", w.numWavesToFinish)
	return w.numWavesToFinish
}

func SimulateRace(w []*Worm, initialTrack *Race) []*Race {
	fmt.Println(WinnerIs(initialTrack))
	numGens := int(WinnerIs(initialTrack).numWavesToFinish)
	fmt.Println(numGens)
	timePoints := make([]*Race, numGens)
	timePoints[0] = initialTrack

	for i := 0; i < numGens; i++ {
		timePoints[i+1] = UpdateRace(timePoints[i])
	}

	return timePoints
}

func UpdateRace(r *Race) *Race {
	for _, competitor := range r.athletes {
		competitor.UpdatePosition()
	}
	return r
}

func (w *Worm) UpdatePosition() {
	if w.body[0].contracted == false {
		w.Contract()
	} else {
		w.Expand()
	}
}

func (w *Worm) Contract() {
	for i, seg := range w.body {
		seg.length = seg.length * 0.5

		if i != 0 {
			w.body[i].position.y = w.body[i-1].position.y - seg.length
		}

		seg.contracted = true
	}
}

func (w *Worm) Expand() {
	for i := len(w.body) - 2; i > 0; i-- {
		w.body[i].length = w.body[i].length * 1.5

		w.body[i-1].position.y = w.body[i].position.y + w.body[i].length

		w.body[i].contracted = false
	}
}
