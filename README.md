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
$ go test -count=10 -bench . -benchmem -cpuprofile=cpu.prof | tee 20170922.log
goos: linux
goarch: amd64
pkg: github.com/hnakamur/go-log-benchmarks
BenchmarkLTSVLog-2                       	  500000	      2102 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	  500000	      2097 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      2075 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      2072 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      2093 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	  500000	      2088 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      2084 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      2066 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      2075 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-2                       	 1000000	      2093 ns/op	       0 B/op	       0 allocs/op
BenchmarkStandardLog-2                   	  500000	      2666 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLog-2                   	  500000	      2627 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLog-2                   	  500000	      2686 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLog-2                   	  500000	      2655 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLog-2                   	  500000	      2666 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLog-2                   	  500000	      2652 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLog-2                   	  500000	      2691 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLog-2                   	  500000	      2658 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLog-2                   	  500000	      2627 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLog-2                   	  500000	      2682 ns/op	      64 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 5000000	       388 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 5000000	       388 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 3000000	       433 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 3000000	       353 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 5000000	       402 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 3000000	       410 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 5000000	       399 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 5000000	       375 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 5000000	       378 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-2          	 5000000	       401 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      7969 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      7763 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      7868 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      7975 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      7832 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      7935 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      7780 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      7874 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      7921 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapJSONDevelopmentLog-2         	  200000	      8170 ns/op	     288 B/op	       7 allocs/op
BenchmarkZapLTSVProductionLog-2          	 3000000	       394 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       415 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 3000000	       341 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       382 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 3000000	       389 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       411 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 3000000	       388 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       413 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 3000000	       374 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-2          	 5000000	       404 ns/op	     128 B/op	       1 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      8307 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      8211 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      8254 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      8499 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      8368 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      8600 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      8445 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      8451 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      8347 ns/op	     208 B/op	       3 allocs/op
BenchmarkZapLTSVDevelopmentLog-2         	  200000	      8246 ns/op	     208 B/op	       3 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      2194 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      2198 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      2185 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      2191 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      2190 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      2194 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      2173 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      2170 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	 1000000	      2202 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-2              	  500000	      2194 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1951 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1935 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1955 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1935 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1937 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1937 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1946 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1932 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1949 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-2   	 1000000	      1935 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1981 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1966 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1976 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1967 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1960 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1972 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1975 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1970 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1997 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-2            	 1000000	      1958 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1983 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1987 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1985 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1971 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1959 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1966 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1976 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1971 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1980 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-2        	 1000000	      1981 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     18031 ns/op	   16467 B/op	       4 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     17755 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     17929 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     18079 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     17937 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     17908 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     18074 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     17798 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     18041 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/StackAndTime-2       	  100000	     17653 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     17629 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     17459 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     16972 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     17642 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     17529 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     17096 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     16916 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     17097 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     17453 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Stack-2              	  100000	     17272 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      9168 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      9141 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      9166 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      9112 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      9219 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      9122 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      9079 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      9036 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      9160 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-2               	  200000	      9145 ns/op	   16464 B/op	       4 allocs/op
PASS
ok  	github.com/hnakamur/go-log-benchmarks	248.394s
```

```
$ benchstat new.log 20170922.log | tee benchstat-20170922.log
name                           old time/op    new time/op    delta
LTSVLog-2                        2.11µs ± 0%    2.08µs ± 1%  -1.05%  (p=0.000 n=10+10)
StandardLog-2                    2.66µs ± 1%    2.66µs ± 1%    ~     (p=0.858 n=9+10)
ZapJSONProductionLog-2            410ns ±14%     393ns ±10%    ~     (p=0.148 n=10+10)
ZapJSONDevelopmentLog-2          8.14µs ± 3%    7.91µs ± 3%  -2.79%  (p=0.002 n=10+10)
ZapLTSVProductionLog-2            430ns ± 5%     391ns ±13%  -9.07%  (p=0.000 n=9+10)
ZapLTSVDevelopmentLog-2          8.49µs ± 2%    8.37µs ± 3%  -1.39%  (p=0.027 n=10+10)
ZerologTimestamp-2               2.21µs ± 2%    2.19µs ± 1%  -1.17%  (p=0.000 n=9+10)
ZerologTimeSecondsFromEpoch-2    1.94µs ± 1%    1.94µs ± 1%    ~     (p=0.929 n=10+10)
ZerologRFC3339Time-2             1.95µs ± 1%    1.97µs ± 1%  +1.18%  (p=0.002 n=10+9)
ZerologRFC3339NanoTime-2         1.96µs ± 1%    1.98µs ± 1%  +0.84%  (p=0.004 n=10+10)
LTSVLogErr/StackAndTime-2        17.8µs ± 2%    17.9µs ± 1%    ~     (p=0.353 n=10+10)
LTSVLogErr/Stack-2               17.3µs ± 2%    17.3µs ± 2%    ~     (p=0.720 n=9+10)
LTSVLogErr/Time-2                9.15µs ± 1%    9.13µs ± 1%    ~     (p=0.400 n=9+10)

name                           old alloc/op   new alloc/op   delta
LTSVLog-2                         0.00B          0.00B         ~     (all equal)
StandardLog-2                     64.0B ± 0%     64.0B ± 0%    ~     (all equal)
ZapJSONProductionLog-2             128B ± 0%      128B ± 0%    ~     (all equal)
ZapJSONDevelopmentLog-2            288B ± 0%      288B ± 0%    ~     (all equal)
ZapLTSVProductionLog-2             128B ± 0%      128B ± 0%    ~     (all equal)
ZapLTSVDevelopmentLog-2            208B ± 0%      208B ± 0%    ~     (all equal)
ZerologTimestamp-2                0.00B          0.00B         ~     (all equal)
ZerologTimeSecondsFromEpoch-2     0.00B          0.00B         ~     (all equal)
ZerologRFC3339Time-2              0.00B          0.00B         ~     (all equal)
ZerologRFC3339NanoTime-2          0.00B          0.00B         ~     (all equal)
LTSVLogErr/StackAndTime-2        16.5kB ± 0%    16.5kB ± 0%    ~     (all equal)
LTSVLogErr/Stack-2               16.5kB ± 0%    16.5kB ± 0%    ~     (all equal)
LTSVLogErr/Time-2                16.5kB ± 0%    16.5kB ± 0%    ~     (all equal)

name                           old allocs/op  new allocs/op  delta
LTSVLog-2                          0.00           0.00         ~     (all equal)
StandardLog-2                      1.00 ± 0%      1.00 ± 0%    ~     (all equal)
ZapJSONProductionLog-2             1.00 ± 0%      1.00 ± 0%    ~     (all equal)
ZapJSONDevelopmentLog-2            7.00 ± 0%      7.00 ± 0%    ~     (all equal)
ZapLTSVProductionLog-2             1.00 ± 0%      1.00 ± 0%    ~     (all equal)
ZapLTSVDevelopmentLog-2            3.00 ± 0%      3.00 ± 0%    ~     (all equal)
ZerologTimestamp-2                 0.00           0.00         ~     (all equal)
ZerologTimeSecondsFromEpoch-2      0.00           0.00         ~     (all equal)
ZerologRFC3339Time-2               0.00           0.00         ~     (all equal)
ZerologRFC3339NanoTime-2           0.00           0.00         ~     (all equal)
LTSVLogErr/StackAndTime-2          4.00 ± 0%      4.00 ± 0%    ~     (all equal)
LTSVLogErr/Stack-2                 4.00 ± 0%      4.00 ± 0%    ~     (all equal)
LTSVLogErr/Time-2                  4.00 ± 0%      4.00 ± 0%    ~     (all equal)
```

## License
MIT
