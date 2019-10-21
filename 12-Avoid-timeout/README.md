# 12-Avoid-timeout

In the previous example, sometimes timeout happens and we cannot get the proper result. To avoid timeout, in this example we want to get results from multiple backends. We do so by initializing `result` to a different function in `main`.

```
result := first("golang",
    fakeSearch("replica 1"),
    fakeSearch("replica 2"),
)
```

`first` function takes a `query` string and a `variadic function` (replicas ...Search):

- Here, that `variadic function` takes an arbitrary number of `Search` functions as arguments. 
- A channel of `Result` is generated.
- `searchReplica` variable is initialized as a function.
    - Inside the function, channel receives data from replicas.
- In `for` loop, a search with a seperate goroutine for each replica is started.
- Afterwards, the data in the channel is returned.

So, whichever replica completes the search quickest, its result will return.