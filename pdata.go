package main

// const, types and globals

type PData struct {
	NumberOfNeos    int // NumberOfNeos : Number of Neos per generation.
	NumberOfGenes   int // NumberOfGenes : Number of genes in a Neo
	NumberOfLayers  int // NumberOfLayers : Number of layers where 0 is input, MaxLayers-1 is output. Inbetween are hidden.
	NumberOfNeurons int // NumberOfNeurons : Number of Neurons per layer
	NumberOfSteps   int // NumberOfSteps : Number of steps in this generation.

	WorldSize int // WorldSize : Total number of cells making up the world.
	WorldX    int // WorldX : Number of cells in the X plane.
	WorldY    int // WorldY : Number of cells in the Y plane.

	NumberOfInputs  int // NumberOfInputs : The number of Input nodes. Must be equal to the NumberOfNeurons or less.
	NumberOfOutputs int // NumberOfOutputs : The number of Output nodes. Must be equal to the NumberOfNeurons or less.

	MaxHunger       int // MaxHunger :
	MaxDistanceLook int // MaxDistanceLook :

	FoodCount int // FoodCount :
}

type Neuron struct {
	InValue     float64
	OutValue    float64
	Weight      float64
	SourceLayer int
	Source      int
	OutLayer    int
	Out         int
}

type Neo struct {
	ID        int
	Age       float64
	X         int
	Y         int
	Genes     []int
	Neurons   []Neuron
	Inputs    []float64
	Outputs   []float64
	LocationX int
	LocationY int
	Hunger    int
	Direction int
	Dead      bool // Dead : True dead. False alive
}

type Point struct {
	X int
	Y int
}

var Program PData

var Neos []Neo

var World []int  // World : The world slice
var Food []Point // Food in the world

var CurrentStep int
