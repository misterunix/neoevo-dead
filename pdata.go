package main

// const, types and globals

// PData : Struct for holding most of the programs running data.
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

	MaxHunger       int // MaxHunger : Starting hunger
	MaxDistanceLook int // MaxDistanceLook : How far the Nenos can see

	FoodCount int // FoodCount : Number of food items in the world.
}

// Neuron : Struct that holds a neurons information.
type Neuron struct {
	InValue     float64 // InValue : The neuron's In value
	OutValue    float64 // OutValue : The neurons Out value that has been passed through Tanh
	Weight      float64 // Weight : The neuron's weighted value.
	SourceLayer int     // SourceLayer : The neuron's source layer.
	Source      int     // Source : The neuron's Source ID
	OutLayer    int     // OutLayer : The neuron's output layer
	Out         int     // Out : // The neuron's Out ID
	LinkForward int
}

// Neo : Struct that contains all the information for a Neo.
type Neo struct {
	ID        int       // ID : Not currently used.
	Age       float64   // Age : Float value of the Neo's age. 0 - 1.0
	X         int       // X : Current X location of the Neo
	Y         int       // Y: Current Y location of the Neo
	Genes     []int     // Genes : Slice of ints holding the Genome.
	Neurons   []Neuron  // Neurons : Slive of the Neurons for the Neos.
	Inputs    []float64 // Inputs : Slice of the enviroment inputs.
	Outputs   []float64 // Outputs : Slice of the output neorons.
	LocationX int       // LocationX : Current location of the Neo.
	LocationY int       // LocationY : Current location of the Neo.
	Hunger    int       // Hunger : Current hunger of the Neo. 0 - MaxHunger. 0 is death.
	Direction int       // Direction : Current direction the Neo is moving in.
	Dead      bool      // Dead : True dead. False alive
}

// Point : Generic struct for holding X,Y locations.
type Point struct {
	X int // X : X location.
	Y int // Y : Y location.
}

var Program PData // Program : Variable expression of PData.

var Neos []Neo // Neos : Slice of the Neos.

var World []int // World : The world slice

var WorldTmp []int // WorldTmp : Update to world. Copy back to World when done.

var Food []Point // Food : Food in the world

var CurrentStep int // CurrentStep : Redundant I think.
