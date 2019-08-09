package env

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// GetString gets the environment variable with the given name and parses it as
// a string.
func GetString(name string, defaultVal string) string {
	val := os.Getenv(name)
	if val == "" {
		return defaultVal
	}
	return val
}

// GetInt64 gets the environment variable with the given name and parses it as
// a an int64
func GetInt64(name string, defaultVal int64) (int64, error) {
	val := os.Getenv(name)
	if val == "" {
		return defaultVal, nil
	}
	return strconv.ParseInt(val, 10, 64)
}

// GetInt gets the environment variable with the given name and parses it as
// an int.
func GetInt(name string, defaultVal int64) (int, error) {
	i, err := GetInt64(name, defaultVal)
	return int(i), err
}

// GetBool gets the environment variable with the given name and parses it as a
// bool.
func GetBool(name string, defaultVal bool) (bool, error) {
	val := os.Getenv(name)
	if val == "" {
		return defaultVal, nil
	}
	return strconv.ParseBool(val)
}

// GetFloat64 gets the environment variable with the given name and parses it as
// a float64.
func GetFloat64(name string, defaultVal float64) (float64, error) {
	val := os.Getenv(name)
	if val == "" {
		return defaultVal, nil
	}
	return strconv.ParseFloat(val, 64)
}

// GetFloat32 gets the environment variable with the given name and parses it as
// a float32.
func GetFloat32(name string, defaultVal float32) (float32, error) {
	val := os.Getenv(name)
	if val == "" {
		return defaultVal, nil
	}

	i, err := strconv.ParseFloat(val, 32)
	if err != nil {
		return 0, err
	}

	return float32(i), nil
}

// GetTimestamp gets the environment variable with the given name and parses it
// as a timestamp in RFC3339 format.
func GetTimestamp(name string, defaultVal time.Time) (time.Time, error) {
	val := os.Getenv(name)
	if val == "" {
		return defaultVal, nil
	}

	return time.Parse(time.RFC3339, val)
}

// GetStringSlice gets the environment variable with the given name and parses
// it as a slice of strings.
func GetStringSlice(name string, defaultVal []string) []string {
	val := os.Getenv(name)
	if val == "" {
		return defaultVal
	}
	return strings.Split(val, ",")
}

// TOOD: Add GetIntSlice, GetFloat64Slice, and GetFloat32Slie
