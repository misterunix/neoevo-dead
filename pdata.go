package main

// const, types and globals

type PData struct {
	NumberOfNeos    int // NumberOfNeos : Number of Neos per generation.
	NumberOfGenes   int // NumberOfGenes : Number of genes in a Neo
	NumberOfLayers  int // NumberOfLayers : Number of layers where 0 is input, MaxLayers-1 is output. Inbetween are hidden.
	NumberOfNeurons int // NumberOfNeurons : Number of Neurons per layer
	NumberOfSteps   int // NumberOfSteps : Number of steps in this generation.
}

type Neuron struct {
	InValue  float64
	OutValue float64
	Weight   float64
	Layer    int
	Out      int
}

type Neo struct {
	ID      int
	Age     float64
	X       int
	Y       int
	Genes   []int
	Neurons []Neuron
}

var Program PData

var Neos []Neo
