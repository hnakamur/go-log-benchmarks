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
BenchmarkLTSVLog-2                       	 1000000	      1630 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1626 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1671 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1642 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1646 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1641 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1630 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1631 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1637 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      1647 ns/op	       0 B/op	       0 allocs/op
BenchmarkStandardLog-2                   	  500000	      2379 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2446 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2427 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	 1000000	      2425 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2368 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2397 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2424 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2390 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2368 ns/op	      96 B/op	       3 allocs/op
BenchmarkStandardLog-2                   	  500000	      2458 ns/op	      96 B/op	       3 allocs/op
BenchmarkZapJSONProductionLog-2          	 5000000	       345 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 5000000	       302 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 5000000	       345 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 5000000	       340 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 5000000	       350 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 5000000	       359 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 5000000	       364 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 5000000	       326 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 5000000	       327 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 5000000	       326 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      6875 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      6765 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      6779 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      6792 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      6814 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      6746 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      6872 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      6768 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      6804 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      6864 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       382 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       354 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       345 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       363 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       304 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       362 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       303 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       335 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       302 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       325 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6599 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6548 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6519 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6582 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6555 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6653 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6598 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6559 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6647 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      6553 ns/op	     208 B/op	       3 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2013 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2006 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      2003 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      2024 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      2009 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      2002 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      2004 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      2003 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      2005 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      2004 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1800 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1779 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1770 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1783 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1778 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1772 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1788 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1771 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1795 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1777 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1783 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1798 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1779 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1781 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1771 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1773 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1770 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1779 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1803 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1780 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1784 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1795 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1786 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1790 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1787 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1794 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1780 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1785 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1779 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1782 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     22684 ns/op	   16544 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     24013 ns/op	   16546 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     22391 ns/op	   16544 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     23007 ns/op	   16544 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     22835 ns/op	   16545 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     23087 ns/op	   16545 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     22553 ns/op	   16544 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     22921 ns/op	   16544 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     22035 ns/op	   16543 B/op	       5 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	   50000	     21824 ns/op	   16544 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     21760 ns/op	   16544 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     22178 ns/op	   16544 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     22265 ns/op	   16543 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     22287 ns/op	   16544 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     22031 ns/op	   16543 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     21730 ns/op	   16544 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     21314 ns/op	   16543 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     22268 ns/op	   16543 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     22501 ns/op	   16544 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     21601 ns/op	   16542 B/op	       5 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7856 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7893 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      8090 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7944 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      8010 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7994 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      8028 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7985 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      8077 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      7980 ns/op	   16464 B/op	       4 allocs/op
PASS
ok  	github.com/hnakamur/go-log-benchmarks	238.323s
```

## License
MIT
