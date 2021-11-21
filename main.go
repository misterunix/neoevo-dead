package main

import (
	"fmt"
	"log"
	"math"
	"math/big"

	"crypto/rand"
)

const (
	INCOUNT = 5
	AGE     = 0
	OSC     = 1
	NSBD    = 2
	EWBD    = 3
	BLOCKF  = 4
)

const (
	OUTCOUNT = 2
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
	N        Neuron  // N
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
	W  float64 // W weight
	I  float64 // I input
	O  float64 // O output
	PO float64 // PO previous output
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
	Direction  int // Direction 0 - 8
}

var GeneCount int
var HiddenCount int
var BlipCount int

func main() {
	fmt.Println("Start")

	GeneCount = 2
	HiddenCount = 1
	BlipCount = 3

	var Blips []TinyBlip

	Blips = make([]TinyBlip, BlipCount)

	for index := 0; index < BlipCount; index++ {
		Blips[index].ID = index
		g := Gene{}
		g.V = int(cryptoRandSecure(int64(0xFFFFFFFF)))
		g.Build()
		Blips[index].GeneEncode = append(Blips[index].GeneEncode, g)
	}

	g := Gene{}
	g.V = int(cryptoRandSecure(int64(0xFFFFFFFF)))
	g.Build()

	fmt.Printf("%+v\n", g)

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
