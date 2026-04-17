package main

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var luaCode = `
	local key = KEYS[1]
	local value = ARGV[1]
	local data = redis.call("GET", key)

	if data == nil then
		redis.call("SET", key, value)
		redis.call("PEXPIRE", key, 6000)
		return {0, value}
	end

	return {1, data}
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

	key := "test:localhost:user-agent"

	res, err := script.Run(
		ctx,
		client,
		[]string{key},
		time.Now().UnixMilli(),
	).Result()

	if err != nil {
	}

	data := res.([]interface{})
	allowed := data[0].(int64)
	result := data[1].(string)

	_ = allowed
	_ = result
}
