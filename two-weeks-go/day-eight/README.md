# Day 8 — Goroutines & Channels

## Project: Parallel Downloader

A CLI tool that fetches multiple URLs concurrently using goroutines and reports each result's status code and response size. Includes a sequential version for benchmarking the speedup.

```bash
go run main.go
```

## What I learned

### Goroutines
A goroutine is a function running concurrently with other code. You start one by putting `go` in front of a function call. Goroutines are not threads — they're lighter, multiplexed by Go's runtime onto a small number of OS threads.
```go
go func() {
    fmt.Println("running concurrently")
}()
```

### main doesn't wait
The program exits as soon as `main()` returns — it does NOT wait for goroutines. Without synchronization your goroutines may never run.

### sync.WaitGroup
The correct way to wait for goroutines to finish. It's a counter — `Add` before launching, `Done` when the goroutine finishes, `Wait` blocks until the counter is zero.
```go
var wg sync.WaitGroup
for _, url := range urls {
    wg.Add(1)
    go func(u string) {
        defer wg.Done()
        // work
    }(url)
}
wg.Wait()
```

### Channels
Typed pipes that let goroutines pass values safely. Send with `ch <- value`, receive with `<-ch`. Unbuffered channels block until both sender and receiver are ready.
```go
results := make(chan Result, len(urls))  // buffered channel
results <- result                         // send
r := <-results                            // receive
```

### Buffered vs unbuffered
- **Unbuffered** (`make(chan T)`) — send blocks until a receiver is ready
- **Buffered** (`make(chan T, 10)`) — send only blocks when buffer is full

For collecting results from N goroutines, a buffered channel of size N means no goroutine blocks on send.

### The close + range pattern
After all sends are done, `close()` the channel. Ranging over a closed channel reads remaining buffered values then exits the loop. Only the sender should close.
```go
wg.Wait()       // wait for all sends to finish
close(results)  // signal no more values coming
for r := range results {
    fmt.Println(r)
}
```

**Order matters:** Wait → Close → Range. Closing before all sends finish causes a panic.

### defer placement
Register cleanup as soon as you have the resource — not at the end of the function. If a later step fails and returns early, deferred cleanup at the end never runs.
```go
response, err := http.Get(u)
if err != nil {
    return
}
defer response.Body.Close()  // right after success check, not at the end
```

### Goroutine loop variable
Pass loop variables as arguments to the goroutine so the value is captured at launch time. Without this the goroutine may see whatever value the variable has when it eventually runs.
```go
for _, url := range urls {
    go func(u string) {  // u is fixed at this URL
        // ...
    }(url)
}
```

## The Speedup

Sequential vs concurrent fetching 10 URLs from jsonplaceholder.typicode.com:

```
Go routines executed in : 88.7779ms
{https://jsonplaceholder.typicode.com/todos/9 200 92 <nil>}
{https://jsonplaceholder.typicode.com/todos/7 200 98 <nil>}
{https://jsonplaceholder.typicode.com/todos/5 200 128 <nil>}
{https://jsonplaceholder.typicode.com/todos/1 200 83 <nil>}
{https://jsonplaceholder.typicode.com/todos/6 200 114 <nil>}
{https://jsonplaceholder.typicode.com/todos/0 404 2 <nil>}
{https://jsonplaceholder.typicode.com/todos/4 200 80 <nil>}
{https://jsonplaceholder.typicode.com/todos/8 200 92 <nil>}
{https://jsonplaceholder.typicode.com/todos/2 200 99 <nil>}
{https://jsonplaceholder.typicode.com/todos/3 200 84 <nil>}
gets without go routines took:  120.4167ms
{https://jsonplaceholder.typicode.com/todos/0 404 2 <nil>}
{https://jsonplaceholder.typicode.com/todos/1 200 83 <nil>}
{https://jsonplaceholder.typicode.com/todos/2 200 99 <nil>}
{https://jsonplaceholder.typicode.com/todos/3 200 84 <nil>}
{https://jsonplaceholder.typicode.com/todos/4 200 80 <nil>}
{https://jsonplaceholder.typicode.com/todos/5 200 128 <nil>}
{https://jsonplaceholder.typicode.com/todos/6 200 114 <nil>}
{https://jsonplaceholder.typicode.com/todos/7 200 98 <nil>}
{https://jsonplaceholder.typicode.com/todos/8 200 92 <nil>}
{https://jsonplaceholder.typicode.com/todos/9 200 92 <nil>}
```

**~27% faster** even with only 10 fast requests. With slower endpoints or more URLs the gap widens dramatically — sequential time scales linearly while concurrent stays roughly constant.

Also notice the concurrent results return out of order (9, 7, 5, 1...) — goroutines finish whenever the network responds, not in the order they were launched. This is a feature of concurrent code, not a bug.

## Key takeaways

- `main()` exits immediately — you must synchronize with WaitGroup or channels to wait for goroutines
- Always `defer wg.Done()` at the top of the goroutine so it runs no matter how the goroutine exits
- `wg.Add(1)` goes BEFORE launching the goroutine, not inside it (race condition)
- Order is critical: Wait → Close → Range. Close before Wait causes a panic
- Defer placement matters — register cleanup right after acquiring the resource, not at the end of the function
- Concurrent code returns results in arrival order, not launch order
- The speedup comes from doing the network I/O wait in parallel — CPU-bound work won't see the same gain