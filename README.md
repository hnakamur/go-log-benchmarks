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
goos: linux
goarch: amd64
pkg: github.com/hnakamur/go-log-benchmarks
BenchmarkLTSVLog-2                       	 1000000	      2111 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      2102 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      2109 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      2099 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      2102 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      2117 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      2099 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	  500000	      2116 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      2113 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      2098 ns/op	       0 B/op	       0 allocs/op
BenchmarkStandardLog-2                   	  500000	      2673 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLog-2                   	  500000	      2654 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLog-2                   	  500000	      2678 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLog-2                   	  500000	      2663 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLog-2                   	  500000	      2658 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLog-2                   	  500000	      2702 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLog-2                   	  500000	      2636 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLog-2                   	  500000	      2661 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLog-2                   	  500000	      2656 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLog-2                   	  500000	      2725 ns/op	      64 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 3000000	       409 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 5000000	       396 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 5000000	       376 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 3000000	       434 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 3000000	       437 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 3000000	       399 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 3000000	       445 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 5000000	       416 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 3000000	       429 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 3000000	       354 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      8291 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      8146 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      8086 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      8170 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      8011 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      8237 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      8117 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      8140 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      7873 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      8288 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapLTSVProductionLog-2          	 3000000	       452 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 3000000	       418 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 3000000	       441 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 3000000	       429 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       430 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       413 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 3000000	       428 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 3000000	       432 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       428 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 3000000	       386 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      8287 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      8365 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      8515 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      8531 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      8630 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      8457 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      8560 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      8484 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      8566 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      8515 ns/op	     208 B/op	       3 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2252 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2213 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      2277 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2209 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2209 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2213 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      2216 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2195 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2229 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2199 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1952 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1938 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1961 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1938 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1928 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1934 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1935 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1941 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1965 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1934 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1932 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1936 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1932 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1938 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1964 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1967 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1969 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1943 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1945 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1939 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1950 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1977 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1972 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1948 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1952 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1968 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1947 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1963 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1961 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1956 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     17612 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     17547 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     18052 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     17639 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     18003 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     18080 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     17835 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     17962 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     17703 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     17869 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     17135 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     18247 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     16888 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     17247 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     17457 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     16977 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     17410 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     17161 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     17518 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     17600 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      9117 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      9177 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      9251 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      9178 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      9133 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	     12168 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      9183 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      9176 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      9059 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      9098 ns/op	   16464 B/op	       4 allocs/op
PASS
ok  	github.com/hnakamur/go-log-benchmarks	241.961s
```

## License
MIT
