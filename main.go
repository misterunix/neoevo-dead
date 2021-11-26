package main

import (
	"fmt"
	"log"
)

type Neuron struct {
	ID       int
	Source   int  // Source : ID of the neuron
	SrcFlag  bool // SrcFlag : false InputNeuron - true HiddenNeuron
	Sink     int  // Sink : ID of thr neuron
	SnkFlag  bool // SnkFlag : false HiddenNeuron - true OutputNeuron
	InValue  float64
	OutValue float64
	Weight   float64
	OutIndex int
}

type NeoBase struct {
	GeneCount int
	Genes     []int
	Neurons   []Neuron
	Hidden    int
	LocationX int
	LocationY int
	Inputs    []float64
	Outputs   []float64
	Hunger    int
}

var Neo []NeoBase
var BaseNeoCount int
var Generation int
var inputcount int
var hiddencount int
var outputcount int

var WorldX int
var WorldY int

var Reset bool

var StepsPerGeneration int
var CurrentStep int

func main() {

	WorldX = 128
	WorldY = 128

	inputcount = 5
	hiddencount = 3
	outputcount = 4
	BaseNeoCount = 10
	genecount := 8
	StepsPerGeneration = 300
	CurrentStep = 0
	maxhunger := 30

	Neo = make([]NeoBase, BaseNeoCount)
	for i := 0; i < BaseNeoCount; i++ {
		Neo[i].Genes = make([]int, genecount)
		Neo[i].Inputs = make([]float64, 12)
		Neo[i].Outputs = make([]float64, 20)
	}

	Reset = true

	if Generation == 0 {
		for i := 0; i < BaseNeoCount; i++ {
			for j := 0; j < genecount; j++ {
				for {
					Neo[i].Genes[j] = randInt(0xFFFFFFFF)
					if fitness(i, j) {
						break
					}
				}
			}
		}
	}

	if Reset {
		for i := range Neo {
			Neo[i].LocationX = randInt(128)
			Neo[i].LocationY = randInt(128)
			Neo[i].Hunger = maxhunger
		}
	}

	for i := range Neo {
		distanceFromNorth := float64(WorldY-Neo[i].LocationY)/float64(WorldY)*2.0 - 1.0
		distanceFromWest := float64(WorldX-Neo[i].LocationX)/float64(WorldX)*2.0 - 1.0
		Neo[i].Inputs[0] = float64(CurrentStep) / float64(StepsPerGeneration)
		Neo[i].Inputs[3] = distanceFromNorth
		Neo[i].Inputs[4] = distanceFromWest
		Neo[i].Inputs[7] = float64(Neo[i].Hunger) / float64(maxhunger)
	}

	for i := 0; i < BaseNeoCount; i++ {
		err := buildNeurons(i)
		if err != nil {
			log.Fatalln(err)
		}
	}

	for i := 0; i < BaseNeoCount; i++ {
		err := printgenes(i)
		if err != nil {
			log.Fatalln(err)
		}
		err = printneurons(i)
		if err != nil {
			log.Fatalln(err)
		}

	}

}

func Step(id int) error {
	if id >= BaseNeoCount || id < 0 {
		return fmt.Errorf("Step id '%d' is out of bounds", id)
	}

	// Move envor inputs into neurons

	return nil
}
