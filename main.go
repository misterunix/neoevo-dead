package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {

	Program.NumberOfInputs = 10
	Program.NumberOfOutputs = 10

	Program.NumberOfNeos = 3000
	Program.NumberOfGenes = 16
	Program.NumberOfNeurons = 32
	Program.NumberOfLayers = 3

	Program.NumberOfSteps = 300

	Program.WorldX = 128
	Program.WorldY = 128
	Program.WorldSize = Program.WorldX * Program.WorldY
	World = make([]int, Program.WorldSize)
	WorldTmp = make([]int, Program.WorldSize)

	Program.MaxDistanceLook = 30
	Program.FoodCount = Program.NumberOfNeos / 4
	Program.MaxHunger = 30

	for i := 0; i < Program.WorldSize; i++ {
		World[i] = 0
		WorldTmp[i] = 0
	}

	err := initNeos()
	if err != nil {
		log.Fatalln(err)
	}

	PutNeosInWorld()

	for count := 0; count < Program.NumberOfSteps; count++ {
		Step0()
		Step1()

		var wg sync.WaitGroup
		quickCount := 0
		for ni := 1; ni <= Program.NumberOfNeos; ni++ {
			//fmt.Println("quickCount", quickCount)
			if quickCount == 0 {
				wg.Add(10)
			}
			go Step2(ni, &wg)
			quickCount++
			if quickCount == 10 {
				wg.Wait()
				quickCount = 0
			}
		}

		CurrentStep++

		//fmt.Println(CurrentStep)

	}

	/*
		for i := range Neos {
			if i == 0 { // skip 0
				continue
			}
			fmt.Println("Neo", i)
			err := PrintGenes(i)
			if err != nil {
				log.Fatalln(err)
			}

			//err = PrintNeuron(i)
			err = PrintNet(i)
			if err != nil {
				log.Fatalln(err)
			}
		}
	*/
}

func PutNeosInWorld() error {

	for i := 1; i <= Program.NumberOfNeos; i++ {
		stupidCount := 0
		for {
			x := randInt(Program.WorldX)
			y := randInt(Program.WorldY)
			if World[XYtoIndex(x, y)] == 0 {
				World[XYtoIndex(x, y)] = i
				Neos[i].LocationX = x
				Neos[i].LocationY = y
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

	Neos = make([]Neo, Program.NumberOfNeos+1)

	err := PutNeosInWorld()
	if err != nil {
		return err
	}

	for i := 1; i <= Program.NumberOfNeos; i++ {

		Neos[i].Neurons = make([]Neuron, 0)
		Neos[i].Genes = make([]int, 0)
		Neos[i].Inputs = make([]float64, Program.NumberOfInputs)
		Neos[i].Outputs = make([]float64, Program.NumberOfOutputs)
		Neos[i].Hunger = 0
		for j := 0; j < Program.NumberOfGenes; j++ {
			buildGenes(i)
		}

	}
	return nil
}
