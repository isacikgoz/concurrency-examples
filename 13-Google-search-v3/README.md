# 13-Google-search-v3

Here another version of Google Search example. Differences are:

- `results` in `main` is defined like this:
```
results := Google("golang")
```
- In `Google`, in each goroutine's channels receive the `first` function:
```
go func() { c <- first(query, web, web) }()
go func() { c <- first(query, image, image) }()
go func() { c <- first(query, video, video) }()
```

So, replicas does the same thing as in the previous example. But this time they are doing so in the seperate goroutines. As a result of that, the search time takes shorter.