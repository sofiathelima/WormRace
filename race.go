package main

// BuildRaceTrack returns a Race with track of #length
// and sets athletes to the list of worms
// and sets each worm to the same starting y position with SetMark()
// and sets them at different x posiions (lanes)
func BuildRaceTrack(w []*Worm, trackLength float64) *RaceBoard {

	track := &RaceBoard{
		height:      float64(len(w)) * 10,
		width:       trackLength * 1.5,
		trackLength: trackLength,
		athletes:    w,
	}

	lanes := float64(len(w) + 1)

	for i, competitor := range w {
		competitor.SetMark(trackLength * 0.25)
		for _, seg := range competitor.body {
			seg.position.x = float64(i+1) * track.height / lanes
		}
	}

	return track
}

// SetMark sets a worm to a given y position
// and sets all of its segments behind it
func (w *Worm) SetMark(y float64) {

	w.body[0].position.y = y

	for i, seg := range w.body {

		if i != 0 {
			w.body[i].position.y = w.body[i-1].position.y - seg.length
		}

	}

}

// WinnerIs returns the worm in the race with the fasted
// calculated speed (number of waves to cross the track)
func WinnerIs(r *RaceBoard) *Worm {
	winner := r.athletes[0]
	for i := 1; i < len(r.athletes); i++ {
		if r.athletes[i].body[0].position.y > winner.body[0].position.y {
			winner = r.athletes[i]
		}
	}
	return winner
}

// SimulateRace returns a list of race timepoints
// with numGens calculated from the winner
func SimulateRace(w []*Worm, initialTrack *RaceBoard) []*RaceBoard {

	timePoints := make([]*RaceBoard, 0)
	timePoints = append(timePoints, initialTrack)
	winner := WinnerIs(initialTrack)

	for winner.body[0].position.y <= initialTrack.trackLength {
		nextStep := UpdateRace(timePoints[len(timePoints)-1])
		winner = WinnerIs(nextStep)
		timePoints = append(timePoints, nextStep)
	}

	return timePoints
}

func (s *Segment) CopySegment() *Segment {
	seg := *s
	return &seg
}

func (w *Worm) CopyWorm() *Worm {
	var wormy Worm
	wormy.genotype = w.genotype
	wormy.body = make([]*Segment, len(w.body))
	wormy.contractionFactor = w.contractionFactor
	for i, seg := range w.body {
		wormy.body[i] = seg.CopySegment()
	}
	return &wormy
}

func (currRace *RaceBoard) CopyRaceBoard() *RaceBoard {
	var nextStep RaceBoard
	nextStep.height = currRace.height
	nextStep.width = currRace.width
	nextStep.trackLength = currRace.trackLength
	nextStep.athletes = make([]*Worm, len(currRace.athletes))

	for i, currWorm := range currRace.athletes {
		nextStep.athletes[i] = currWorm.CopyWorm()
	}

	return &nextStep
}

func UpdateRace(r *RaceBoard) *RaceBoard {
	nextStep := r.CopyRaceBoard()
	for _, competitor := range nextStep.athletes {
		competitor.UpdatePosition()
	}
	return nextStep
}

func (w *Worm) UpdatePosition() {

	if !w.body[0].contracted {
		w.Contract()
	} else {
		w.Expand()
	}
}

func (w *Worm) Contract() {

	for _, seg := range w.body {
		seg.length = seg.length * w.contractionFactor

		seg.contracted = true
	}

	w.SetMark(w.body[0].position.y)
}

func (w *Worm) Expand() {

	for _, seg := range w.body {
		seg.length = seg.length / w.contractionFactor

		seg.contracted = false
	}

	totalLen := float64(len(w.body)) * w.body[0].length
	newPosition := w.body[len(w.body)-1].position.y + (totalLen)

	w.SetMark(newPosition)

}
