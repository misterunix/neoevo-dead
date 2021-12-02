<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# neoevo

```go
import "neoevo"
```

## Index

- [Constants](<#constants>)
- [Variables](<#variables>)
- [func GetAngle(x1, y1, x2, y2 float64) float64](<#func-getangle>)
- [func GetDistance(x1, y1, x2, y2 float64) float64](<#func-getdistance>)
- [func PlaceFood()](<#func-placefood>)
- [func PrintGenes(id int) error](<#func-printgenes>)
- [func PrintIO(id int) error](<#func-printio>)
- [func PrintNet(id int) error](<#func-printnet>)
- [func PrintNeuron(id int) error](<#func-printneuron>)
- [func PutNeosInWorld() error](<#func-putneosinworld>)
- [func Step0(i int)](<#func-step0>)
- [func Step1(i int)](<#func-step1>)
- [func Step2(i int)](<#func-step2>)
- [func Step3(i int)](<#func-step3>)
- [func Step4(i int)](<#func-step4>)
- [func XYtoIndex(x, y int) int](<#func-xytoindex>)
- [func buildGenes(id int) error](<#func-buildgenes>)
- [func createpng(i int)](<#func-createpng>)
- [func initNeos() error](<#func-initneos>)
- [func main()](<#func-main>)
- [func probability(p float64) bool](<#func-probability>)
- [func randFloat() float64](<#func-randfloat>)
- [func randFloatFullValue() float64](<#func-randfloatfullvalue>)
- [func randInt(max int) int](<#func-randint>)
- [type Neo](<#type-neo>)
- [type Neuron](<#type-neuron>)
  - [func decode(g int) Neuron](<#func-decode>)
- [type PData](<#type-pdata>)
- [type Point](<#type-point>)
  - [func DirectionToStep(d int) Point](<#func-directiontostep>)
  - [func IndexToXY(i int) Point](<#func-indextoxy>)


## Constants

```go
const (
    DEG2RAD = 0.0174532925
    RAD2DEG = 57.2957795130
)
```

## Variables

```go
var CurrentStep int // CurrentStep : Redundant I think.
```

```go
var Food []Point // Food : Food in the world
```

```go
var Neos []Neo // Neos : Slice of the Neos.
```

```go
var World []int // World : The world slice
```

```go
var WorldTmp []int // WorldTmp : Update to world. Copy back to World when done.
```

## func [GetAngle](<https://github.com/misterunix/neoevo/blob/main/location.go#L19>)

```go
func GetAngle(x1, y1, x2, y2 float64) float64
```

GetAngle : Return the angle in degrees from x1\,y1 to x2\,y2

## func [GetDistance](<https://github.com/misterunix/neoevo/blob/main/location.go#L13>)

```go
func GetDistance(x1, y1, x2, y2 float64) float64
```

GetDistance : Return the distance between x1\,y1 to x2\,y2

## func [PlaceFood](<https://github.com/misterunix/neoevo/blob/main/steps.go#L8>)

```go
func PlaceFood()
```

## func [PrintGenes](<https://github.com/misterunix/neoevo/blob/main/printing.go#L8>)

```go
func PrintGenes(id int) error
```

PrintGenes : Print the genes of the Neo in hex format

## func [PrintIO](<https://github.com/misterunix/neoevo/blob/main/printing.go#L26>)

```go
func PrintIO(id int) error
```

PrintIO : Prints the Neos Iinput and Output slices\.

## func [PrintNet](<https://github.com/misterunix/neoevo/blob/main/printing.go#L55>)

```go
func PrintNet(id int) error
```

PrintNet : Print the Neo's net list\.

## func [PrintNeuron](<https://github.com/misterunix/neoevo/blob/main/printing.go#L43>)

```go
func PrintNeuron(id int) error
```

PrintNeuron : Prints the Neos Neurons\. Just a dump of the slice\.

## func [PutNeosInWorld](<https://github.com/misterunix/neoevo/blob/main/main.go#L114>)

```go
func PutNeosInWorld() error
```

## func [Step0](<https://github.com/misterunix/neoevo/blob/main/steps.go#L25>)

```go
func Step0(i int)
```

Step1 : Fill in all the inputs from the environment\.

## func [Step1](<https://github.com/misterunix/neoevo/blob/main/steps.go#L185>)

```go
func Step1(i int)
```

Step1 : Move the Neo's inputs to the neurons\.

## func [Step2](<https://github.com/misterunix/neoevo/blob/main/steps.go#L196>)

```go
func Step2(i int)
```

Step2 : Propagate out to in and sum and pass through Tanh\.

## func [Step3](<https://github.com/misterunix/neoevo/blob/main/steps.go#L238>)

```go
func Step3(i int)
```

Step3 : Check if Neos died from hunger\.

## func [Step4](<https://github.com/misterunix/neoevo/blob/main/steps.go#L252>)

```go
func Step4(i int)
```

Step4 : Do movement

## func [XYtoIndex](<https://github.com/misterunix/neoevo/blob/main/location.go#L31>)

```go
func XYtoIndex(x, y int) int
```

Convert X\,Y to index value for positioning in the world space\.

## func [buildGenes](<https://github.com/misterunix/neoevo/blob/main/decode.go#L42>)

```go
func buildGenes(id int) error
```

buildGenes :

## func [createpng](<https://github.com/misterunix/neoevo/blob/main/imaging.go#L9>)

```go
func createpng(i int)
```

## func [initNeos](<https://github.com/misterunix/neoevo/blob/main/main.go#L142>)

```go
func initNeos() error
```

## func [main](<https://github.com/misterunix/neoevo/blob/main/main.go#L12>)

```go
func main()
```

## func [probability](<https://github.com/misterunix/neoevo/blob/main/steps.go#L263>)

```go
func probability(p float64) bool
```

probability : return true if random is less than p\.

## func [randFloat](<https://github.com/misterunix/neoevo/blob/main/rnd.go#L11>)

```go
func randFloat() float64
```

randFloat : returns a positive float from 0 to 1

## func [randFloatFullValue](<https://github.com/misterunix/neoevo/blob/main/rnd.go#L22>)

```go
func randFloatFullValue() float64
```

randFullRange : Return a float64 in the range of \> \-1\.0 and \< \+1\.0

## func [randInt](<https://github.com/misterunix/neoevo/blob/main/rnd.go#L33>)

```go
func randInt(max int) int
```

randInt : returns a integer that is between 0 and max\.

## type [Neo](<https://github.com/misterunix/neoevo/blob/main/pdata.go#L38-L52>)

Neo : Struct that contains all the information for a Neo\.

```go
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
```

## type [Neuron](<https://github.com/misterunix/neoevo/blob/main/pdata.go#L27-L35>)

Neuron : Struct that holds a neurons information\.

```go
type Neuron struct {
    InValue     float64 // InValue : The neuron's In value
    OutValue    float64 // OutValue : The neurons Out value that has been passed through Tanh
    Weight      float64 // Weight : The neuron's weighted value.
    SourceLayer int     // SourceLayer : The neuron's source layer.
    Source      int     // Source : The neuron's Source ID
    OutLayer    int     // OutLayer : The neuron's output layer
    Out         int     // Out : // The neuron's Out ID
}
```

### func [decode](<https://github.com/misterunix/neoevo/blob/main/decode.go#L6>)

```go
func decode(g int) Neuron
```

decode :

## type [PData](<https://github.com/misterunix/neoevo/blob/main/pdata.go#L6-L24>)

PData : Struct for holding most of the programs running data\.

```go
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
```

```go
var Program PData // Program : Variable expression of PData.
```

## type [Point](<https://github.com/misterunix/neoevo/blob/main/pdata.go#L55-L58>)

Point : Generic struct for holding X\,Y locations\.

```go
type Point struct {
    X   int // X : X location.
    Y   int // Y : Y location.
}
```

### func [DirectionToStep](<https://github.com/misterunix/neoevo/blob/main/location.go#L45>)

```go
func DirectionToStep(d int) Point
```

DirectionToStep : Convert direction to x\,y stepping

### func [IndexToXY](<https://github.com/misterunix/neoevo/blob/main/location.go#L36>)

```go
func IndexToXY(i int) Point
```

IndexToXY : Convert the index into X\,Y for the world space\.



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)