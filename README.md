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
BenchmarkLTSVLog-2                  	 1000000	      1597 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      1591 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      1595 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      1599 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      1596 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      1595 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      1598 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      1595 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      1595 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      1605 ns/op	       0 B/op	       0 allocs/op
BenchmarkStandardLog-2              	  500000	      2406 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2370 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2427 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2391 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2416 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2433 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2395 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2457 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	 1000000	      2383 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2424 ns/op	      96 B/op	       3 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       314 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       393 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       380 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       354 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       354 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       311 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       332 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       358 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 3000000	       367 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       325 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6586 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6587 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6596 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6520 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6543 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6550 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6450 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6515 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6528 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6499 ns/op	     208 B/op	       3 allocs/op
BenchmarkZerologTimestamp-2         	 1000000	      1758 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	 1000000	      1761 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	 1000000	      1760 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	 1000000	      1760 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	 1000000	      2045 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	 1000000	      1761 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	 1000000	      1760 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	 1000000	      1762 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	 1000000	      1757 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	 1000000	      1758 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	 1000000	      1775 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	 1000000	      1776 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	 1000000	      1776 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	 1000000	      1779 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	 1000000	      1777 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	 1000000	      1781 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	 1000000	      1780 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	 1000000	      1775 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	 1000000	      1778 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	 1000000	      1778 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	 1000000	      1746 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	 1000000	      1752 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	 1000000	      1752 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	 1000000	      1746 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	 1000000	      1753 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	 1000000	      1756 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	 1000000	      1754 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	 1000000	      1756 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	 1000000	      1751 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	 1000000	      1753 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2  	  100000	     18752 ns/op	    8262 B/op	       3 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2  	  100000	     20524 ns/op	    8263 B/op	       3 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2  	  100000	     21456 ns/op	    8264 B/op	       3 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2  	  100000	     20574 ns/op	    8263 B/op	       3 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2  	  100000	     18602 ns/op	    8261 B/op	       3 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2  	  100000	     20631 ns/op	    8263 B/op	       3 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2  	  100000	     19430 ns/op	    8262 B/op	       3 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2  	  100000	     18966 ns/op	    8262 B/op	       3 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2  	  100000	     20009 ns/op	    8263 B/op	       3 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2  	  100000	     20551 ns/op	    8263 B/op	       3 allocs/op
BenchmarkLTSVLogErr/Stack-2         	  100000	     20009 ns/op	    8264 B/op	       3 allocs/op
BenchmarkLTSVLogErr/Stack-2         	  100000	     21155 ns/op	    8263 B/op	       3 allocs/op
BenchmarkLTSVLogErr/Stack-2         	  100000	     19434 ns/op	    8262 B/op	       3 allocs/op
BenchmarkLTSVLogErr/Stack-2         	  100000	     18469 ns/op	    8269 B/op	       3 allocs/op
BenchmarkLTSVLogErr/Stack-2         	  100000	     19320 ns/op	    8266 B/op	       3 allocs/op
BenchmarkLTSVLogErr/Stack-2         	  100000	     18213 ns/op	    8264 B/op	       3 allocs/op
BenchmarkLTSVLogErr/Stack-2         	  100000	     17773 ns/op	    8264 B/op	       3 allocs/op
BenchmarkLTSVLogErr/Stack-2         	  100000	     18781 ns/op	    8264 B/op	       3 allocs/op
BenchmarkLTSVLogErr/Stack-2         	  100000	     18276 ns/op	    8264 B/op	       3 allocs/op
BenchmarkLTSVLogErr/Stack-2         	  100000	     19311 ns/op	    8266 B/op	       3 allocs/op
BenchmarkLTSVLogErr/Time-2          	  200000	      5871 ns/op	    8209 B/op	       2 allocs/op
BenchmarkLTSVLogErr/Time-2          	  200000	      5924 ns/op	    8209 B/op	       2 allocs/op
BenchmarkLTSVLogErr/Time-2          	  200000	      5794 ns/op	    8209 B/op	       2 allocs/op
BenchmarkLTSVLogErr/Time-2          	  300000	      5775 ns/op	    8209 B/op	       2 allocs/op
BenchmarkLTSVLogErr/Time-2          	  300000	      5855 ns/op	    8210 B/op	       2 allocs/op
BenchmarkLTSVLogErr/Time-2          	  200000	      5661 ns/op	    8209 B/op	       2 allocs/op
BenchmarkLTSVLogErr/Time-2          	  300000	      5777 ns/op	    8209 B/op	       2 allocs/op
BenchmarkLTSVLogErr/Time-2          	  300000	      5864 ns/op	    8209 B/op	       2 allocs/op
BenchmarkLTSVLogErr/Time-2          	  200000	      5746 ns/op	    8209 B/op	       2 allocs/op
BenchmarkLTSVLogErr/Time-2          	  300000	      5863 ns/op	    8209 B/op	       2 allocs/op
PASS
ok  	github.com/hnakamur/go-log-benchmarks	178.411s
```

## License
MIT
