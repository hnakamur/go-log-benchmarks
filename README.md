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
BenchmarkLTSVLog-2                  	 1000000	      2103 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      2100 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      2118 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      2104 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      2123 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                  	  500000	      2096 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      2092 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      2109 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                  	  500000	      2120 ns/op	      48 B/op	       3 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      2099 ns/op	      48 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2420 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	 1000000	      2414 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2398 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2391 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2361 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2361 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2363 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2378 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	 1000000	      2389 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2430 ns/op	      96 B/op	       3 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       279 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       364 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       351 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       353 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       390 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       376 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       360 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       324 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       319 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       359 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6497 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6389 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6424 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6384 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6419 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6449 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6445 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6487 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6411 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6526 ns/op	     208 B/op	       3 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2663 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2612 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2647 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2616 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2657 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2621 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2653 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2616 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2663 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	  500000	      2618 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2711 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2663 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2712 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2668 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2702 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2684 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2716 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2667 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2707 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	  500000	      2659 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2670 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2624 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2673 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2624 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2674 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2627 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2672 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2628 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2675 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	  500000	      2624 ns/op	       0 B/op	       0 allocs/op
BenchmarkErrLVWithStackAndTime-2    	   50000	     22394 ns/op	   10017 B/op	      11 allocs/op
BenchmarkErrLVWithStackAndTime-2    	  100000	     21828 ns/op	   10017 B/op	      11 allocs/op
BenchmarkErrLVWithStackAndTime-2    	   50000	     23231 ns/op	   10017 B/op	      11 allocs/op
BenchmarkErrLVWithStackAndTime-2    	  100000	     23266 ns/op	   10017 B/op	      11 allocs/op
BenchmarkErrLVWithStackAndTime-2    	  100000	     22758 ns/op	   10017 B/op	      11 allocs/op
BenchmarkErrLVWithStackAndTime-2    	  100000	     22447 ns/op	   10017 B/op	      11 allocs/op
BenchmarkErrLVWithStackAndTime-2    	   50000	     23808 ns/op	   10017 B/op	      11 allocs/op
BenchmarkErrLVWithStackAndTime-2    	  100000	     21605 ns/op	   10017 B/op	      11 allocs/op
BenchmarkErrLVWithStackAndTime-2    	  100000	     22623 ns/op	   10017 B/op	      11 allocs/op
BenchmarkErrLVWithStackAndTime-2    	  100000	     23431 ns/op	   10017 B/op	      11 allocs/op
BenchmarkErrLVWithStack-2           	   50000	     20904 ns/op	    9969 B/op	       9 allocs/op
BenchmarkErrLVWithStack-2           	  100000	     23502 ns/op	    9969 B/op	       9 allocs/op
BenchmarkErrLVWithStack-2           	   50000	     23616 ns/op	    9969 B/op	       9 allocs/op
BenchmarkErrLVWithStack-2           	  100000	     22671 ns/op	    9975 B/op	       9 allocs/op
BenchmarkErrLVWithStack-2           	   50000	     22218 ns/op	    9970 B/op	       9 allocs/op
BenchmarkErrLVWithStack-2           	   50000	     23121 ns/op	    9970 B/op	       9 allocs/op
BenchmarkErrLVWithStack-2           	   50000	     22489 ns/op	    9970 B/op	       9 allocs/op
BenchmarkErrLVWithStack-2           	   50000	     22861 ns/op	    9970 B/op	       9 allocs/op
BenchmarkErrLVWithStack-2           	   50000	     23061 ns/op	    9970 B/op	       9 allocs/op
BenchmarkErrLVWithStack-2           	  100000	     22668 ns/op	    9970 B/op	       9 allocs/op
BenchmarkErrLVWithTime-2            	  300000	      3854 ns/op	     400 B/op	       7 allocs/op
BenchmarkErrLVWithTime-2            	  500000	      3915 ns/op	     400 B/op	       7 allocs/op
BenchmarkErrLVWithTime-2            	  300000	      3942 ns/op	     400 B/op	       7 allocs/op
BenchmarkErrLVWithTime-2            	  300000	      3879 ns/op	     400 B/op	       7 allocs/op
BenchmarkErrLVWithTime-2            	  500000	      3899 ns/op	     400 B/op	       7 allocs/op
BenchmarkErrLVWithTime-2            	  500000	      3943 ns/op	     400 B/op	       7 allocs/op
BenchmarkErrLVWithTime-2            	  300000	      3810 ns/op	     400 B/op	       7 allocs/op
BenchmarkErrLVWithTime-2            	  300000	      3924 ns/op	     400 B/op	       7 allocs/op
BenchmarkErrLVWithTime-2            	  500000	      3854 ns/op	     400 B/op	       7 allocs/op
BenchmarkErrLVWithTime-2            	  300000	      3865 ns/op	     400 B/op	       7 allocs/op
PASS
ok  	github.com/hnakamur/go-log-benchmarks	163.507s
```

## License
MIT
