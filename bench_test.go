package box_test

import (
	ƒ "github.com/tedsuo/box"
	"testing"
)

func BenchmarkSetNativeMapFromLoop(b *testing.B) {
	iMap := map[int]int{}
	for i := 0; i < b.N; i++ {
		iMap[i] = i
	}
}

func BenchmarkSetBoxedMapFromLoop(b *testing.B) {
	mapp := ƒ.NewMap()
	for i := 0; i < b.N; i++ {
		mapp.Set(i, i)
	}
}

func BenchmarkSetBoxedMapFromSeq(b *testing.B) {
	seq := ƒ.NewSeq()
	go func() {
		defer close(seq)
		for i := 0; i < b.N; i++ {
			seq <- ƒ.Box{i, i}
		}
	}()
	mapp := ƒ.NewMap(seq)
	mapp = mapp
}

func BenchmarkSetMapFromNativeMap(b *testing.B) {
	iMap := map[int]int{}
	for i := 0; i < b.N; i++ {
		iMap[i] = i
	}
	b.ResetTimer()
	mapp := ƒ.NewMap(iMap)
	mapp = mapp
}

func BenchmarkCountMapFromNativeMap(b *testing.B) {
	iMap := map[int]int{}
	for i := 0; i < b.N; i++ {
		iMap[i] = i
	}
	b.ResetTimer()
	mapp := ƒ.NewMap(iMap)
	ƒ.Count(mapp)
}

func BenchmarkCountSeqFromNativeMap(b *testing.B) {
	iMap := map[int]int{}
	for i := 0; i < b.N; i++ {
		iMap[i] = i
	}
	b.ResetTimer()
	seq := ƒ.NewSeq(iMap)
	ƒ.Count(seq)
}

func BenchmarkCountSeqViaEach(b *testing.B) {
	seq := ƒ.NewSeq()
	go func() {
		defer close(seq)
		for i := 0; i < b.N; i++ {
			seq <- ƒ.Box{i}
		}
	}()
	ƒ.Count(seq)
}

func BenchmarkCountSeqViaRangeInc(b *testing.B) {
	seq := ƒ.NewSeq()
	go func() {
		defer close(seq)
		for i := 0; i < b.N; i++ {
			seq <- ƒ.Box{i}
		}
	}()
	i := length(seq)
	i++
}

func BenchmarkFillNativeMapViaLoop(b *testing.B) {
	iMap := map[int]int{}
	for i := 0; i < b.N; i++ {
		iMap[i] = i
	}
}

func length(iMap ƒ.Sequence) (i int) {
	for _ = range iMap {
		i++
	}
	return
}
