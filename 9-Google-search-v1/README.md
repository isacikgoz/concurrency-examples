# 9-Google-search-v1

In this example, a set of `var` is defined (`web`, `image` and `video`). These are defined as `FakeSearch` function.

- `FakeSearch` function takes a `kind` string and returns the `Search`.
    - `Search` is defined as a type of a function which takes a `query` string parameter and returns the `Result`.
        - `Result` is defined as a string type  
    - Consequently, the `FakeSearch` sleeps for a while and returns whatever the `Result` is.

`main` function, 

- starts with `rand.Seed(time.Now().Unix())`. In `GO` language, `rand` generates the same number each time the program runs. To prevent this, we can use `seed`. By doing that, every time when program runs, a different value is produced. 
- After that, a `start` time is defined. 
- `results` is defined as short declaration. It initialized with the `Google` function.
    - `Google` function takes a string `query` parameter and returns `results` (an array of `Result` type).
        - In `Google` function, some `append` functions are performed, respectively `web("golang")`, `image("golang")`, `video("golang")` are appended to `results`.
- `elapsed` time is defined as since `start` time.

Finally, the `results` and the `elapsed` time are printed.