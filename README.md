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
BenchmarkLTSVLog-2                       	 1000000	      1663 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1655 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1677 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1670 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1659 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1667 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1665 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1665 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1670 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1671 ns/op	       0 B/op	       0 allocs/op
BenchmarkStandardLog-2                   	  500000	      2473 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2446 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2504 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2430 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2483 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2441 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2465 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2501 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2447 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2530 ns/op	      96 B/op	       3 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       304 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       325 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       304 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       353 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       299 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       288 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       348 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       354 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       314 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 3000000	       357 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6953 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6739 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6739 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6793 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6720 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6746 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6761 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6722 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6809 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6757 ns/op	     208 B/op	       3 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2594 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2595 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2592 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2599 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2602 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2598 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2596 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2603 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2610 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2607 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1772 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1779 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1771 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1775 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1770 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1776 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1775 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1774 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1771 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1780 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1805 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1823 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1795 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1794 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1804 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1799 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1791 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1804 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1796 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1802 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1793 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1788 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1804 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1803 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1797 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1789 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1787 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1791 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1809 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1807 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	   50000	     22333 ns/op	   16537 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	   50000	     22195 ns/op	   16536 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     22388 ns/op	   16536 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     22326 ns/op	   16538 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     23240 ns/op	   16539 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	   50000	     23871 ns/op	   16542 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     22200 ns/op	   16537 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     22927 ns/op	   16539 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	   50000	     22316 ns/op	   16537 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     22373 ns/op	   16537 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     22019 ns/op	   16537 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     21567 ns/op	   16543 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     23493 ns/op	   16546 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	   50000	     23494 ns/op	   16547 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	   50000	     23711 ns/op	   16547 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	   50000	     23137 ns/op	   16545 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	   50000	     23913 ns/op	   16546 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	   50000	     22328 ns/op	   16544 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     23324 ns/op	   16546 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	   50000	     22645 ns/op	   16545 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7886 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7890 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7657 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7804 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7810 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7792 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7808 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7954 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7957 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7995 ns/op	   16464 B/op	       4 allocs/op
PASS
ok  	github.com/hnakamur/go-log-benchmarks	186.276s
```

## License
MIT
