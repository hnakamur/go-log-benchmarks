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
BenchmarkLTSVLog-2                  	 1000000	      1609 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      1627 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      1637 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      1618 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      1626 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      1639 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      1617 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      1622 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      1612 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                  	 1000000	      1607 ns/op	       0 B/op	       0 allocs/op
BenchmarkStandardLog-2              	  500000	      2406 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2409 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2378 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	 1000000	      2412 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	 1000000	      2360 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2388 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2336 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2442 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	  500000	      2401 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2              	 1000000	      2415 ns/op	      96 B/op	       3 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       327 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       330 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       322 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       287 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       318 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       318 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       300 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       321 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       328 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2     	 5000000	       292 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6393 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6485 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6465 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6597 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6444 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6473 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6489 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6480 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6458 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2    	  200000	      6413 ns/op	     208 B/op	       3 allocs/op
BenchmarkZerologTimestamp-2         	 1000000	      1803 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	 1000000	      1795 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	 1000000	      1793 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	 1000000	      1801 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	 1000000	      1802 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	 1000000	      1787 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	 1000000	      1811 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	 1000000	      1791 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	 1000000	      1797 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2         	 1000000	      1796 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	 1000000	      1799 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	 1000000	      1804 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	 1000000	      1800 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	 1000000	      1799 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	 1000000	      1803 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	 1000000	      1805 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	 1000000	      1800 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	 1000000	      1800 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	 1000000	      1800 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2       	 1000000	      1803 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	 1000000	      1782 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	 1000000	      1779 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	 1000000	      1787 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	 1000000	      1787 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	 1000000	      1779 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	 1000000	      1784 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	 1000000	      1779 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	 1000000	      1789 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	 1000000	      1803 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2   	 1000000	      1808 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2  	  100000	     20634 ns/op	   16535 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2  	   50000	     21301 ns/op	   16536 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2  	  100000	     22078 ns/op	   16537 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2  	   50000	     23392 ns/op	   16540 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2  	  100000	     22466 ns/op	   16539 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2  	  100000	     22012 ns/op	   16538 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2  	   50000	     22101 ns/op	   16537 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2  	   50000	     21793 ns/op	   16537 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2  	  100000	     22282 ns/op	   16538 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2  	  100000	     22329 ns/op	   16538 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2         	  100000	     22365 ns/op	   16538 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2         	  100000	     21937 ns/op	   16538 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2         	  100000	     22796 ns/op	   16538 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2         	  100000	     22016 ns/op	   16546 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2         	  100000	     22274 ns/op	   16543 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2         	  100000	     21860 ns/op	   16544 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2         	  100000	     21762 ns/op	   16543 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2         	  100000	     22574 ns/op	   16546 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2         	  100000	     21754 ns/op	   16543 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2         	  100000	     22732 ns/op	   16546 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Time-2          	  200000	      7598 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2          	  200000	      7884 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2          	  200000	      7658 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2          	  200000	      7689 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2          	  200000	      7620 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2          	  200000	      7580 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2          	  200000	      7783 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2          	  200000	      7753 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2          	  200000	      7882 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2          	  200000	      7748 ns/op	   16464 B/op	       4 allocs/op
PASS
ok  	github.com/hnakamur/go-log-benchmarks	181.503s
```

## License
MIT
