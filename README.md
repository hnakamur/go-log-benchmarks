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
BenchmarkLTSVLog-2                       	 1000000	      1702 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1701 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1692 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1697 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1690 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1703 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1688 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1697 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1688 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1691 ns/op	       0 B/op	       0 allocs/op
BenchmarkStandardLog-2                   	  500000	      2404 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2438 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2369 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2368 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2444 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2407 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2408 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2434 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2398 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2388 ns/op	      96 B/op	       3 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       298 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       347 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       374 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 3000000	       374 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       348 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       336 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       291 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       345 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       338 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       341 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6494 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6499 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6611 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6545 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6522 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6548 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6628 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6554 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6554 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6574 ns/op	     208 B/op	       3 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2714 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2708 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2706 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2703 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2714 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2707 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2702 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2711 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2701 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2703 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1835 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1851 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1835 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1837 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1830 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1827 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1840 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1832 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1835 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1893 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1838 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1848 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1844 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1837 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1841 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1839 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1837 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1835 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1836 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1842 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1844 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1847 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1852 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1839 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1851 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1841 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1847 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1846 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1880 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1843 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     22046 ns/op	   16537 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     22240 ns/op	   16537 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     23200 ns/op	   16537 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     21353 ns/op	   16536 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	   50000	     22062 ns/op	   16539 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     22065 ns/op	   16537 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     21884 ns/op	   16537 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     21686 ns/op	   16538 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     21059 ns/op	   16535 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     21961 ns/op	   16544 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     21380 ns/op	   16543 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     21198 ns/op	   16544 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     21859 ns/op	   16544 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     22367 ns/op	   16543 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     21063 ns/op	   16542 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     20960 ns/op	   16543 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     22551 ns/op	   16545 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     21109 ns/op	   16543 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     20666 ns/op	   16541 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     20995 ns/op	   16542 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7670 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7440 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7824 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7682 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7760 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7669 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7665 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7544 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7723 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7673 ns/op	   16464 B/op	       4 allocs/op
PASS
ok  	github.com/hnakamur/go-log-benchmarks	196.199s
```

## License
MIT
