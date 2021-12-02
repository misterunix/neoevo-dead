package main

import "fmt"

// decode :
func decode(g int) Neuron {

	weightI := (g & 0x0000FFFF)
	weight := (float64(weightI)/65535.0)*8 - 4.0

	source := (g & 0xFF000000 >> 24) & 0x1F
	sourcelayer := (((g & 0xFF000000 >> 24) & 0xE0) >> 5) % Program.NumberOfLayers
	out := (g & 0x00FF0000 >> 16) & 0x1F
	outlayer := (((g & 0x00FF0000 >> 16) & 0xE0) >> 5) % Program.NumberOfLayers

	//node := (g & 0x0FF00000 >> 24) % Program.NumberOfNeurons

	if sourcelayer == 0 {
		source %= Program.NumberOfInputs
	} else {
		source %= Program.NumberOfNeurons
	}

	if outlayer == Program.NumberOfLayers-1 {
		out %= Program.NumberOfOutputs
	} else {
		out %= Program.NumberOfNeurons
	}

	tn := Neuron{
		Weight:      weight,
		Source:      source,
		SourceLayer: sourcelayer,
		Out:         out,
		OutLayer:    outlayer,
	}

	return tn
}

// buildGenes :
func buildGenes(id int) error {
	if id > Program.NumberOfNeos || id < 1 {
		return fmt.Errorf("buildNeurons id '%d' is out of bounds", id)
	}

	if len(Neos[id].Genes) == 0 {
		for {

			tg := randInt(0xFFFFFFFF)
			n1 := decode(tg)

			/*
				if n1.SourceLayer == 0 {
					n1.Source %= Program.NumberOfInputs
				} else {
					n1.Source %= Program.NumberOfNeurons
				}

				if n1.OutLayer == Program.NumberOfLayers-1 {
					n1.Out %= Program.NumberOfOutputs
				} else {
					n1.Out %= Program.NumberOfNeurons
				}
			*/
			if n1.SourceLayer == n1.OutLayer {
				continue
			}
			if n1.SourceLayer >= n1.OutLayer {
				continue
			}

			Neos[id].Neurons = append(Neos[id].Neurons, n1)
			Neos[id].Genes = append(Neos[id].Genes, tg)
			break
		}
	}

	for {
		tg := randInt(0xFFFFFFFF)

		n1 := decode(tg)
		/*
			if n1.SourceLayer == 0 {
				n1.Source %= Program.NumberOfInputs
			} else {
				n1.Source %= Program.NumberOfNeurons
			}

			if n1.OutLayer == Program.NumberOfLayers-1 {
				n1.Out %= Program.NumberOfOutputs
			} else {
				n1.Out %= Program.NumberOfNeurons
			}
		*/
		//n1.SourceLayer %= Program.NumberOfLayers
		//n1.OutLayer %= Program.NumberOfLayers

		if n1.SourceLayer == n1.OutLayer {
			continue
		}
		if n1.SourceLayer >= n1.OutLayer {
			continue
		}
		drop := false
		for _, j := range Neos[id].Genes {
			n2 := decode(j)
			/*
				if n2.SourceLayer == 0 {
					n2.Source %= Program.NumberOfInputs
				} else {
					n2.Source %= Program.NumberOfNeurons
				}

				if n2.OutLayer == Program.NumberOfLayers-1 {
					n2.Out %= Program.NumberOfOutputs
				} else {
					n2.Out %= Program.NumberOfNeurons
				}
			*/
			//n2.SourceLayer %= Program.NumberOfLayers
			//n2.OutLayer %= Program.NumberOfLayers
			if n1.SourceLayer != n2.SourceLayer {
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
