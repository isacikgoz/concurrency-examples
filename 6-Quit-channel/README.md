# 6-Quit-channel

In that `boring`, there is an additional boolean channel named as `quit`. In this goroutine whenever `c` channel receives a `msg`, it does not anything other than formatting and returning data. And if any data receives in the `quit` channel, the goroutine ends and the `boring` returns with receive-only channel `c`.

In `main`, firstly `quit` channel is created. And then `boring` is called. In `for` loop, an iteration number is determined by randomly and the data in the channel is printed. After that, `quit` receives `true`. As a result the program ends.