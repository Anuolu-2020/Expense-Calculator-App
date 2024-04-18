package pkg

import "os"

type Env struct{}

func (e Env) GetDbUrl() string {
	DATABASE_URL := os.Getenv("DATABASE_URL")
	return DATABASE_URL
}

func (e Env) GetGoEnv() string {
	GO_ENV := os.Getenv("GO_ENV")
	return GO_ENV
}

func (e Env) GetRedisUrl() string {
	REDIS_URL := os.Getenv("REDIS_URL")
	return REDIS_URL
}

func (e Env) GetRedisPass() string {
	REDIS_PASS := os.Getenv("REDIS_PASS")
	return REDIS_PASS
}

func (e Env) GetSessionDBUrl() string {
	SESSION_DB := os.Getenv("SESSION_DB")
	return SESSION_DB
}
