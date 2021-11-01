package main

type Race struct {
	trackLength float64
	athletes    []*Worm
	startLine   OrderedPair
}

type Worm struct {
	genotype          string
	contractionFactor int
	body              []*Segment
	numWavesToFinish  float64
}

type Segment struct {
	position   OrderedPair
	length     float64
	contracted bool
}

type OrderedPair struct {
	x float64
	y float64
}
