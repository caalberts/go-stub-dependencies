package main

import "github.com/go-redis/redis"

type RedisProxy interface {
	HGetAll(string) (map[string]string, error)
	HMSet(string, map[string]interface{}) (string, error)
}

type Redis struct {
	*redis.Client
}

func (r *Redis) HGetAll(key string) (map[string]string, error) {
	return r.Client.HGetAll(key).Result()
}

func (r *Redis) HMSet(key string, data map[string]interface{}) (string, error) {
	return r.Client.HMSet(key, data).Result()
}

