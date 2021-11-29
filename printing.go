package main

import (
	"fmt"
)

func PrintGenes(id int) error {
	if id >= Program.NumberOfNeos || id < 0 {
		return fmt.Errorf("PrintGenes id '%d' is out of bounds", id)
	}
	//fmt.Println("Printing genes.")
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

func PrintNeuron(id int) error {
	if id >= Program.NumberOfNeos || id < 0 {
		return fmt.Errorf("PrintGenes id '%d' is out of bounds", id)
	}

	for _, n := range Neos[id].Neurons {
		fmt.Printf("%+v \n", n)
	}
	return nil
}
