# Golang Benchmark - map, slice, array

```shell
go test -bench=.                                                          
goos: darwin
goarch: arm64
pkg: github.com/Nikolay200669/bench_map-slice
cpu: Apple M1 Pro
BenchmarkArrayAccess-10         1000000000               0.3747 ns/op
BenchmarkSliceAccess-10         1000000000               0.5844 ns/op
BenchmarkMapAccess-10           206825692                7.128 ns/op
PASS
ok      github.com/Nikolay200669/bench_map-slice        3.637s
```