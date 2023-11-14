package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type EnvVars struct {
	DSN           string
	RedisAddr     string
	RedisPassword string
	RedisDb       int
}

func LoadEnv() EnvVars {
	if err := godotenv.Load(); err != nil {
		panic("cannot load the config file as env variables")
	}

	dsn := os.Getenv("DSN")
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPass := os.Getenv("REDIS_PASSWORD")
	redisDb := os.Getenv("REDIS_DB")
	parsedRedisDb, err := strconv.Atoi(redisDb)
	if err != nil {
		panic("cannot parse redis DB number")
	}

	return EnvVars{
		DSN:           dsn,
		RedisAddr:     redisAddr,
		RedisPassword: redisPass,
		RedisDb:       parsedRedisDb,
	}
}
