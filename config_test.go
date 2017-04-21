package mgobench

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	fname := "test.toml"
	c, err := LoadConfig(fname)
	assert.Nil(t, err, "error loading toml file: %s", err)
	assert.NotNil(t, c, "nil config returned")
	t.Logf("%#v", c)
}