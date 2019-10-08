# 8-Daisy-chain

Here is a `f` function that takes 2 integer channels as parameters. It takes the value from `right` channel, add `1` to that value and sends it to the `left` channel. 

In `main`, a constant `n` is declared as 100000. The leftmost of daisy chain is generated as a channel. `right` and `left` are initialized to `leftmost`. After that, in the `for` loop the chain starts to form:

- `right` channel is reassigned to a new channel
- a goroutine starts with `f` function
- `left` channel is set as the value of `right` channel  

After the `for` loop, another goroutine starts with a function. That function takes an integer channel as parameter and sends `1` to that channel. And than returns the `right` channel.

Finally, `leftmost` is printed. We will see where the chain ends up.