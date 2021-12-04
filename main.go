package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

// Starting over 12-03-2021
// Goal, small steps

func main() {

	//var numberOfThreads int
	//flag.IntVar(&numberOfThreads, "threads", 2, "Number of threads to use.")
	//flag.Parse()

	//Program.NumberOfInputs = INPUTCOUNT
	//Program.NumberOfOutputs = OUTPUTCOUNT
	var fcd int
	flag.IntVar(&Program.NumberOfNeos, "neos", 1000, "Number of Neos.")
	flag.IntVar(&Program.NumberOfGenes, "genes", 4, "Number of genes.")
	flag.IntVar(&Program.NumberOfLayers, "layers", 1, "Number of hidden layers.")
	flag.IntVar(&Program.NumberOfNeurons, "neurons", 3, "Number of neurons in hidden layers.")
	flag.IntVar(&Program.NumberOfSteps, "steps", 300, "Number of steps per generation.")
	flag.IntVar(&Program.NumberOfGenerations, "generations", 100, "Number of generations per simulation.")
	flag.Float64Var(&Program.Mutaions, "mutation", 0.001, "Frequecy of mutations. 0.0 to 1.0.")
	flag.IntVar(&Program.WorldX, "x", 128, "World Size in X. Image size will be 5x this amount.")
	flag.IntVar(&Program.WorldX, "y", 128, "World Size in Y. Image size will be 5x this amount.")
	flag.IntVar(&Program.MaxDistLook, "maxdist", 30, "Maximum distance a Neo can detect out to.")
	flag.IntVar(&Program.MaxHunger, "hunger", 30, "Maximum hunger.")
	flag.IntVar(&fcd, "food", 4, "Divisor for Number of Neos to food.")
	flag.Parse()

	Program.NumberOfNeos = Program.NumberOfNeos + 1 // Index starts at 1, so increase count

	if Program.NumberOfGenes < 4 {
		Program.NumberOfGenes = 4
	}
	if Program.NumberOfGenes > 255 {
		Program.NumberOfGenes = 255
	} // WOW

	if Program.NumberOfLayers == 0 {
		Program.NumberOfNeurons = INPUTCOUNT + OUTPUTCOUNT
	} else {
		Program.NumberOfNeurons = INPUTCOUNT + OUTPUTCOUNT + (Program.NumberOfLayers * Program.NumberOfNeurons)
	}

	Program.FworldX = float64(Program.WorldX)
	Program.FworldY = float64(Program.WorldY)
	Program.WorldSize = Program.WorldX * Program.WorldY

	World = make([]int, Program.WorldSize)
	WorldTmp = make([]int, Program.WorldSize)

	Program.FMaxDistLook = float64(Program.MaxDistLook)

	if fcd == 0 || fcd > Program.NumberOfNeos {
		fmt.Println("Food count devisor out of bounds.")
		os.Exit(-1)
	}

	Program.FoodCount = Program.NumberOfNeos / fcd

	Neos = make([]Neo, Program.NumberOfNeos)
	for i := 1; i < Program.NumberOfNeos; i++ {
		Neos[i].Neurons = make([]Neuron, Program.NumberOfNeurons)
		Neos[i].Genes = make([]uint32, Program.NumberOfGenes)
	}

	startsim := time.Now() // Get the current time. Used for timming the execution of the sim.

	clearworld() // Reset the world

	// Clear the world slices.
	for i := 0; i < Program.WorldSize; i++ {
		World[i] = 0
		WorldTmp[i] = 0
	}

	err := gen0init()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Simulation run time: %v\n", time.Since(startsim))

}

func gen0init() error {

	for i := 1; i < Program.NumberOfNeos; i++ {

		var nid int // nid : Neuron ID
		for j := 0; j < Program.NumberOfGenes; j++ {

			tempgene := uint32(randInt(0xFFFFFFFF))
			Neos[i].Genes[j] = tempgene
			neu := decode(tempgene)
			neu.ID = nid
			nid++

		}

		err := printgenes(i)
		if err != nil {
			return err
		}

		err = printneuron(i)
		if err != nil {
			return err
		}

	}

	return nil

}
