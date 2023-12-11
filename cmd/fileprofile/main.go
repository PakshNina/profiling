package main

import (
	"context"
	"log"
	"os"
	"runtime/pprof"

	"profiling/factorial"
)

func main() {
	// CPU профайлинг
	cpuFile, err := os.Create("cpu1.prof")
	if err != nil {
		log.Fatal(err)
	}
	err = pprof.StartCPUProfile(cpuFile)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		pprof.StopCPUProfile()
	}()

	// Запуск кода с лейблом
	ctx := context.Background()
	pprof.Do(ctx, pprof.Labels("label", "calculate"), func(ctx context.Context) {
		factorial.Calculate()
	})

	// Запись профайла heap
	memFile, err := os.Create("mem1.prof")
	if err != nil {
		log.Fatal(err)
	}
	err = pprof.WriteHeapProfile(memFile)
	if err != nil {
		log.Fatal(err)
	}
}
