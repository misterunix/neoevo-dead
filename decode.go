package main

import "fmt"

func decode(g int) Neuron {

	weightI := (g & 0x0000FFFF)
	weight := (float64(weightI)/65535.0)*8 - 4.0
	layer := (g & 0x000F0000 >> 16) & Program.NumberOfLayers
	node := (g & 0x0FF00000 >> 20) & Program.NumberOfNeurons

	tn := Neuron{
		Weight: weight,
		Layer:  layer,
		Out:    node,
	}

	return tn
}

func buildNeurons(id int) error {
	if id >= Program.NumberOfNeos || id < 0 {
		return fmt.Errorf("buildNeurons id '%d' is out of bounds", id)
	}

	for _, g := range Neos[id].Genes {

		n1 := decode(g)

	}

	return nil

}
