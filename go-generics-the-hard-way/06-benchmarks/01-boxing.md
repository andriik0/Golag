# Boxing

This benchmark demonstrates how Go generics speedup scenarios where boxing can be avoided.

* [**What is boxing?**](#what-is-boxing?): a summary of boxing
* [**The example**](#the-example): an exemplar case
* [**The benchmark**](#the-benchmark): how does it stack up?
* [**Key takeaways**](#key-takeaways): what it all means

## What is boxing?

In general, the term _boxing_ refers to wrapping a value with some container, aka, "putting it inside a box." For the purposes of _this_ page, the term "boxing" will refer to wrapping some value in Golang with the empty interface, ex.:

```go
x := int(0)
i := interface{}(x)
```

Generics can be used to eliminate the need for boxing in many situations, which we will highlight with `List[T]`.


## The example

This page describes how to run `BenchmarkBoxing`, a Go-based benchmark that benchmarks lists from the following packages:

* **`./lists/boxed`**: defines `type List []interface{}`
* **`./lists/typed`**: defines `type IntList []int`
* **`./lists/generic`**: defines `type List[T any] []T`


## The benchmark

In order to determine the impact of boxing, we will run a Go benchmark that adds random integers to instances of a `boxed.List`, `typed.IntList`, and `generic.List[int]`:

---

:warning: **Please note** that Docker may hang if not provided enough resources. The performance improvements are significant enough that the container may be starved for CPU and killed.

---

```bash
docker run -it --rm go-generics-the-hard-way \
  go test -bench Boxing -run Boxing -benchmem -count 5 -v ./06-benchmarks
```

The following table was generated by piping the above command to `hack/b2md.py -t boxing` and represents an average across the total number of benchmarks determined by the `-count` flag:

| List type | Number of types | Operations | ns/op | Bytes/op | Allocs/op |
|:---------:|:---------------:|:----------:|:-----:|:--------:|:---------:|
| Boxed | 1 | 28639768.6 | 49.42 | 100.4 | 0 |
| Generic | 1 | 217233399 | 8.31 | 45.4 | 0 |
| Typed | 1 | 300006311.8 | 8.49 | 43.4 | 0 |


## Key takeaways

A few, key takeaways:

* On average the implementation of `List[T any]` was more performant than the boxed list:
  * operations were 10x faster
  * consumed half the memory
* The performance improvements were the result of removing the need to box the integer values

---

Next: [Build times](./02-build-times.md)