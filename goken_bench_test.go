package goken_test

import (
	"testing"

	"github.com/kpango/goken"
)

func BenchmarkGokenGenerate(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			goken.GenerateToken("key", "userID", "actionID")
			goken.GenerateToken("key", "userID", "actionID")
			goken.GenerateToken("key", "userID", "actionID")
			goken.GenerateToken("key", "userID", "actionID")
			goken.GenerateToken("key", "userID", "actionID")
		}
	})
}

func BenchmarkXSRFGenerate(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			goken.GenerateTokenAtTime("key", "userID", "actionID")
			goken.GenerateTokenAtTime("key", "userID", "actionID")
			goken.GenerateTokenAtTime("key", "userID", "actionID")
			goken.GenerateTokenAtTime("key", "userID", "actionID")
			goken.GenerateTokenAtTime("key", "userID", "actionID")
		}
	})
}
