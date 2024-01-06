package chord

import (
	"context"
	"github.com/ryogrid/gord-overlay/pkg/model"
	log "github.com/sirupsen/logrus"
	"net/http"
	_ "net/http/pprof"
	"testing"
)

func runBench(b *testing.B, processCount int) {
	log.SetLevel(log.FatalLevel)
	processes := generateStabilizedProcesses(context.Background(), processCount)
	log.Warn("generateStabilizedProcesses end.")
	findingKey := model.NewHashID("test")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := processes[0].FindSuccessorByTable(context.Background(), findingKey)
		if err != nil {
			b.Fatalf("err = %#v", err)
		}
	}
}

func BenchmarkStabilize(b *testing.B) {
	id1 := model.NewHashID("test1")
	//id2 := model.NewHashID("test2")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//id1.GreaterThanEqual(id2)
		NewFingerTable(id1)
		//NewFingerTable(id2)
	}
}

func BenchmarkFingerTableStabilize_1(b *testing.B) {
	runBench(b, 1)
}

func BenchmarkFingerTableStabilize_2(b *testing.B) {
	runBench(b, 2)
}

func BenchmarkFingerTableStabilize_4(b *testing.B) {
	runBench(b, 4)
}

func BenchmarkFingerTableStabilize_8(b *testing.B) {
	runBench(b, 8)
}

func BenchmarkFingerTableStabilize_16(b *testing.B) {
	runBench(b, 16)
}

func BenchmarkFingerTableStabilize_32(b *testing.B) {
	runBench(b, 32)
}

func BenchmarkFingerTableStabilize_64(b *testing.B) {
	runBench(b, 64)
}

func BenchmarkFingerTableStabilize_128(b *testing.B) {
	runBench(b, 128)
}

func BenchmarkFingerTableStabilize_256(b *testing.B) {
	runBench(b, 256)
}

func BenchmarkFingerTableStabilize_512(b *testing.B) {
	runBench(b, 512)
}

func BenchmarkFingerTableStabilize_1024(b *testing.B) {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	runBench(b, 1024)
}
