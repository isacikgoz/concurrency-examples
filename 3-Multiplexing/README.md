# 3-Multiplexing

Here, again a `boring` function is the same as previous `generator pattern`. But also here has `fan in pattern`: a function that returns a channel.

`fanInSelect` function takes 2 channels of strings as parameters (a and b) and returns a channel of strings (c). Also it contains a single goroutine. In `fanInSelect`, a channel (c) is generated. And in goroutine; whenever a data comes from `a` or `b` channel, it is passed to the `c` channel. Afterwards, function returns the `c`.

`fanIn` function also takes 2 channels and combines them into 1 channel (c). But while doing like this, it uses 2 goroutines. One goroutine is waiting data from `a`, the other one is waiting data from `b`. 

In `main`, at the beginning there is a multiplexing function `c` is defined. `fainIn` is called with 2 function calls which are return with channels.

These are all channels: 

- `boring("boring!")`
- `boring("yay!")`
- `fanIn(boring("boring!"), boring("yay!"))`

After the multiplexing function, the known `for` loop comes. This time, it iterates 9 times. And formats and prints the data which comes from `c`.