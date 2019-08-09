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

func TestGetString(t *testing.T) {
	cleanTestEnv()

	assertEqual(t, "default", GetString("ENV_TEST", "default"))
	os.Setenv("ENV_TEST", "not-default")
	assertEqual(t, "not-default", GetString("ENV_TEST", "not-default"))
}

func TestGetInt(t *testing.T) {
	cleanTestEnv()

	val, err := GetInt("ENV_TEST", 42)
	assertEqual(t, nil, err)
	assertEqual(t, 42, val)

	os.Setenv("ENV_TEST", strconv.Itoa(13))

	val, err = GetInt("ENV_TEST", 42)
	assertEqual(t, nil, err)
	assertEqual(t, 13, val)
}

func TestGetInt64(t *testing.T) {
	cleanTestEnv()

	val, err := GetInt("ENV_TEST", 42)
	assertEqual(t, nil, err)
	assertEqual(t, 42, val)

	os.Setenv("ENV_TEST", strconv.Itoa(13))

	val, err = GetInt("ENV_TEST", 42)
	assertEqual(t, nil, err)
	assertEqual(t, 13, val)
}

func TestGetBool(t *testing.T) {
	cleanTestEnv()

	val, err := GetBool("ENV_TEST", false)
	assertEqual(t, nil, err)
	assertEqual(t, false, val)

	os.Setenv("ENV_TEST", "true")

	val, err = GetBool("ENV_TEST", false)
	assertEqual(t, nil, err)
	assertEqual(t, true, val)
}

func TestGetFloat64(t *testing.T) {
	cleanTestEnv()

	val, err := GetFloat64("ENV_TEST", 42.7)
	assertEqual(t, nil, err)
	assertEqual(t, 42.7, val)

	os.Setenv("ENV_TEST", strconv.FormatFloat(13.3, 'f', 12, 64))

	val, err = GetFloat64("ENV_TEST", 42.7)
	assertEqual(t, nil, err)
	assertEqual(t, 13.3, val)
}

func TestGetFloat32(t *testing.T) {
	cleanTestEnv()

	val, err := GetFloat32("ENV_TEST", 42.7)
	assertEqual(t, nil, err)
	assertEqual(t, float32(42.7), val)

	os.Setenv("ENV_TEST", strconv.FormatFloat(13.3, 'f', 12, 32))

	val, err = GetFloat32("ENV_TEST", 42.7)
	assertEqual(t, nil, err)
	assertEqual(t, float32(13.3), val)

	// Now test an invalid float
	os.Setenv("ENV_TEST", "asdf")
	_, err = GetFloat32("ENV_TEST", 42.7)
	assertNotEqual(t, nil, err)
	assertEqual(t, true, strings.Contains(err.Error(), "strconv.ParseFloat"))
}

func TestGetTimestamp(t *testing.T) {
	cleanTestEnv()

	time1 := time.Now()
	time2, err := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	assertEqual(t, nil, err)

	val, err := GetTimestamp("ENV_TEST", time1)
	assertEqual(t, nil, err)
	assertEqual(t, time1, val)

	os.Setenv("ENV_TEST", time1.Format(time.RFC3339))

	val, err = GetTimestamp("ENV_TEST", time1)
	assertEqual(t, nil, err)
	if val.UTC().Equal(time2) {
		t.Fatal("Wanted", time2, "but got:", time1)
	}
}

func TestGetStringSlice(t *testing.T) {
	cleanTestEnv()

	val := GetStringSlice("ENV_TEST", []string{"a", "b", "c"})
	for i, v := range []string{"a", "b", "c"} {
		assertEqual(t, val[i], v)
	}
	os.Setenv("ENV_TEST", "x,y,z")

	val = GetStringSlice("ENV_TEST", []string{"a", "b", "c"})
	for i, v := range []string{"x", "y", "z"} {
		assertEqual(t, val[i], v)
	}
}

func cleanTestEnv() {
	os.Unsetenv("ENV_TEST")
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatal("Wanted:", a, "but got:", b)
	}
}

func assertNotEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		t.Fatal("Wanted value to not be", a, "but it was")
	}
}
