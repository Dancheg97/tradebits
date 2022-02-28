package config

import "testing"

func TestReadConfig(t *testing.T) {
	test_config := GetConfig("config.json")
	if test_config.Dgraph != "localhost:9080" {
		t.Error("config was read badly")
	}
}

func TestReadUnexistingConfig(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	GetConfig(".s")
}

func TestReadBadConig(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	GetConfig("config.go")
}
