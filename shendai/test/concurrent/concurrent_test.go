package concurrent

import "testing"

func BenchmarkConcurrentAtomicAdd(b *testing.B) {
	b.ResetTimer()
	println(b.N)
	for i := 0; i < b.N; i++ {
		ConcurrentAtomicAdd()
	}
}

func BenchmarkConcurrentMutexAdd(b *testing.B) {
	b.ResetTimer()
	println(b.N)
	for i := 0; i < b.N; i++ {
		ConcurrentMutexAdd()
	}
}
