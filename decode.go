package main

import "fmt"

func checkgene(gene uint32) bool {

	//fmt.Printf("checkgene %08X\n", gene)

	var src, dst uint32
	var srcl, dstl uint32

	srcl = ((gene & 0xE0000000) >> 29) % uint32(Program.NumberOfLayers)
	dstl = ((gene & 0x00E00000) >> 21) % uint32(Program.NumberOfLayers)

	// dst layer can not be layer 0
	if dstl == 0 {
		return false
	}

	// src layer can not be the output layer
	if srcl == uint32(Program.NumberOfLayers-1) {
		return false
	}

	// src layer cannot be more than dst layer (cant go backwards - at least for now.
	if srcl > dstl {
		return false
	}

	//fmt.Printf("srcl:%02d dstl:%02d\n", srcl, dstl)

	if srcl == 0 {
		src = ((gene & 0x1F000000) >> 24) % INPUTCOUNT
	} else {
		src = ((gene & 0x1F000000) >> 24) % uint32(Program.NumberOfNeurons)
	}

	if dstl == uint32(Program.NumberOfLayers-1) {
		dst = ((gene & 0x001F0000) >> 16) % OUTPUTCOUNT
	} else {
		dst = ((gene & 0x001F0000) >> 16) % uint32(Program.NumberOfNeurons)
	}

	// no loops on same neuron
	if srcl != uint32(Program.NumberOfLayers)-1 {
		if srcl == dstl && src == dst {
			return false
		}
	}
	//fmt.Printf("srcl:%02d src:%02d dstl:%02d dst:%02d\n", srcl, src, dstl, dst)
	return true
}

func decode(gene uint32) Neuron {

	neu := Neuron{}

	weightI := (gene & 0x0000FFFF)               // weightI : unsigned integer value of the weight
	weight := (float64(weightI)/65535.0)*8 - 4.0 // weight : neurons weight -4.0 to +4.0

	var src, dst uint32
	var srcl, dstl uint32

	srcl = ((gene & 0xE0000000) >> 29) % uint32(Program.NumberOfLayers)
	dstl = ((gene & 0x00E00000) >> 21) % uint32(Program.NumberOfLayers)

	if srcl == 0 {
		src = ((gene & 0x1F000000) >> 24) % INPUTCOUNT
	} else {
		src = ((gene & 0x1F000000) >> 24) % uint32(Program.NumberOfNeurons)
	}

	if dstl == uint32(Program.NumberOfLayers-1) {
		dst = ((gene & 0x001F0000) >> 16) % OUTPUTCOUNT
	} else {
		dst = ((gene & 0x001F0000) >> 16) % uint32(Program.NumberOfNeurons)
	}

	/*
		srcl = ((gene & 0xE0000000) >> 29) % uint32(Program.NumberOfLayers)
		dstl = ((gene & 0x00E00000) >> 21) % uint32(Program.NumberOfLayers)

		if srcl == 0 {
			src = ((gene & 0x1F000000) >> 24) % INPUTCOUNT
		} else {
			src = ((gene & 0x1F000000) >> 24) % uint32(Program.NumberOfNeurons)
			//fmt.Printf("--src--- %d %d %d -----\n", srcl, src, Program.NumberOfNeurons)
		}

		if dstl == uint32(Program.NumberOfLayers-1) {
			dst = ((gene & 0x001F0000) >> 16) % OUTPUTCOUNT
		} else {
			dst = ((gene & 0x001F0000) >> 16) % uint32(Program.NumberOfNeurons)
			//fmt.Printf("--dst--- %d %d %d-----\n", dstl, dst, Program.NumberOfNeurons)
		}
	*/
	//	fmt.Printf("srcl:%02d src:%02d dstl:%02d dst:%02d\n", srcl, src, dstl, dst)
	neu.Weight = weight
	neu.SourceID = int(src)
	neu.SourceLayer = int(srcl)
	neu.OutID = int(dst)
	neu.OutLayer = int(dstl)

	return neu
}

// link the neurns for easier navigation
func linkneurons(id int) error {

	if id < 1 || id > Program.NumberOfNeos {
		return fmt.Errorf("linkneurons id `%d` out of bounds", id)
	}

	for i, n := range Neos[id].Neurons {

		/*
			if n.OutLayer == Program.NumberOfLayers-1 {
				Neos[id].Neurons[i].LinkForward = -99
				continue
			}
		*/
		for j, k := range Neos[id].Neurons {

			// Cant do itself. :)
			if i == j {
				continue
			}

			if n.OutLayer == k.SourceLayer && n.OutID == k.SourceID {

				Neos[id].Neurons[i].LinkForward = j
				//fmt.Println("Link", id, i, j)
				break

			} else {
				Neos[id].Neurons[i].LinkForward = -1
			}

		}

	}

	return nil

}
