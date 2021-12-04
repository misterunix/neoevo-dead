package main

import (
	"fmt"
)

// PrintGenes : Print the genes of the Neo in hex format
func PrintGenes(id int) error {
	if id > Program.NumberOfNeos || id < 1 {
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

// PrintIO : Prints the Neos Iinput and Output slices.
func PrintIO(id int) error {
	if id > Program.NumberOfNeos || id < 1 {
		return fmt.Errorf("PrintIO id '%d' is out of bounds", id)
	}
	fmt.Println("Neo", id)
	for i := 0; i < Program.NumberOfInputs; i++ {
		fmt.Printf("%5.3f ", Neos[id].Inputs[i])
	}
	fmt.Println()
	for i := 0; i < Program.NumberOfOutputs; i++ {
		fmt.Printf("%5.3f ", Neos[id].Outputs[i])
	}
	fmt.Println()
	return nil
}

// PrintNeuron : Prints the Neos Neurons. Just a dump of the slice.
func PrintNeuron(id int) error {
	if id > Program.NumberOfNeos || id < 1 {
		return fmt.Errorf("PrintGenes id '%d' is out of bounds", id)
	}

	for _, n := range Neos[id].Neurons {
		fmt.Printf("NEO %d %+v \n", id, n)
	}
	return nil
}

// PrintNet : Print the Neo's net list.
func PrintNet(id int) error {
	if id > Program.NumberOfNeos || id < 1 {
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
				s1 = "pES"
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
			s1 = fmt.Sprintf("%d-%d", n.SourceLayer, n.Source)
		}

		if n.OutLayer == Program.NumberOfLayers-1 {
			switch n.Out {
			case 0:
				s2 = "mRD"
			case 1:
				s2 = "mFW"
			case 2:
				s2 = "mBK"
			case 3:
				s2 = "tLF"
			case 4:
				s2 = "tRT"
			case 5:
				s2 = "mNT"
			case 6:
				s2 = "mST"
			case 7:
				s2 = "mWS"
			case 8:
				s2 = "mES"
			case 9:
				s2 = "NOP"
			case 10:
				s2 = "mX"
			case 11:
				s2 = "mY"
			}
		} else {
			s2 = fmt.Sprintf("%d-%d", n.OutLayer, n.Out)
		}
		w1 := n.Weight * 8192
		w := int(w1)

		fmt.Printf("%s %s %d\n", s1, s2, w)
	}
	return nil
}
