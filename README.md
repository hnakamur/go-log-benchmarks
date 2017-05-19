go-log-benchmarks
=================

A go logging library benchmarks for me.

## Benchmark result

Note these benchmarks print roughly same outputs, but not the exactly same outputs.

Especially BenchmarkZapLTSVProductionLog uses
[zapcore.EpochTimeEncoder](https://godoc.org/go.uber.org/zap/zapcore#EpochTimeEncoder).
It prints time as floating-point number of seconds since the Unix epoch, and this is a
low cost operation compared to printing time in ISO8609 format.

Other benchmarks (BenchmarkLTSVLog, BenchmarkStandardLog and BenchmarkZapLTSVDevelopmentLog)
uses ISO8609 time format, though there is a slight difference.
BenchmarkZapLTSVDevelopmentLog uses [zapcore.ISO8601TimeEncoder](https://godoc.org/go.uber.org/zap/zapcore#ISO8601TimeEncoder) which prints times with millisecond precision.
The other two prints times with microsecond precision.

```
$ go test -count=10 -bench . -benchmem -cpuprofile=cpu.prof
BenchmarkLTSVLog-2                  	 1000000	      2015 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      2045 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      2036 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      2000 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      2043 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      2011 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      2026 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      2015 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      2046 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      2007 ns/op	      48 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2356 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2340 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	 1000000	      2380 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2334 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2357 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2361 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2381 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2385 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2348 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2406 ns/op	      96 B/op	       3 allocs/op
BenchmarkZapLTSVProductionLog-2     	 3000000	       402 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       345 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       360 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       328 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       341 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       349 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       323 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       317 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       337 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       302 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6404 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6386 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6428 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6368 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6393 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6449 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6354 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6333 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6389 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6334 ns/op	     208 B/op	       3 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2708 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2663 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2657 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2657 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2661 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2662 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2656 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2654 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2655 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2656 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2635 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2648 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2644 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2644 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2645 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2641 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2636 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2638 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2650 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2637 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2629 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2632 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2629 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2648 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2630 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2634 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2640 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2634 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2634 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2632 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/hnakamur/go-log-benchmarks	108.095s
```

## License
MIT
