package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/misterunix/sniffle/systemcpu"
)

func main() {

	var numberOfThreads int

	flag.IntVar(&numberOfThreads, "threads", 2, "Number of threads to use.")

	flag.Parse()

	buffers = make(chan bool, numberOfThreads)

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

	sr := systemcpu.SysRun{}

	sr.Update()

	fmt.Println("CPUCount", sr.CPUCount)
	fmt.Println("Architecture", sr.Architecture)
	fmt.Println("OS", sr.OS)
	fmt.Println("GoVersion", sr.GoVersion)
	fmt.Println("Goroutines", sr.Goroutines)
	fmt.Println("PID", sr.PID)

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

		//fmt.Println("Loop Count:", count)

		var wg sync.WaitGroup
		for ni := 1; ni <= Program.NumberOfNeos; ni++ {
			wg.Add(1)
			go Step2(ni, &wg)
			//go Step2(ni+1, &wg)
			//fmt.Println("Len:", len(buffers), cap(buffers))
		}
		fmt.Println("len:", len(buffers), cap(buffers))
		//sr.GetGo()
		//fmt.Println("Goroutines", sr.Goroutines)
		wg.Wait()

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
