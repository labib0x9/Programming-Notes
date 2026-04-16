package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type Value struct {
	X int    `json:"x"`
	Y string `json:"y"`
}

func main() {
	client := redis.NewClient(
		&redis.Options{
			Addr: "localhost:6379",
		},
	)
	defer client.Close()

	ctx := context.Background()

	// Simple key-value, string return...
	err := client.Set(
		ctx,
		"key",
		123,
		0,
	).Err()

	if err != nil {

	}

	res, err := client.Get(
		ctx,
		"key",
	).Result()

	if err != nil {

	}
	_ = res

	// struct{} cache, use serialization to store struct, here is json

	data, err := json.Marshal(
		Value{
			X: 10,
			Y: "Hello",
		},
	)

	err = client.Set(
		ctx,
		"key-struct",
		data,
		time.Minute,
	).Err()

	if err != nil {

	}

	ress, err := client.Get(
		ctx,
		"key",
	).Bytes()

	var v Value
	err = json.Unmarshal(ress, &v)

	// use hash
	client.HSet(
		ctx,
		"key",
		"x", 10,
		"y", "hello",
	)
}
