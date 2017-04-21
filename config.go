package mgobench

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Thread int
	Duration int
	Logfile string `toml:"logfile"`
	Mongos []string
	Database string
	Collection string
}

func (c Config) Validate() bool {
	return true
}

func LoadConfig(p string) (Config, error) {
	var c = Config{}

	_, err:= toml.DecodeFile(p, &c)
	if err != nil {
		return c, err
	}
	return c, nil
}