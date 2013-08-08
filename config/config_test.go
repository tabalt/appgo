package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	conf = Instance("config_test.json")
	if !conf {
		t.Errorf("Init() failed. Got %v, expected true.", result)
	}
}
