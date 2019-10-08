# 4-Restoring-sequence

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