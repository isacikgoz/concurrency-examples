# 7-Receive-on-quit

In this example, the program says it's done and then finishes.

In `boring`, everything is the same as the previous example except for `quit`. This time, `quit` is a channel of strings not booleans. And in `quit` case of `select`, before it returns, there is a function call (`cleanup`) and after that `quit` receives a string message `see ya!`.

There is also a `cleanup` function as mentioned above. In this function, an `all clean` message is printed. 

In `main`, differently from the previous example, `bye` string is sent to `quit` channel. And the leaving message contains a data from `quit` channel. But after execution of that example, we cannot see the `bye`. `quit` channel receives `see ya!` string in `boring`.