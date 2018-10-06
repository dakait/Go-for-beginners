## Questions

In this section I will answer some of the problems faced while coding in Go

1. How to get interface field values from an interface without using  struct?

2. Does goroutines always execute in “last in first out” manner? (I have tested it on go playground with many examples)
**Ans:** The otuput is pseudo-random. The Go playground, by design, is deterministic, which makes it easier to cache programs. So, you'll find the same order of excution everytime. If you running locally and still getting the same results, then you should check the vaule of runtime.GOMAXPROCS. If it's set to 1 as in goplayground, it will run one thing at a time sequentially. For more extended answer [link](https://test.com)   
