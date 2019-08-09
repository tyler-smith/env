# env
[![Build Status](https://travis-ci.org/tyler-smith/env.svg?branch=master)](https://travis-ci.org/tyler-smith/env)
[![license](https://img.shields.io/github/license/tyler-smith/env.svg?maxAge=2592000)](https://github.com/tyler-smith/env/blob/master/LICENSE)
[![Documentation](https://godoc.org/github.com/tyler-smith/env?status.svg)](http://godoc.org/github.com/tyler-smith/env)
[![Go Report Card](https://goreportcard.com/badge/github.com/tyler-smith/env)](https://goreportcard.com/report/github.com/tyler-smith/env)
[![GitHub issues](https://img.shields.io/github/issues/tyler-smith/env.svg)](https://github.com/tyler-smith/env/issues)


Functions for easily working with environment variables

## Example
```go
func NewConfigFromEnv() Config {
  return Config{
    APIBindInterface: env.GetString("SEARCH_API_BIND_INTERFACE", "localhost:8080"),
    MySQLDSN:  env.GetString("SEARCH_MYSQL_DSN", "root@tcp(localhost:3306)/yourdb"),
    MaxThreads: env.GetInt("TESTAPP_MAX_THREADS"),
  }
```