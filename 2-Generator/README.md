# 2-Generator

In this example's `boring` function is differentiate from first example's. This time, the function returns receive-only channel of strings. The channel is created in the `boring` function and the goroutine is launched in that function. (The same `for` loop as in previous example's `boring` function.) 

In `main`, firstly a function is generated named as `c`. And the same `for` loop as in previous example's `main` is performed.