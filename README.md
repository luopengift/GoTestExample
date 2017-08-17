# GoTestExample
some test example for golang

## NOTE:测试用例仅供参考!!!

1. 创建1KB的Buffer和创建[]byte，性能几乎一样
```
[root@localhost GoTestExample]# go test -v -test.benchmem -bench="."  buf_test.go
Benchmark_Buffer 	 1000000	      1805 ns/op	    8192 B/op	       1 allocs/op
Benchmark_Byte   	 1000000	      1808 ns/op	    8192 B/op	       1 allocs/op
PASS
ok  	command-line-arguments	3.659s
```

2. mutex的性能比chan高大约4倍
```
[root@iZm5egf7xb48axmu4z1t3fZ GoTestExample]# go test -v -test.benchmem -bench="."  chan_test.go
Benchmark_Mutex 	100000000	        21.6 ns/op	       0 B/op	       0 allocs/op
Benchmark_Chan  	20000000	        81.4 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	command-line-arguments	3.903s
```
3. "xx = xx + 1" 会比 "xx += 1" 性能高一点点
```
[root@iZm5egf7xb48axmu4z1t3fZ GoTestExample]# go test -v -test.benchmem -bench="."  plus_test.go
Benchmark_plusEqual     2000000000           0.49 ns/op        0 B/op          0 allocs/op
Benchmark_plusOne       2000000000           0.41 ns/op        0 B/op          0 allocs/op
PASS
ok      command-line-arguments  1.890s
```


