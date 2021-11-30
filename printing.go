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

func PrintNet(id int) error {
	if id >= Program.NumberOfNeos || id < 0 {
		return fmt.Errorf("PrintNet id '%d' is out of bounds", id)
	}

	for _, n := range Neos[id].Neurons {
		var s1 string
		var s2 string
		if n.SourceLayer == 0 {
			switch n.Source {
			case 0:
				s1 = "age"
			case 1:
				s1 = "clF"
			case 2:
				s1 = "clN"
			case 3:
				s1 = "pNS"
			case 4:
				s1 = "pWE"
			case 5:
				s1 = "dsF"
			case 6:
				s1 = "dsN"
			case 7:
				s1 = "Hgr"
			case 8:
				s1 = "dFB"
			case 9:
				s1 = "dBB"
			}
		} else {
			s1 = fmt.Sprintf("N%02d-%02d", n.SourceLayer, n.Source)
		}

		if n.OutLayer == Program.NumberOfLayers-1 {
			s2 = fmt.Sprintf("O-%02d", n.Out)
		} else {
			s2 = fmt.Sprintf("N%02d-%02d", n.OutLayer, n.Out)
		}
		w1 := n.Weight * 8192
		w := int(w1)

		fmt.Printf("%s %s %d\n", s1, s2, w)
	}
	return nil
}
