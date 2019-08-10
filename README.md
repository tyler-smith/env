# env
[![Build Status](https://travis-ci.org/tyler-smith/env.svg?branch=master)](https://travis-ci.org/tyler-smith/env)
[![license](https://img.shields.io/github/license/tyler-smith/env.svg?maxAge=2592000)](https://github.com/tyler-smith/env/blob/master/LICENSE)
[![Documentation](https://godoc.org/github.com/tyler-smith/env?status.svg)](http://godoc.org/github.com/tyler-smith/env)
[![Go Report Card](https://goreportcard.com/badge/github.com/tyler-smith/env)](https://goreportcard.com/report/github.com/tyler-smith/env)
[![GitHub issues](https://img.shields.io/github/issues/tyler-smith/env.svg)](https://github.com/tyler-smith/env/issues)


Functions for easily working with environment variables

## Example
```go
package main

// Config represents the information needed to connected to your backing
// services; layer IV of the 12 factor app.
type Config struct {
	APIBindInterface string
	MySQLDSN         string
	MaxThreads       string
}

// NewConfigFromEnv creates a new `Config` populated with information from the
// local environment using the "env" lib
func NewConfigFromEnv() Config {
	return Config{
		APIBindInterface: env.GetString("TESTAPP_API_BIND_INTERFACE", "localhost:8080"),
		MySQLDSN:         env.GetString("TESTAPP_MYSQL_DSN", "root@tcp(localhost:3306)/yourdb"),
		MaxThreads:       env.GetInt("TESTAPP_MAX_THREADS", 10),
	}
}

```