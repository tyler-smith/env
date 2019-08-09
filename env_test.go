package env

// In each test we'll unset the env var and test returning the default. Then
// we'll set the env var and check that the proper value is returned

import (
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func cleanTestEnv() {
	os.Unsetenv("ENV_TEST")
}

func TestGetString(t *testing.T) {
	cleanTestEnv()

	val := GetString("ENV_TEST", "default")
	if val != "default" {
		t.Fatal("Wanted: 'default' but got :", val)
	}

	os.Setenv("ENV_TEST", "not-default")

	val = GetString("ENV_TEST", "not-default")
	if val != "not-default" {
		t.Fatal("Wanted: 'not-default' but got :", val)
	}
}

func TestGetInt(t *testing.T) {
	cleanTestEnv()

	val, err := GetInt("ENV_TEST", 42)
	if err != nil {
		t.Fatal("No error was expected but got:", err)
	}
	if val != 42 {
		t.Fatal("Wanted: 'default' but got :", val)
	}

	os.Setenv("ENV_TEST", strconv.Itoa(13))

	val, err = GetInt("ENV_TEST", 42)
	if err != nil {
		t.Fatal("No error was expected but got:", err)
	}
	if val != 13 {
		t.Fatal("Wanted: 14 but got :", val)
	}
}

func TestGetInt64(t *testing.T) {
	cleanTestEnv()

	val, err := GetInt("ENV_TEST", 42)
	if err != nil {
		t.Fatal("No error was expected but got:", err)
	}
	if val != 42 {
		t.Fatal("Wanted: 'default' but got :", val)
	}

	os.Setenv("ENV_TEST", strconv.Itoa(13))

	val, err = GetInt("ENV_TEST", 42)
	if err != nil {
		t.Fatal("No error was expected but got:", err)
	}
	if val != 13 {
		t.Fatal("Wanted: 14 but got :", val)
	}
}

func TestGetBool(t *testing.T) {
	cleanTestEnv()

	val, err := GetBool("ENV_TEST", false)
	if err != nil {
		t.Fatal("No error was expected but got:", err)
	}
	if val {
		t.Fatal("Wanted false but got :", val)
	}

	os.Setenv("ENV_TEST", "true")

	val, err = GetBool("ENV_TEST", false)
	if err != nil {
		t.Fatal("No error was expected but got:", err)
	}
	if !val {
		t.Fatal("Wanted: false but got :", val)
	}
}

func TestGetFloat64(t *testing.T) {
	cleanTestEnv()

	val, err := GetFloat64("ENV_TEST", 42.7)
	if err != nil {
		t.Fatal("No error was expected but got:", err)
	}
	if val != 42.7 {
		t.Fatal("Wanted: 'default' but got :", val)
	}

	os.Setenv("ENV_TEST", strconv.FormatFloat(13.3, 'f', 12, 64))

	val, err = GetFloat64("ENV_TEST", 42.7)
	if err != nil {
		t.Fatal("No error was expected but got:", err)
	}
	if val != 13.3 {
		t.Fatal("Wanted: 14 but got :", val)
	}
}

func TestGetFloat32(t *testing.T) {
	cleanTestEnv()

	val, err := GetFloat32("ENV_TEST", 42.7)
	if err != nil {
		t.Fatal("No error was expected but got:", err)
	}
	if val != 42.7 {
		t.Fatal("Wanted: 'default' but got :", val)
	}

	os.Setenv("ENV_TEST", strconv.FormatFloat(13.3, 'f', 12, 32))

	val, err = GetFloat32("ENV_TEST", 42.7)
	if err != nil {
		t.Fatal("No error was expected but got:", err)
	}
	if val != 13.3 {
		t.Fatal("Wanted: 14 but got :", val)
	}

	// Now test an invalid float
	os.Setenv("ENV_TEST", "asdf")
	_, err = GetFloat32("ENV_TEST", 42.7)
	if err == nil || !strings.Contains(err.Error(), "strconv.ParseFloat") {
		t.Fatal("A parse error was expected but got:", err)
	}
}

func TestGetTimestamp(t *testing.T) {
	cleanTestEnv()

	time1 := time.Now()
	time2, err := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	if err != nil {
		t.Fatal("No error was expected but got:", err)
	}

	val, err := GetTimestamp("ENV_TEST", time1)
	if err != nil {
		t.Fatal("No error was expected but got:", err)
	}
	if !val.Equal(time1) {
		t.Fatal("Wanted:", time1, "but got :", val)
	}

	os.Setenv("ENV_TEST", time1.Format(time.RFC3339))

	val, err = GetTimestamp("ENV_TEST", time1)
	if err != nil {
		t.Fatal("No error was expected but got:", err)
	}
	if val.UTC().Equal(time2) {
		t.Fatal("Wanted: 14 but got :", val)
	}
}

func TestGetStringSlice(t *testing.T) {
	cleanTestEnv()

	val := GetStringSlice("ENV_TEST", []string{"a", "b", "c"})
	for i, v := range val {
		if val[i] != v {
			t.Fatal("Wanted:", val[i], "but got:", v)
		}
	}
	os.Setenv("ENV_TEST", "x,y,z")

	val = GetStringSlice("ENV_TEST", []string{"a", "b", "c"})
	for i, v := range []string{"x", "y", "z"} {
		if val[i] != v {
			t.Fatal("Wanted:", val[i], "but got:", v)
		}
	}
}
