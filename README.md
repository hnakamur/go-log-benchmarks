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
BenchmarkLTSVLog-2                       	 1000000	      1646 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1631 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1633 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1635 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1639 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1637 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1624 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1613 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1624 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1623 ns/op	       0 B/op	       0 allocs/op
BenchmarkStandardLog-2                   	  500000	      2455 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2458 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2405 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2396 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2395 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2429 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2418 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2480 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2471 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2410 ns/op	      96 B/op	       3 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       321 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       310 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       327 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       347 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       364 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       363 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       343 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       286 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       327 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       325 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6632 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6523 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6538 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6643 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6599 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6639 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6593 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6682 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6608 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6549 ns/op	     208 B/op	       3 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      1967 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      1951 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      1965 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      1954 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      1957 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      1964 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      1975 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      1970 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      1953 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      1955 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1765 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1769 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1761 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1761 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1769 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1762 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1765 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1764 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1762 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1763 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1805 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1799 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1804 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1800 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1798 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1796 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1799 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1802 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1803 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1804 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1800 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1800 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1798 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1814 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1780 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1804 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1799 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1800 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1814 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1802 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     22114 ns/op	   16538 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     21742 ns/op	   16536 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     21895 ns/op	   16537 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     22521 ns/op	   16537 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     22543 ns/op	   16538 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	   50000	     22368 ns/op	   16539 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	   50000	     22718 ns/op	   16540 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	   50000	     24000 ns/op	   16541 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     22242 ns/op	   16537 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     21004 ns/op	   16545 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	   50000	     21938 ns/op	   16545 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     22492 ns/op	   16546 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     21750 ns/op	   16543 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     22838 ns/op	   16545 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     21980 ns/op	   16544 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	   50000	     22565 ns/op	   16544 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     22300 ns/op	   16544 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     22087 ns/op	   16544 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     22211 ns/op	   16546 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     22904 ns/op	   16547 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7750 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7738 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7827 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7790 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7751 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7834 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7554 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7672 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7807 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7807 ns/op	   16464 B/op	       4 allocs/op
PASS
ok  	github.com/hnakamur/go-log-benchmarks	198.427s
```

## License
MIT
