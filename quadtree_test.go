package rquad

import (
	"image"
	"testing"

	"github.com/aurelien-rainone/imgtools/imgscan"
)

func benchmarkQuadtreeCreation(b *testing.B, pngfile string, fn newQuadtreeFunc, resolution int) {
	var (
		bm      image.Image
		err     error
		scanner imgscan.Scanner
	)

	bm, err = loadPNG(pngfile)
	checkB(b, err)
	scanner, err = imgscan.NewScanner(bm)
	checkB(b, err)

	// run N times
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := fn(scanner, resolution)
		checkB(b, err)
	}
}

func BenchmarkBasicTreeCreationRes32(b *testing.B) {
	benchmarkQuadtreeCreation(b, "./testdata/bigsquare.png", newBasicTree, 32)
}

func BenchmarkBasicTreeCreationRes16(b *testing.B) {
	benchmarkQuadtreeCreation(b, "./testdata/bigsquare.png", newBasicTree, 16)
}

func BenchmarkBasicTreeCreationRes8(b *testing.B) {
	benchmarkQuadtreeCreation(b, "./testdata/bigsquare.png", newBasicTree, 8)
}

func BenchmarkBasicTreeCreationRes4(b *testing.B) {
	benchmarkQuadtreeCreation(b, "./testdata/bigsquare.png", newBasicTree, 4)
}

func BenchmarkBasicTreeCreationRes2(b *testing.B) {
	benchmarkQuadtreeCreation(b, "./testdata/bigsquare.png", newBasicTree, 2)
}

func BenchmarkBasicTreeCreationRes1(b *testing.B) {
	benchmarkQuadtreeCreation(b, "./testdata/bigsquare.png", newBasicTree, 1)
}

func BenchmarkCNTreeCreationRes32(b *testing.B) {
	benchmarkQuadtreeCreation(b, "./testdata/bigsquare.png", newCNTree, 32)
}

func BenchmarkCNTreeCreationRes16(b *testing.B) {
	benchmarkQuadtreeCreation(b, "./testdata/bigsquare.png", newCNTree, 16)
}

func BenchmarkCNTreeCreationRes8(b *testing.B) {
	benchmarkQuadtreeCreation(b, "./testdata/bigsquare.png", newCNTree, 8)
}

func BenchmarkCNTreeCreationRes4(b *testing.B) {
	benchmarkQuadtreeCreation(b, "./testdata/bigsquare.png", newCNTree, 4)
}

func BenchmarkCNTreeCreationRes2(b *testing.B) {
	benchmarkQuadtreeCreation(b, "./testdata/bigsquare.png", newCNTree, 2)
}

func BenchmarkCNTreeCreationRes1(b *testing.B) {
	benchmarkQuadtreeCreation(b, "./testdata/bigsquare.png", newCNTree, 1)
}
