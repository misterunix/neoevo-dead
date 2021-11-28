package main

func main() {

	Program.MaxNeos = 10

}

func initNeos() {

	Neos = make([]Neo, Program.MaxNeos)

	for i := 0; i < Program.MaxNeos; i++ {

		Neos[i].Neurons = make(map[int]Neuron)

	}

}
