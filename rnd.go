package main

import (
	"log"
	"math/big"

	"crypto/rand"
)

// randFloat : returns a positive float from 0 to 1
func randFloat() float64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		log.Println(err)
	}
	t := nBig.Int64()
	r := float64(t) / 1000000.0
	return r
}

// randFullRange : Return a float64 in the range of > -1.0 and < +1.0
func randFloatFullValue() float64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		log.Println(err)
	}
	t := nBig.Int64()
	r := ((float64(t) / 1000000.0) * 2.0) - 1.0
	return r
}

// randInt : returns a integer that is between 0 and max.
func randInt(max int) int {
	t := int64(max)
	nBig, err := rand.Int(rand.Reader, big.NewInt(t))
	if err != nil {
		log.Println(err)
	}
	return int(nBig.Int64())
}
