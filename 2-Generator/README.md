# 2-Generator

In this example's `boring` function is differentiate from first example's. This time, the function returns receive-only channel of strings. And it is called as `Generator Pattern` which is a function that returns a channel of something. Also it should or can return a read only channel, and by the end of the generation the channel can be closed by the generator function itself. So, in this example, the channel is created in the `boring` function and the goroutine is launched in that function. (The same `for` loop as in previous example's `boring` function.) 

In `main`, firstly a function is generated named as `c`. And the same `for` loop as in previous example's `main` is performed.