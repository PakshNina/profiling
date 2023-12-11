## Бенчмаркинг
1. Запуск бенчмарков

`go test -v ./... -bench=. -benchtime 100x -count=2 -benchmem`

3. Сравнение бенчмарков

`go test ./... -bench="BenchmarkCalculate" -run=^# -count=2 | tee old.txt`

`go test ./... -bench="BenchmarkCalculate" -run=^# -count=2 | tee new.txt`

Сравниваем:

`benchstat old.txt new.txt`

3. Профилирование с помощью pprof для памяти и ЦПУ (из пакета)

`go test -bench='BenchmarkCalculate' -cpuprofile='cpu.prof' -memprofile='mem.prof'`

Профилирование ЦПУ:

`go tool pprof cpu.prof`

`go tool pprof -http=:6060 cpu.prof`

Профилирование памяти:

`go tool pprof mem.prof`

`go tool pprof -http=:6060 mem.prof`

# Профилирование

1. Сравнить два варианта:

`go tool pprof -http=:6060 -diff_base cpu1.prof cpu.prof`

2. Диагностика утечки памяти.

Получение данных о куче с вызовом сборщика мусора

`http://localhost:5555/debug/pprof/heap?gc=1`

`go tool pprof -http=:6060 -diff_base heap1.prof heap2.prof`

3. CPU

`go tool pprof -http=:6061 http://localhost:5555/debug/pprof/profile`

4. Heap

`go tool pprof -http=:6061 http://localhost:5555/debug/pprof/heap`

5. Мьютексы

`go tool pprof -http=:6061 http://localhost:5555/debug/pprof/mutex`

6. Блокирующие функции

`go tool pprof -http=:6061 http://localhost:5555/debug/pprof/block` 
