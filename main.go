package main

import "github.com/zyfdegh/redis-bench/db"

func main() {
	db.Ping()
	db.Bench()
}
