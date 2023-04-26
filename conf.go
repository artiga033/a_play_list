package a_play_list

import (
	"github.com/caarlos0/env/v8"
)

type Config struct {
	Endpoint   string `env:"ENDPOINT"`
	User       string `env:"USER"`
	Pass       string `env:"PASS"`
	VideoTypes string `env:"VIDEO_TYPES"`
}

var VideoTypes = make(map[string]struct{})

func init() {
	for _, v := range []string{"mp4", "mkv", "avi", "mov", "rmvb", "webm", "flv"} {
		VideoTypes[v] = struct{}{}
	}
}

func LoadEnvConf(conf *Config) error {
	err := env.ParseWithOptions(conf, env.Options{Prefix: "A_PLAY_LIST_"})
	return err
}
