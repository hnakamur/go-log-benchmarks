go-log-benchmarks
=================

A go logging library benchmarks for me.

## Benchmark result

Note these benchmarks print different outputs.
Its purpose is not to choose the most performant library, but to get rough cost
of each operation in each library.

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
BenchmarkLTSVLog-2                  	 1000000	      2075 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      2069 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                  	  500000	      2083 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      2089 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                  	  500000	      2084 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      2052 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      2066 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      2036 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      2041 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      2072 ns/op	      48 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2332 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	 1000000	      2343 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2343 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2362 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	 1000000	      2350 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2335 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	 1000000	      2358 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2344 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2390 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2346 ns/op	      96 B/op	       3 allocs/op
BenchmarkZapLTSVProductionLog-2     	 3000000	       371 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       362 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       377 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       336 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       344 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       344 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       341 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 3000000	       380 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       381 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       352 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6431 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6340 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6487 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6520 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6422 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6447 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6382 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6441 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6409 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6540 ns/op	     208 B/op	       3 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2611 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2608 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2605 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2602 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2608 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2604 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2614 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2615 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2604 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2613 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2657 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2654 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2661 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2656 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2653 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2648 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2649 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2653 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2659 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2658 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2621 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2625 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2624 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2625 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2630 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2623 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2627 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2626 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2624 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2627 ns/op	       0 B/op	       0 allocs/op
BenchmarkErrWithStackAndTime-2      	   50000	     21598 ns/op	   10017 B/op	      11 allocs/op
BenchmarkErrWithStackAndTime-2      	  100000	     20383 ns/op	   10017 B/op	      11 allocs/op
BenchmarkErrWithStackAndTime-2      	  100000	     24455 ns/op	   10017 B/op	      11 allocs/op
BenchmarkErrWithStackAndTime-2      	  100000	     23510 ns/op	   10017 B/op	      11 allocs/op
BenchmarkErrWithStackAndTime-2      	   50000	     21495 ns/op	   10017 B/op	      11 allocs/op
BenchmarkErrWithStackAndTime-2      	  100000	     21513 ns/op	   10017 B/op	      11 allocs/op
BenchmarkErrWithStackAndTime-2      	  100000	     22671 ns/op	   10017 B/op	      11 allocs/op
BenchmarkErrWithStackAndTime-2      	  100000	     22986 ns/op	   10017 B/op	      11 allocs/op
BenchmarkErrWithStackAndTime-2      	  100000	     22877 ns/op	   10017 B/op	      11 allocs/op
BenchmarkErrWithStackAndTime-2      	  100000	     22511 ns/op	   10017 B/op	      11 allocs/op
BenchmarkErrWithStack-2             	   50000	     21795 ns/op	    9841 B/op	       9 allocs/op
BenchmarkErrWithStack-2             	  100000	     22860 ns/op	    9841 B/op	       9 allocs/op
BenchmarkErrWithStack-2             	  100000	     21572 ns/op	    9841 B/op	       9 allocs/op
BenchmarkErrWithStack-2             	   50000	     22533 ns/op	    9841 B/op	       9 allocs/op
BenchmarkErrWithStack-2             	  100000	     22329 ns/op	    9847 B/op	       9 allocs/op
BenchmarkErrWithStack-2             	  100000	     21058 ns/op	    9841 B/op	       9 allocs/op
BenchmarkErrWithStack-2             	  100000	     22096 ns/op	    9842 B/op	       9 allocs/op
BenchmarkErrWithStack-2             	  100000	     22186 ns/op	    9842 B/op	       9 allocs/op
BenchmarkErrWithStack-2             	  100000	     22059 ns/op	    9842 B/op	       9 allocs/op
BenchmarkErrWithStack-2             	  100000	     21170 ns/op	    9841 B/op	       9 allocs/op
BenchmarkErrWithTime-2              	  300000	      3900 ns/op	     400 B/op	       7 allocs/op
BenchmarkErrWithTime-2              	  300000	      3832 ns/op	     400 B/op	       7 allocs/op
BenchmarkErrWithTime-2              	  300000	      3794 ns/op	     400 B/op	       7 allocs/op
BenchmarkErrWithTime-2              	  300000	      3770 ns/op	     400 B/op	       7 allocs/op
BenchmarkErrWithTime-2              	  300000	      3812 ns/op	     400 B/op	       7 allocs/op
BenchmarkErrWithTime-2              	  300000	      3814 ns/op	     400 B/op	       7 allocs/op
BenchmarkErrWithTime-2              	  300000	      3848 ns/op	     400 B/op	       7 allocs/op
BenchmarkErrWithTime-2              	  500000	      3836 ns/op	     400 B/op	       7 allocs/op
BenchmarkErrWithTime-2              	  300000	      3807 ns/op	     400 B/op	       7 allocs/op
BenchmarkErrWithTime-2              	  300000	      3798 ns/op	     400 B/op	       7 allocs/op
BenchmarkErr-2                      	  500000	      3173 ns/op	     352 B/op	       5 allocs/op
BenchmarkErr-2                      	  500000	      3198 ns/op	     352 B/op	       5 allocs/op
BenchmarkErr-2                      	  500000	      3192 ns/op	     352 B/op	       5 allocs/op
BenchmarkErr-2                      	  500000	      3171 ns/op	     352 B/op	       5 allocs/op
BenchmarkErr-2                      	  500000	      3212 ns/op	     352 B/op	       5 allocs/op
BenchmarkErr-2                      	  500000	      3226 ns/op	     352 B/op	       5 allocs/op
BenchmarkErr-2                      	  500000	      3189 ns/op	     352 B/op	       5 allocs/op
BenchmarkErr-2                      	  500000	      3117 ns/op	     352 B/op	       5 allocs/op
BenchmarkErr-2                      	  500000	      3038 ns/op	     352 B/op	       5 allocs/op
BenchmarkErr-2                      	  500000	      3009 ns/op	     352 B/op	       5 allocs/op
PASS
ok  	github.com/hnakamur/go-log-benchmarks	183.393s
```

## License
MIT
