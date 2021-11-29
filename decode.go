package main

import "fmt"

func decode(g int) Neuron {

	weightI := (g & 0x0000FFFF)
	weight := (float64(weightI)/65535.0)*8 - 4.0
	layer := (g & 0x000F0000 >> 16) % Program.NumberOfLayers
	node := (g & 0x0FF00000 >> 24) % Program.NumberOfNeurons

	tn := Neuron{
		Weight: weight,
		Layer:  layer,
		Out:    node,
	}

	return tn
}

func buildGenes(id int) error {
	if id >= Program.NumberOfNeos || id < 0 {
		return fmt.Errorf("buildNeurons id '%d' is out of bounds", id)
	}

	if len(Neos[id].Genes) == 0 {
		tg := randInt(0xFFFFFFFF)
		n1 := decode(tg)
		Neos[id].Neurons = append(Neos[id].Neurons, n1)
		Neos[id].Genes = append(Neos[id].Genes, tg)
	}

	for {
		tg := randInt(0xFFFFFFFF)

		n1 := decode(tg)
		drop := false
		for _, j := range Neos[id].Genes {
			n2 := decode(j)

			if n1.Layer != n2.Layer && n1.Out != n2.Out {
				drop = true
				break
			}
		}
		if drop {
			Neos[id].Neurons = append(Neos[id].Neurons, n1)
			Neos[id].Genes = append(Neos[id].Genes, tg)
			break
		}
	}

	return nil
}
