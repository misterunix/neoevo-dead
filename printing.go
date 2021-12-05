package main

import "fmt"

// printgenes : Display the genes in 8 digit hex format
func printgenes(id int) error {
	if id > Program.NumberOfNeos || id < 1 {
		return fmt.Errorf("printgenes id '%d' is out of bounds", id)
	}
	fmt.Println("Neo", id)
	var w int
	for i := 0; i < Program.NumberOfGenes; i++ {
		w++
		fmt.Printf("%08X ", Neos[id].Genes[i])
		if w == 8 {
			fmt.Println()
		}
	}
	fmt.Println()
	return nil
}

// printneuron : Prints the Neos Neurons. Just a dump of the slice.
func printneuron(id int) error {
	if id > Program.NumberOfNeos || id < 1 {
		return fmt.Errorf("printneuron id '%d' is out of bounds", id)
	}

	for i, n := range Neos[id].Neurons {
		//fmt.Printf("NEO %d Neuron %d ", id, i)
		//fmt.Printf("ID:%d inv:%4.3f outv:%4.3f ", n.ID, n.InValue, n.OutValue)
		//fmt.Printf("srcl:%02d src:%02d dstl:%02d dst:%02d ", n.SourceLayer, n.SourceID, n.OutLayer, n.OutID)
		//	fmt.Printf("lf:%02d lb:%02d\n", n.LinkForward, n.LinkBackward)
		fmt.Printf("NEO %d Neuron %d %+v \n", id, i, n)
	}

	return nil
}
