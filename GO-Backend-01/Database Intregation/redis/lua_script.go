package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// data becomes false if key is missing
// data is stored value if key isn't missing.

var luaCode = `
	local key = KEYS[1]
	local value = ARGV[1]
	local data = redis.call("GET", key)

	if data == false then
		redis.call("SET", key, value)
		redis.call("PEXPIRE", key, 6000)
		return {0, value}
	end

	return {1, tostring(data)}
`

var script = redis.NewScript(luaCode)

func main() {
	client := redis.NewClient(
		&redis.Options{
			Addr: "localhost:6379",
		},
	)
	defer client.Close()

	ctx := context.Background()

	if err := client.Ping(ctx).Err(); err != nil {
		panic(err)
	}

	key := "test:localhost:user-agent"

	res, err := script.Run(
		ctx,
		client,
		[]string{key},
		time.Now().UnixMilli(),
	).Result()

	if err != nil {
		panic(err)
	}

	data := res.([]interface{})
	allowed := data[0].(int64)
	result := data[1].(string)

	// _ = allowed
	// _ = result

	if allowed == 1 {
		fmt.Println("OK")
		fmt.Println(result)
		fmt.Printf("RAW: %#v, TYPE: %T\n", data[1], data[1])
	} else {
		fmt.Println("NOT OK")
	}
}
