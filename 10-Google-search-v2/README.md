# 10-Google-search-v2

This example is very similar to the previous one except for `Google` function. Here, that function has a difference. In `Google`:

- A channel of `Result` is created.
- 3 functions are defined and start with 3 different goroutines.
    - In each goroutine, the channel receives data from `web`, `image` and `video`.
- In `for` loop, whenever the channel receives a data, it is appended to the `results`.

By using goroutines, the `elapsed` time is shortened. In the previous example, `elapsed time` was total time of each opeartion's time. As it is seen, which of the operations `web`, `image`, `video` has the longest `elapsed` time, its duration is printed in this example. 