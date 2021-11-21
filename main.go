package main

import (
	"fmt"
	"log"
	"math"
	"math/big"

	"crypto/rand"
)

const (
	INCOUNT = 10
	AGE     = 0
	OSC     = 1
	NSBD    = 2
	EWBD    = 3
	BLOCKF  = 4
)

const (
	OUTCOUNT = 10
	OMF      = 0
	MLR      = 1
)

type Gene struct {
	V        int
	Source   bool    // Source
	Sink     bool    // Sink
	SourceID int     // SourceID
	SinkID   int     // SinkID
	WeightI  int     // WeightI
	Weight   float64 // Weight

}

func (g *Gene) Build() {
	g.WeightI = (g.V & 0x0000FFFF)
	g.Weight = (float64(g.WeightI)/65535.0)*2 - 1.0

	g.SinkID = (g.V & 0x00FF0000 >> 16) & 0x7f
	t := ((g.V & 0x00FF0000 >> 16) & 0x80) >> 7
	if t == 0 {
		g.Sink = false
	} else {
		g.Sink = true
	}
	g.SourceID = (g.V & 0xFF000000 >> 24) & 0x7f
	t = ((g.V & 0xFF000000 >> 24) & 0x80) >> 7
	if t == 0 {
		g.Source = false
	} else {
		g.Source = true
	}
}

type Neuron struct {
	ID int     // ID
	W  float64 // W weight
	I  float64 // I input
	O  float64 // O output
	PO float64 // PO previous output
	C  []int   // C connections from other neurons
}

func (n *Neuron) Output() {

	ti := n.I + n.PO
	n.O = math.Tanh(n.W * ti)
	n.PO = n.O
}

type TinyBlip struct {
	ID         int
	GeneEncode []Gene
	In         [INCOUNT]float64
	Out        [OUTCOUNT]float64
	N          []Neuron // N
	Direction  int      // Direction 0 - 8
}

func (tb *TinyBlip) Step() {
	var ii float64
	for i, tn := range tb.N {
		ii = 0
		for ti := range tn.C {
			ii += tb.N[ti].O
		}
		tb.N[i].I = ii
	}

}

var GeneCount int
var HiddenCount int
var BlipCount int
var Blips []TinyBlip

func main() {
	fmt.Println("Start")

	GeneCount = 2
	HiddenCount = 1
	BlipCount = 3

	Blips = make([]TinyBlip, BlipCount)

	for index := 0; index < BlipCount; index++ {
		Blips[index].ID = index
		Blips[index].GeneEncode = make([]Gene, GeneCount)
		/*
			g := Gene{}
			g.V = int(cryptoRandSecure(int64(0xFFFFFFFF)))
			g.Build()
		*/

		Blips[index].N = make([]Neuron, INCOUNT+HiddenCount)
		//Blips[index].GeneEncode = append(Blips[index].GeneEncode, g)
	}

	for index := 0; index < BlipCount; index++ {
		for j := 0; j < GeneCount; j++ {
			g := Gene{}
			g.V = int(cryptoRandSecure(int64(0xFFFFFFFF)))
			g.Build()
			Blips[index].GeneEncode[j] = g

		}
	}

	for index := 0; index < BlipCount; index++ {
		fmt.Printf("%d ", index)
		for j := 0; j < GeneCount; j++ {
			fmt.Printf("%08X ", Blips[index].GeneEncode[j].V)
		}
		fmt.Println()
	}

	/*
		count := 4

		i := make([]float64, count)
		for j := 0; j < count; j++ {
			i[j] = craprand()
		}

		m := make([]float64, count)
		for j := 0; j < count; j++ {
			m[j] = (craprand() * 8.0) - 4.0
		}

		s := make([]float64, count)
		for j := 0; j < count; j++ {
			s[j] = i[j] * m[j]
		}

		t := make([]float64, count)
		for j := 0; j < count; j++ {
			t[j] = math.Tanh(s[j])
		}

		//

		for j := 0; j < count; j++ {
			fmt.Printf("%6.3f ", i[j])
		}
		fmt.Println()
		for j := 0; j < count; j++ {
			fmt.Printf("%6.3f ", m[j])
		}
		fmt.Println()
		for j := 0; j < count; j++ {
			fmt.Printf("%6.3f ", s[j])
		}
		fmt.Println()
		for j := 0; j < count; j++ {
			fmt.Printf("%6.3f ", t[j])
		}
		fmt.Println()
	*/
	/*
		ma := []float64{1, 2, 3, 0, 4, 5, 0, 0, 6}
		mb := []float64{0, 1, 1, 1, 1, 1, 1, 1, 2}
		mm := []float64{0.5, 1.5, 2, 2.5, 3, 3.5, 4, -4, -2}

		a := mat.NewDense(3, 3, ma)
		b := mat.NewDense(3, 3, mb)
		c := mat.NewDense(3, 3, nil)
		m := mat.NewDense(3, 3, mm)

		fa := mat.Formatted(a, mat.Prefix("    "), mat.Squeeze())
		fb := mat.Formatted(b, mat.Prefix("    "), mat.Squeeze())
		fm := mat.Formatted(m, mat.Prefix("    "), mat.Squeeze())

		fmt.Printf("with all values:\na = %v\n\n", fa)
		fmt.Printf("with all values:\na = %v\n\n", fb)
		fmt.Printf("with all values:\na = %v\n\n", fm)

		c.MulElem(a, m)
		fc := mat.Formatted(c, mat.Prefix("    "), mat.Squeeze())
		fmt.Printf("with all values:\na = %v\n\n", fc)
	*/

}

func craprand() float64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		log.Println(err)
	}
	t := nBig.Int64()
	r := float64(t) / 1000000.0
	return r
}

func cryptoRandSecure(max int64) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		log.Println(err)
	}
	return nBig.Int64()
}
