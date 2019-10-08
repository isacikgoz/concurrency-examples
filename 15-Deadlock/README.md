# 15-Deadlock

In this example, we see a `deadlock`. Remember the previous example. We were sending a `ball` to the `table` to start the ping-pong game:
```
table <- new(ball)
```
If we don't sent `ball` to the `table`, the goroutines wait for nothing. So they asleep and deadlock occurs.