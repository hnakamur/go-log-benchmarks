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
BenchmarkLTSVLog-2                 	 1000000	      2036 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                 	 1000000	      2028 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                 	 1000000	      2046 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                 	 1000000	      2049 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                 	 1000000	      2035 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                 	 1000000	      2061 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                 	 1000000	      2049 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                 	 1000000	      2041 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                 	 1000000	      2046 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                 	 1000000	      2043 ns/op	      48 B/op	       3 allocs/op
BenchmarkStandardLog-2             	  500000	      2408 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2             	  500000	      2412 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2             	  500000	      2448 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2             	  500000	      2434 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2             	  500000	      2400 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2             	  500000	      2421 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2             	  500000	      2431 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2             	  500000	      2449 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2             	  500000	      2432 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2             	  500000	      2386 ns/op	      96 B/op	       3 allocs/op
BenchmarkZapLTSVProductionLog-2    	 3000000	       415 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2    	 5000000	       412 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2    	 3000000	       419 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2    	 5000000	       398 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2    	 5000000	       374 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2    	 3000000	       365 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2    	 5000000	       419 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2    	 3000000	       414 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2    	 5000000	       398 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2    	 3000000	       404 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVDevelopmentLog-2   	  200000	      6123 ns/op	     197 B/op	       4 allocs/op
BenchmarkZapLTSVDevelopmentLog-2   	  200000	      6104 ns/op	     197 B/op	       4 allocs/op
BenchmarkZapLTSVDevelopmentLog-2   	  200000	      6145 ns/op	     197 B/op	       4 allocs/op
BenchmarkZapLTSVDevelopmentLog-2   	  200000	      6169 ns/op	     197 B/op	       4 allocs/op
BenchmarkZapLTSVDevelopmentLog-2   	  200000	      6149 ns/op	     197 B/op	       4 allocs/op
BenchmarkZapLTSVDevelopmentLog-2   	  200000	      6124 ns/op	     197 B/op	       4 allocs/op
BenchmarkZapLTSVDevelopmentLog-2   	  200000	      6190 ns/op	     197 B/op	       4 allocs/op
BenchmarkZapLTSVDevelopmentLog-2   	  200000	      6144 ns/op	     197 B/op	       4 allocs/op
BenchmarkZapLTSVDevelopmentLog-2   	  200000	      6105 ns/op	     197 B/op	       4 allocs/op
BenchmarkZapLTSVDevelopmentLog-2   	  200000	      6102 ns/op	     197 B/op	       4 allocs/op
PASS
ok  	github.com/hnakamur/ltsvlog	67.083s
```

## License
MIT
