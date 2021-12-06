package main

import "fmt"

/*????|PPER|TTAA|QQQQ

E Destination Layer 0-F
R Source Layer 0-F
TT Destination ID 00-FF
AA Source ID 00-FF
QQQQ Weight 65535 : (X / 65535)*8-4 : -4.0 to +4.0*/

func decode(gene uint64) Neuron {

	n := Neuron{}

	weighti := gene & 0x000000000000FFFF
	weightf := ((float64(weighti) / 65535.0) * 8) - 4

	src := (gene & 0x0000000000FF0000) >> 16
	dst := (gene & 0x00000000FF000000) >> 24
	srclayer := (gene & 0x0000000F00000000) >> 28
	dstlayer := (gene & 0x000000F000000000) >> 32

	n.Weight = weightf
	n.Weighti = int(weighti)
	n.SourceID = int(src)
	n.SourceLayer = int(srclayer)
	n.OutID = int(dst)
	n.OutLayer = int(dstlayer)

	return n
}

func genecheck(gene uint64) bool {

	//weighti := gene & 0x000000000000FFFF
	//weightf := ((float64(weighti) / 65535.0) * 8) - 4

	src := (gene & 0x0000000000FF0000) >> 16
	dst := (gene & 0x00000000FF000000) >> 24
	srclayer := (gene & 0x0000000F00000000) >> 28
	dstlayer := (gene & 0x000000F000000000) >> 32

	if dstlayer == 0 {
		return false
	}

	if srclayer == uint64(Program.NumberOfLayers-1) {
		return false
	}

	if srclayer == dstlayer && src == dst {
		return false
	}

	return true
}

func linkneurons(id int) error {

	if id > Program.NumberOfNeos || id < 1 {
		return fmt.Errorf("PrintIO id '%d' is out of bounds", id)
	}

	for layer := 0; layer < Program.NumberOfLayers; layer++ {

		for i, neo := range Neos[id].Neurons {

		}

	}

}
