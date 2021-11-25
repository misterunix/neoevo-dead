package main

type Neuron struct {
	ID       int
	Source   int // Source : ID of the neuron
	Sink     int
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
}

var Neo []NeoBase
var BaseNeoCount int
var Generation int

func main() {

	BaseNeoCount = 10
	genecount := 8

	Neo = make([]NeoBase, BaseNeoCount)
	for i := 0; i < BaseNeoCount; i++ {
		Neo[i].Genes = make([]int, genecount)
	}

	if Generation == 0 {
		for i := 0; i < BaseNeoCount; i++ {
			for j := 0; j < genecount; j++ {
				Neo[i].Genes[j] = randInt(0xFFFFFFFF)
			}
		}
	}

}
