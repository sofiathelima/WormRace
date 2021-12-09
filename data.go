package main

type RaceBoard struct {
	height, width float64
	trackLength   float64
	athletes      []*Worm
}

type Worm struct {
	genotype          string
	contractionFactor float64 // percentage of body length when contracted
	body              []*Segment
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
