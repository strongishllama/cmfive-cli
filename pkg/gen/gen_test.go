package gen_test

import (
	"os"
	"testing"
)

func setup(t *testing.T) {
	_ = os.RemoveAll("modules")
}

func teardown(t *testing.T) {
	_ = os.RemoveAll("modules")
}
