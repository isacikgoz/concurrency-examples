# 16-Task-queue

In this example, there are some `constants`: number of workers and number of tasks.

A `task` struct which contains `id` string is defined.

`process` function takes `t` task as parameter. And prints `id` of `processing task`.

`worker` function takes a channel of `task`. And whenever channel receives a task, the worker processes that task.

`getTask` function has no arguments and returns an array of tasks. Inside the function, `tasks` initialized as a zero lenght slice. And the number of tasks as numTasks is appended to the `tasks` slice by specifying the `id`.
- To specify the `id`; a randomly `int` is converted to `string` with `strconv.Itoa(rand.Intn(1e9))`. `Itoa` means integer to string conversion.

In `main`, 

- a buffered channel is generated.
- goroutines are created for each workers. 
`hellaTasks` is initialized as `getTasks()`.
- number of processed tasks initialized to 0 (`processeds`). 
- Then we have a different `for` loop definition: 

    `for _, task := range hellaTasks `

    - Above declaration means; iterate this loop for every `hellaTasks` member. But I don't consider the iteration number so let's underscore it (`_`).

    - Back to the `for` loop; each time that a channel receives a task, `processeds` is increased by 1.

- And closing the channel, so that no more values will be sent on it.
- Finally, the total number of processed tasks is printed.