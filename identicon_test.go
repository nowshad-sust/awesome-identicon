package main

import (
	"testing"

	"github.com/Pallinder/go-randomdata"
)

func BenchmarkGenerateIdenticon(b *testing.B) {

	for i := 0; i < 10; i++ {
		generateIdenticon(randomdata.SillyName(), 60)
	}
}
