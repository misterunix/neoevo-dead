package main

// const, types and globals

type PData struct {
	MaxNeos   int // MaxNeos : Number of Neos per generation.
	MaxLayers int // MaxLayers : Number of layers where 0 is input, MaxLayers-1 is output. Inbetween are hidden.

	MaxSteps int // MaxSteps : Number of steps in this generation.

}

type Neuron struct {
	InValue  float64
	OutValue float64
	Weight   float64
	ID       int
	Layer    int
	Out      []int
}

type Neo struct {
	ID      int
	Age     float64
	X       int
	Y       int
	Neurons map[int]Neuron
}

var Program PData

var Neos []Neo
