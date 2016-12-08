# redis-bench
A simple single-thread write/read benchmark program for redis in golang.

##  Start a redis

To benchmark in-memory redis,
```sh
docker run -d -p 6379:6379 redis:alpine
```

To benchmark redis with persistence,
```sh
# Persistence in AOF way. https://redis.io/topics/persistence
docker run -d -v /tmp/redis:/data -p 6379:6379 redis:alpine redis-server --appendonly yes
```

## Build
```sh
go get gopkg.in/redis.v5
# Use mirror of gopkg.in/redis.v5 if go get failed.
# https://github.com/go-redis/redis
go build
```

## Bench
```sh
time ./redis-bench
```
