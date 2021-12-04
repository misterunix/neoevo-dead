package main

// const, types and globals

// Inputs
const (
	AGE               = 0  // AGE : Age of Neo. 0.0 - 1.0
	DIRCLOSESTFOOD    = 1  // CLOSESTFOOD : direction to closest food 0.0 to 1.0 - 0.0 to 360.0, -1 if no food.
	DISTANCEFOOD      = 2  // DISTANCEFOOD : Distace to closest food. 0.0 - 1.0, -1 if to far.
	CLOSESTNEO        = 3  // CLOSESTNEO : Direction to closest Neo 0.0 to 1.0 - 0.0 to 360.0, -1 if no Neo.
	DISTANCENEO       = 4  // DISTANCENEO : Distance to closest Neo 0.0 to 1.0, -1 if to far.
	POSITIONNS        = 5  // Position between North and South wall - IE Y -1.0 North 1.0 South.
	POSITIONWE        = 6  // Position between West and East wall - IE X -1.0 West 1.0 East.
	DISTANCEFORWARD   = 7  // DISTANCEFORWARD : Distance to nearest blockage forward.
	DISTANCEBACKWARDS = 8  // DISTANCEBACKWARDS : Distance to nearest blockage backwards.
	HUNGER            = 9  // HUNGER : Hunger lever. 0.0 - 1.0, 0 not hungry, 1.0 dead.
	INPUTCOUNT        = 10 // INPUTCOUNT : number of inputs.
)

// Outputs
const (
	MOVERANDOM    = 0 // MOVERANDOM : Move in any of the 4 directions.
	MOVEFORWARD   = 1 // MOVEFORWARD : Move in the forward direction.
	MOVEBACKWARDS = 2 // MOVEBACKWARDS : Move in the backwards direction - not turning.
	TURNLEFT      = 3 // Turn 90 degrees counter clockwise.
	TURNRIGHT     = 4 // Turn 90 degress clockwise.
	MOVENORTH     = 5 // MOVENORTH : Move y-1
	MOVESOUTH     = 6 // MOVESOUTH : Move y+1
	MOVEWEST      = 7 // MOVEWEST : Move x-1
	MOVEEAST      = 8 // MOVEEAST : Move x+1
	OUTPUTCOUNT   = 9 // OUTPUTCOUNT : number of outputs
)

// Point : Generic struct for holding X,Y locations.
// Storing floats so fewer conversions from int to float will need to be done.
type Point struct {
	X  int     // X : X location.
	Y  int     // Y : Y location.
	FX float64 // FX : float64 of the X location.
	FY float64 // FY : float64 of the Y location.
}

// Neuron : Struct that holds a neurons information.
type Neuron struct {
	ID           int     // ID : ID of the neuron.
	InValue      float64 // InValue : The neuron's In value
	OutValue     float64 // OutValue : The neurons Out value that has been passed through Tanh
	Weight       float64 // Weight : The neuron's weighted value.
	SourceLayer  int     // SourceLayer : The neuron's source layer.
	SourceID     int     // SourceID : The neuron's Source ID
	OutLayer     int     // OutLayer : The neuron's output layer
	OutID        int     // OutID : // The neuron's Out ID
	LinkForward  int     // LinkForward : Neuron to forward reference
	LinkBackward int     // LinkBackward : Neuron to backdirection
}

// Neo : Struct that contains all the information for a Neo.
type Neo struct {
	ID        int       // ID : Not currently used.
	Age       float64   // Age : Float value of the Neo's age. 0 - 1.0
	Location  Point     // Location : Point of X,Y of Neo's location in the world.
	Genes     []uint32  // Genes : Slice of ints holding the Genome.
	Neurons   []Neuron  // Neurons : Slive of the Neurons for the Neos.
	Inputs    []float64 // Inputs : Slice of the enviroment inputs.
	Outputs   []float64 // Outputs : Slice of the output neorons.
	LocationX int       // LocationX : Current location of the Neo.
	LocationY int       // LocationY : Current location of the Neo.
	Hunger    int       // Hunger : Current hunger of the Neo. 0 - MaxHunger. 0 is death.
	Direction int       // Direction : Current direction the Neo is moving in.
	Dead      bool      // Dead : True dead. False alive
}

// PData : Struct for holding most of the programs running data.
type PData struct {
	NumberOfNeos        int // NumberOfNeos : Number of Neos per generation.
	NumberOfGenes       int // NumberOfGenes : Number of genes in a Neo
	NumberOfLayers      int // NumberOfLayers : Number of layers where 0 is input, MaxLayers-1 is output. Inbetween are hidden.
	NumberOfNeurons     int // NumberOfNeurons : Number of Neurons per hidden layer
	NumberOfSteps       int // NumberOfSteps : Number of steps in this generation.
	NumberOfGenerations int // NumberOfGenerations : Number of generations per simulation.

	WorldSize int     // WorldSize : Total number of cells making up the world.
	WorldX    int     // WorldX : Number of cells in the X plane.
	WorldY    int     // WorldY : Number of cells in the Y plane.
	FworldX   float64 // FworldX : Float version of WorldX. Storing floats so less conversions will be needed.
	FworldY   float64 // FworldY : Float version of WorldY. Sroting floats so less conversions will be needed.

	//NumberOfInputs  int // NumberOfInputs : The number of Input nodes. Must be equal to the NumberOfNeurons or less.
	//NumberOfOutputs int // NumberOfOutputs : The number of Output nodes. Must be equal to the NumberOfNeurons or less.

	MaxHunger    int     // MaxHunger : Starting hunger
	MaxDistLook  int     // MaxDistLook : How far the Nenos can see
	FMaxDistLook float64 // FMaxDistLook : Float of MaxDistLook so less conversions needed.

	FoodCount int // FoodCount : Number of food items in the world.

	Mutaions float64 // Mutations : Frequecy of mutations. 0.0 to 1.0. 0 - no mutations. 1 - Mutation every generation.
}

var Program PData   // Program : Variable expression of PData.
var Neos []Neo      // Neos : Slice of the Neos. Neo count starts at 1 !!!!! World will use 0 to be open cell.
var World []int     // World : The world slice
var WorldTmp []int  // WorldTmp : Update to world. Copy back to World when done.
var Food []Point    // Food : Food in the world
var CurrentStep int // CurrentStep : Redundant I think.
