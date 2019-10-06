# concurrency-examples

`concurrency-examples` contains some examples about concurrency phenomenon in `GO` language

## 1-Simple

In this example, there is a `boring` function which takes a string msg and a channel of strings as parameters. In that function, there is a `for` loop that takes string msg and iteration number (i), and formats them with `Sprintf`. Sprintf formats and returns a string without printing it anywhere. So we send the result to channel `c`. In the second line of `for` loop, a little wait is added.

In the `main` function, a channel of string `c` is created. And after the channel is created, a string message and the channel is sent to `boring` function. Here, the point to be considered is `boring` is invoked in a goroutine. If a function is called with the `go` keyword, a goroutine works in the background concurrently:

`go boring("boring!", c)`

And also there is a `for` loop with 5 iteration in the `main`. In that loop, the data which comes from the channel is printed. Finally, a leaving message is printed. 

## 2-Generator

In this example's `boring` function is differentiate from first example's. This time, the function returns receive-only channel of strings. The channel is created in the `boring` function and the goroutine is launched in that function. (The same `for` loop as in previous example's `boring` function.) 

In `main`, firstly a function is generated named as `c`. And the same `for` loop as in previous example's `main` is performed.

## 3-Multiplexing

Here, again a `boring` function is the same as previous `generator pattern`. But also here has `fan in pattern`: a function that returns a channel.

`fanInSelect` function takes 2 channels of strings as parameters (a and b) and returns a channel of strings (c). Also it contains a single goroutine. In `fanInSelect`, a channel (c) is generated. And in goroutine; whenever a data comes from `a` or `b` channel, it is passed to the `c` channel. Afterwards, function returns the `c`.

`fanIn` function also takes 2 channels and combines them into 1 channel (c). But while doing like this, it uses 2 goroutines. One goroutine is waiting data from `a`, the other one is waiting data from `b`. 

In `main`, at the beginning there is a multiplexing function `c` is defined. `fainIn` is called with 2 function calls which are return with channels.

These are all channels: 

- `boring("boring!")`
- `boring("yay!")`
- `fanIn(boring("boring!"), boring("yay!"))`

After the multiplexing function, the known `for` loop comes. This time, it iterates 9 times. And formats and prints the data which comes from `c`.

## 4-Restoring-sequence

Restoring sequencing is sending a channel to a channel making goroutine and wait its turn.

In this example, a `message` struct is defined. It consist of a `str` string and a `wait` channel of boolean. That channel acts as a signalier, it will block on the wait channel until other person says "OK I want you to go ahead".

In `boring`, again a string `msg` parameter comes in and a channel output comes out. There are some differences in `boring` function. First is; after the `c` channel generation, there is another channel `waitForIt` is generated. That channel is shared between all messages. And the other is; inside the `for` loop, the data which is sent to the channel is in `message` type. So, the string and the boolean values are sent to `c`.

Again the same `fanIn` as in the `3-Multiplexing` is present.

`main` starts with multiplexing function. After that a different `for` loop comes. In the loop, respectively:

- a `message` type `msg1` is created and takes the data from `c` channel
- `msg1`'s string part is printed 
- another `message` type `msg2` is created and takes the data from `c` channel
- `msg2`'s string part is printed
- `msg1`'s boolean part is changed as `true`
- `msg2`'s boolean part is changed as `true`

And finally, a leaving message is printed.

## 5-Timeout

In this example, `boring` function takes a string `msg` parameter and returns receive-only channel of strings again. It launches goroutine inside the function as we have seen.

In `main`, a channel is generated with `boring` function. And then a time out value is declared, it is defined as 3 seconds. After that, in `for` loop; whenever `c` channel receives a data, it is printed until the time out `to` occurs.

## 6-Quit-channel

In that `boring`, there is an additional boolean channel named as `quit`. In this goroutine whenever `c` channel receives a `msg`, it does not anything other than formatting and returning data. And if any data receives in the `quit` channel, the goroutine ends and the `boring` returns with receive-only channel `c`.

In `main`, firstly `quit` channel is created. And then `boring` is called. In `for` loop, an iteration number is determined by randomly and the data in the channel is printed. After that, `quit` receives `true`. As a result the program ends.

## 7-Receive-on-quit

In this example, the program says it's done and then finishes.

In `boring`, everything is the same as the previous example except for `quit`. This time, `quit` is a channel of strings not booleans. And in `quit` case of `select`, before it returns, there is a function call (`cleanup`) and after that `quit` receives a string message `see ya!`.

There is also a `cleanup` function as mentioned above. In this function, an `all clean` message is printed. 

In `main`, differently from the previous example, `bye` string is sent to `quit` channel. And the leaving message contains a data from `quit` channel. But after execution of that example, we cannot see the `bye`. `quit` channel receives `see ya!` string in `boring`.

## 8-Daisy-chain

Here is a `f` function that takes 2 integer channels as parameters. It takes the value from `right` channel, add `1` to that value and sends it to the `left` channel. 

In `main`, a constant `n` is declared as 100000. The leftmost of daisy chain is generated as a channel. `right` and `left` are initialized to `leftmost`. After that, in the `for` loop the chain starts to form:

- `right` channel is reassigned to a new channel
- a goroutine starts with `f` function
- `left` channel is set as the value of `right` channel  

After the `for` loop, another goroutine starts with a function. That function takes an integer channel as parameter and sends `1` to that channel. And than returns the `right` channel.

Finally, `leftmost` is printed. We will see where the chain ends up.

## 9-Google-search-v1



## 10-Google-search-v2



## 11-Google-search-v2.1



## 12-Avoid-timeout



## 13-Google-search-v3



## 14-Ping-pong



## 15-Deadlock



## 16-Task-queue
