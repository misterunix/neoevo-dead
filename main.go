package main

import (
	"fmt"
	"log"
	"time"

	"github.com/misterunix/parallel"
	"github.com/misterunix/sniffle/systemcpu"
)

func main() {

	//var numberOfThreads int
	//flag.IntVar(&numberOfThreads, "threads", 2, "Number of threads to use.")
	//flag.Parse()

	Program.NumberOfInputs = 10
	Program.NumberOfOutputs = 10

	Program.NumberOfNeos = 1000
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

	start := time.Now()

	for i := 0; i < Program.WorldSize; i++ {
		World[i] = 0
		WorldTmp[i] = 0
	}

	err := initNeos()
	if err != nil {
		log.Fatalln(err)
	}

	for count := 0; count < Program.NumberOfSteps; count++ {

		p0 := parallel.NewParallel()
		for ni := 1; ni <= Program.NumberOfNeos; ni++ {
			p0.Register(Step0, ni).SetReceivers()
		}
		p0.Run()

		p1 := parallel.NewParallel()
		for ni := 1; ni <= Program.NumberOfNeos; ni++ {
			p1.Register(Step1, ni).SetReceivers()
		}
		p1.Run()

		p2 := parallel.NewParallel()
		for ni := 1; ni <= Program.NumberOfNeos; ni++ {

			p2.Register(Step2, ni).SetReceivers()

		}
		p2.Run()

		createpng(count)

		CurrentStep++

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
	fmt.Println("Time:", time.Since(start))
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
				//ugg++
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
