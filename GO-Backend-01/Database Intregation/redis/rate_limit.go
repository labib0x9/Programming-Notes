package redis
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var luaCode = `
	local key = KEYS[1]
	local capacity = tonumber(ARGV[1])
	local rate = tonumber(ARGV[2])
	local now = tonumber(ARGV[3])

	local data = redis.call("HMGET", key, "token", "last_refill")
	local token = tonumber(data[1])
	local last_refill = tonumber(data[2])

	if token == nil then
		token = rate
		last_refill = now
	end

	local diff = (now - last_refill) / 1000.0
	local add_token = diff * rate
	token = math.min(capacity, token + add_token)
	last_refill = now

	if token >= 1 then
		token = token - 1
		redis.call("HMSET", key, "token", token, "last_refill", last_refill)
		redis.call("PEXPIRE", key, 6000)
		return {1, 0}
	end

	local token_need = 1 - token
	local wait_ms = math.ceil((token_need / rate) * 1000) 

	redis.call("HMSET", key, "token", token, "last_refill", last_refill)
	redis.call("PEXPIRE", key, 6000)

	return {0, wait_ms}
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
		2,
		1,
		time.Now().UnixMilli(),
	).Result()

	if err != nil {
		panic(err)
	}

	data := res.([]interface{})
	allowed := data[0].(int64)
	result := data[1].(int64)

	if allowed == 1 {
		fmt.Println("OK")
		fmt.Println(result)
		fmt.Printf("RAW: %#v, TYPE: %T\n", data[1], data[1])
	} else {
		fmt.Println("NOT OK")
		fmt.Println("WAIT", result)
	}
}
