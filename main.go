package main

func main() {

	Program.NumberOfNeos = 10
	Program.NumberOfSteps = 300
	Program.NumberOfLayers = 3
	Program.NumberOfSteps = 8

	initNeos()

}

func initNeos() {

	Neos = make([]Neo, Program.NumberOfNeos)

	for i := 0; i < Program.NumberOfNeos; i++ {

		Neos[i].Neurons = make(map[int]Neuron)
		Neos[i].Genes = make([]int, Program.NumberOfGenes)
		for j := 0; j < Program.NumberOfGenes; j++ {
			Neos[i].Genes = append(Neos[i].Genes, randInt(0xFFFFFFFF))
		}
	}

}
