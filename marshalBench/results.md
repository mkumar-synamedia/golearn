goos: darwin  
goarch: amd64  

Method|Iterations|Time
---|---|---
BenchmarkHandMarshal-8|1000000000|0.302 ns/op  
BenchmarkBinaryLibMarshal-8|1240303|973 ns/op  
BenchmarkGobLibMarshal-8|336864|3035 ns/op  
PASS

Process finished with exit code 0  
