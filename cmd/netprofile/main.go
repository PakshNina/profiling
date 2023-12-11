package main

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

type Dict struct {
	dict1 map[int]time.Time
	dict2 map[int]int
	dict3 map[time.Time]int
}

func main() {
	var mu sync.Mutex
	dict := Dict{
		dict1: map[int]time.Time{},
		dict2: map[int]int{},
		dict3: map[time.Time]int{},
	}
	ctx := context.Background()
	pprof.Do(ctx, pprof.Labels("label", "dict"), func(ctx context.Context) {
		go func() {
			for {
				for i := 0; i < 10; i++ {
					go func(i int) {
						mu.Lock()
						dict.dict1[i] = time.Now()
						dict.dict2[i] = i
						dict.dict3[time.Now()] = i
						time.Sleep(time.Millisecond * 100)
						mu.Unlock()
					}(i)
				}
				time.Sleep(time.Millisecond * 100)
			}
		}()
	})
	// Анализ блокирующих функций, 0 - отключает профилирование, 1 - отслеживание всех блокирующих функций
	runtime.SetBlockProfileRate(1)
	// Анализ мьютексов
	runtime.SetMutexProfileFraction(1)

	// Запуск веб-профайлинга
	// http://localhost:5555/debug/pprof/
	log.Fatal(http.ListenAndServe("localhost:5555", nil))
}
