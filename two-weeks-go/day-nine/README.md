# Day 9 — Sync Primitives

## Project: Thread-Safe Cache with TTL

An in-memory key-value cache that's safe for concurrent use, with automatic expiry of entries after their time-to-live. A background goroutine sweeps expired entries every second.

```bash
go run main.go              # run stress test
go run -race main.go        # run with race detector (requires C compiler)
```

## Project Structure

```
day-nine/
├── go.mod
├── main.go                 # stress test: 100 goroutines hitting the cache
└── models/
    └── cache.go            # Cache struct + Set/Get/Delete + cleanup goroutine
```

### main.go
Launches 100 goroutines that randomly Set, Get, or Delete keys from a small pool of 5 keys. Small key pool forces contention so the mutex actually matters. `wg.Wait()` ensures all operations complete before the program exits.

### models/cache.go
The cache itself. Three exported methods (`Set`, `Get`, `Delete`) and an unexported `cleanup` goroutine that runs for the lifetime of the cache. All internal state is unexported so callers can't bypass the locking.

## What I learned

### Race conditions
Multiple goroutines accessing the same memory where at least one is writing causes unpredictable behavior. `counter++` looks atomic but is actually read-increment-write — two goroutines can overlap and lose increments.

### The race detector
Go has a built-in race detector. Run with `-race`:
```bash
go run -race main.go
go test -race ./...
```
It instruments memory access and reports any concurrent read/write conflict with stack traces. Catches bugs that may not have manifested yet under your current timing.

### sync.Mutex
A basic exclusive lock. Only one goroutine can hold it at a time. Always pair `Lock()` with `defer Unlock()` so it releases even on panic or early return.
```go
mu.Lock()
defer mu.Unlock()
counter++
```

### sync.RWMutex
Two-mode lock — many readers OR one writer.
- `RLock()` / `RUnlock()` — many goroutines can hold simultaneously, for reads
- `Lock()` / `Unlock()` — exclusive, blocks readers and other writers

Better than plain Mutex for read-heavy workloads like a cache where Gets vastly outnumber Sets.

### Mutexes are NOT reentrant
You cannot acquire a lock you already hold — it deadlocks. If you're already in a locked section and need to call another method that locks, work with the underlying data directly:
```go
c.mu.Lock()
defer c.mu.Unlock()
delete(c.data, key)  // direct map delete, NOT c.Delete(key)
```

### sync.Once
Runs a function exactly once even if called from many goroutines. Idiomatic Go lazy singleton:
```go
var once sync.Once
once.Do(func() { config = loadConfigFromDisk() })
```

### sync/atomic
Lockless operations for simple integer values. Faster than mutex-protected counters but only works for single values.
```go
var counter atomic.Int64
counter.Add(1)
counter.Load()
```

### sync.Map vs map + mutex
Plain Go maps panic on concurrent access. Two options:
- **map + RWMutex** — general purpose, more flexible, often faster
- **sync.Map** — optimized for write-once-read-many or disjoint-key workloads

For a general-purpose cache, map + RWMutex is the right choice.

### time.Ticker
For periodic background work, `time.Ticker` is the right tool. It fires on a regular cadence regardless of how long the work takes:
```go
ticker := time.NewTicker(1 * time.Second)
defer ticker.Stop()
for {
    <-ticker.C
    // do periodic work
}
```

### Background goroutine pattern
Launch a goroutine in the constructor so it runs for the lifetime of the object:
```go
func NewCache() *Cache {
    c := &Cache{data: map[string]entry{}}
    go c.cleanup()
    return c
}
```

### delete during range
Go specifically allows `delete(m, key)` during a `range` loop. Other languages would crash with "modified during iteration."

## Key takeaways

- Always run concurrent code with `-race` — bugs that haven't surfaced yet will be caught deterministically
- Pair every `Lock()` with `defer Unlock()` on the next line — no exceptions
- Mutexes are not reentrant — don't call locking methods from inside locked sections
- Keep critical sections as small as possible — every nanosecond inside a lock blocks other goroutines
- Public methods own the locking; internal helpers work with the data directly
- Unexport internal state — exported fields let callers bypass the locking and break the safety guarantees
- RWMutex for read-heavy workloads, plain Mutex for write-heavy or balanced
- Stress tests need a small key pool to force contention — unique keys per goroutine never actually conflict