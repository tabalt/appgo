package config

import (
	"testing"
)

func TestInit(t *testing.T) {
	conf = Config{}
	result := conf.Init()
	if !result {
		t.Errorf("Init() failed. Got %v, expected true.", result)
	}
}
