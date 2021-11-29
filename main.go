package main

import (
	"fmt"
	"log"
)

func main() {

	Program.NumberOfInputs = 10
	Program.NumberOfOutputs = 5

	Program.NumberOfNeos = 10
	Program.NumberOfGenes = 16
	Program.NumberOfNeurons = 32
	Program.NumberOfLayers = 4

	Program.NumberOfSteps = 300

	Program.WorldX = 128
	Program.WorldY = 128
	Program.WorldSize = Program.WorldX * Program.WorldY
	World = make([]int, Program.WorldSize)

	for i := 0; i < Program.WorldSize; i++ {
		World[i] = -1
	}

	err := initNeos()
	if err != nil {
		log.Fatalln(err)
	}

}

func PutNeosInWorld() error {

	for i := 0; i < Program.NumberOfNeos; i++ {
		stupidCount := 0
		for {
			x := randInt(Program.WorldX)
			y := randInt(Program.WorldY)
			if World[XYtoIndex(x, y)] == -1 {
				World[XYtoIndex(x, y)] = i
				break
			}
			stupidCount++
			if stupidCount == Program.WorldSize*4 {
				return fmt.Errorf("PutNeosInWorld was not able to place all the Neos. Max count reached")
			}
		}
	}
	return nil
}

func initNeos() error {

	Neos = make([]Neo, Program.NumberOfNeos)

	err := PutNeosInWorld()
	if err != nil {
		return err
	}

	for i := 0; i < Program.NumberOfNeos; i++ {

		Neos[i].Neurons = make([]Neuron, 0)
		Neos[i].Genes = make([]int, 0)

		for j := 0; j < Program.NumberOfGenes; j++ {
			buildGenes(i)
		}

		err := PrintGenes(i)
		if err != nil {
			return err
		}

		err = PrintNeuron(i)
		if err != nil {
			return err
		}
	}
	return nil
}
