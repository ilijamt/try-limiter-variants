Limiter
-------

```shell
go test --bench=. -benchmem ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

Benchmark
```shell
go test -run=^$ --bench=. -benchmem -benchtime=10000x . -memprofile memprofile.out -cpuprofile cpuprofile.out
cpu: AMD Ryzen 9 5950X 16-Core Processor
BenchmarkV1_Sync_Take-32                   10000               100.1 ns/op           163 B/op          3 allocs/op
BenchmarkV1_Async_Take-32                  10000             12966 ns/op            7278 B/op        143 allocs/op
BenchmarkV1_Sync_TryTake-32                10000               200.0 ns/op           288 B/op          5 allocs/op
BenchmarkV1_Async_TryTake-32               10000             38434 ns/op           23122 B/op        482 allocs/op
BenchmarkV2_Sync_Take-32                   10000              2124 ns/op            2378 B/op         46 allocs/op
BenchmarkV2_Async_Take-32                  10000             47407 ns/op           28173 B/op        582 allocs/op
BenchmarkV2_Sync_TryTake-32                10000               900.1 ns/op          1661 B/op         31 allocs/op
BenchmarkV2_Async_TryTake-32               10000             35850 ns/op           21130 B/op        437 allocs/op
BenchmarkV3_Sync_Take-32                   10000              2201 ns/op            1721 B/op         31 allocs/op
BenchmarkV3_Async_Take-32                  10000            143054 ns/op           75455 B/op       1569 allocs/op
BenchmarkV3_Sync_TryTake-32                10000              3201 ns/op            1278 B/op         27 allocs/op
BenchmarkV3_Async_TryTake-32               10000            173039 ns/op           86237 B/op       1793 allocs/op
PASS
ok      limiter 22.790s
```