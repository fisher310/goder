package main

import (
	"github.com/kylelemons/go-gypsy/yaml"
	"testing"
)

func TestYaml(t *testing.T) {
	config, err := yaml.ReadFile("config.yaml")
	if err != nil {
		t.Errorf("read the file error: %s\n", err )
	}

	enabled, err := config.GetBool("enabled")
	if err != nil || !enabled {
		t.Errorf("expected enabled is %v, but actual is %v, and the error is %s", true, enabled, err)
	}
	path, err := config.Get("path")
	if err != nil {
		t.Errorf("get path err: %s", err)
	}

	expectedPath := "/usr/local/path"
	if path != expectedPath {
		t.Errorf("path interceptor error, exepcted is %s, but the actual is %s\n", expectedPath, path)
	}
}
