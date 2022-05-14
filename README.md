Limiter
-------

### Testing

```shell
$ ./test.sh
```

### Benchmark

#### Windows - 64GB memory - 16-core

CPU and Memory profiles in [profiling/win](profiling/win)

```shell
$ ./bench.sh
goos: windows
goarch: amd64
pkg: limiter
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

* [CPU Profile](profiling/win/pprof001.svg)
* [Mem Profile](profiling/win/pprof002.svg)

#### Linux - 16GB memory - 4-core

CPU and Memory profiles in [profiling/linux](profiling/linux)

```shell
$ ./bench.sh
goos: linux
goarch: amd64
pkg: limiter
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkV1_Sync_Take-8       	   10000	      1488 ns/op	     656 B/op	       5 allocs/op
BenchmarkV1_Async_Take-8      	   10000	    191669 ns/op	  127361 B/op	    2640 allocs/op
BenchmarkV1_Sync_TryTake-8    	   10000	       555.5 ns/op	     557 B/op	      11 allocs/op
BenchmarkV1_Async_TryTake-8   	   10000	     77459 ns/op	   48701 B/op	    1013 allocs/op
BenchmarkV2_Sync_Take-8       	   10000	      9981 ns/op	    5516 B/op	     111 allocs/op
BenchmarkV2_Async_Take-8      	   10000	    307409 ns/op	  179575 B/op	    3728 allocs/op
BenchmarkV2_Sync_TryTake-8    	   10000	      6154 ns/op	    4523 B/op	      90 allocs/op
BenchmarkV2_Async_TryTake-8   	   10000	    169320 ns/op	   97750 B/op	    2033 allocs/op
BenchmarkV3_Sync_Take-8       	   10000	      6926 ns/op	    2958 B/op	      58 allocs/op
BenchmarkV3_Async_Take-8      	   10000	   1397216 ns/op	  639131 B/op	   13287 allocs/op
BenchmarkV3_Sync_TryTake-8    	   10000	      1322 ns/op	     441 B/op	       9 allocs/op
BenchmarkV3_Async_TryTake-8   	   10000	    844289 ns/op	  337461 B/op	    7029 allocs/op
PASS
ok  	limiter	40.339s                                                                                                                                                                                                      
```

* [CPU Profile](profiling/linux/pprof001.svg)
* [Mem Profile](profiling/linux/pprof002.svg)
