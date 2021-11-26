package main

import (
	"fmt"
)

func printgenes(id int) error {
	if id < 0 || id >= BaseNeoCount {
		return fmt.Errorf("printgenes id '%d' is out of bounds", id)
	}

	fmt.Printf("Neo %d\n", id)
	s := 0
	for _, g := range Neo[id].Genes {
		fmt.Printf("%08X ", g)
		s++
		if s == 8 {
			fmt.Println()
			s = 0
		}
	}
	fmt.Println()

	return nil
}

func printneurons(id int) error {
	if id < 0 || id >= BaseNeoCount {
		return fmt.Errorf("printneurons id '%d' is out of bounds", id)
	}

	for _, n := range Neo[id].Neurons {
		var h1, h2 string

		if n.SrcFlag {
			h1 = "h"
		} else {
			h1 = "i"
		}

		if !n.SnkFlag {
			h2 = "h"
		} else {
			h2 = "o"
		}

		fmt.Printf("id:%d src:%d%s snk:%d%s InV:%2.7f OutV:%2.7f wgt:%2.7f Oinx:%d\n", n.ID, n.Source, h1, n.Sink, h2, n.InValue, n.OutValue, n.Weight, n.OutIndex)
	}
	fmt.Println()

	return nil
}
