# 1-Simple

In this example, there is a `boring` function which takes a string msg and a channel of strings as parameters. In that function, there is a `for` loop that takes string msg and iteration number (i), and formats them with `Sprintf`. Sprintf formats and returns a string without printing it anywhere. So we send the result to channel `c`. In the second line of `for` loop, a little wait is added.

In the `main` function, a channel of string `c` is created. And after the channel is created, a string message and the channel is sent to `boring` function. Here, the point to be considered is `boring` is invoked in a goroutine. If a function is called with the `go` keyword, a goroutine works in the background concurrently:

`go boring("boring!", c)`

And also there is a `for` loop with 5 iteration in the `main`. In that loop, the data which comes from the channel is printed. Finally, a leaving message is printed. 