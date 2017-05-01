package mgobench

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Thread   int
	Duration int

	Testcases map[string]testcase
	Influxdb  influx
	Mongo     mongo
}

type testcase struct {
	Name     string
	Duration string
}

type influx struct {
	ConnectionString string `toml:"connection_string"`
	Database         string
}

type mongo struct {
	ConnectionString string `toml:"connection_string"`
	Database         string
	Collection       string
}

func (c Config) Validate() bool {
	// TODO : validate config struct
	return true
}

func LoadConfig(p string) (Config, error) {
	var c = Config{}

	_, err := toml.DecodeFile(p, &c)
	if err != nil {
		return c, err
	}
	return c, nil
}
