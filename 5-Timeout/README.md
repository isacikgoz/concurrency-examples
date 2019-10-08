# 5-Timeout

In this example, `boring` function takes a string `msg` parameter and returns receive-only channel of strings again. It launches goroutine inside the function as we have seen.

In `main`, a channel is generated with `boring` function. And then a time out value is declared, it is defined as 3 seconds. After that, in `for` loop; whenever `c` channel receives a data, it is printed until the time out `to` occurs.