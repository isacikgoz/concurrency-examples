# 14-Ping-pong

Here we have a `ball` struct that include `hits` number of `int`.

In `player` function, there are 2 arguments. One is `name` string and the other is `table` as a channel of `ball pointer`. This function contains a `for` loop:

- `ball` initialized to data from `table` channel with short declaration.
- number of hits are incremented by 1.
- name and number of hits are printed.
- a little waiting 
- ball is sent to the `table` channel.

In `main`, 

- `table` channel is generated as `ball pointer`.
- 2 different goroutines starts with `player` functions.
- new `ball` struct is sent to the `table`.
- waits 1 seconds
- `table` channel sends its data. So, we read the data from that channel.

This code acts like 2 players play ping-pong with the help of 2 goroutines. In every step, the number of hits are increased by 1 and the player is changes.