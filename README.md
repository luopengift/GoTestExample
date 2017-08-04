# GoTestExample
some test example for golang

## NOTE:测试用例仅供参考!!!

1. 创建1KB的Buffer和创建[]byte，性能几乎一样
```
[root@localhost GoTestExample]# go test -bench="."  buf_test.go buf.go -v  -test.benchmem
Benchmark_Buffer 	 1000000	      1814 ns/op	    8192 B/op	       1 allocs/op
Benchmark_Byte   	 1000000	      1841 ns/op	    8192 B/op	       1 allocs/op
PASS
ok  	command-line-arguments	3.700s
```

2. mutex的性能比chan高大约4倍
```
[root@localhostGoTestExample]# go test -bench="."  chan_test.go -v  -test.benchmem
Benchmark_Mutex 	100000000	        22.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_Chan  	20000000	        81.3 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	command-line-arguments	3.941s
```
