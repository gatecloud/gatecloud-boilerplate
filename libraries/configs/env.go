package configs

import (
	"log"
	"os"
	"strings"
)

type Env struct {
	Config
	DbEngine string `json:"db_engine"`
	DbConn   string `json:"db_conn"`
}

func (env *Env) GetPrefix() string {
	return env.Prefix
}

func (env *Env) SetPrefix(prefix string) {
	env.Prefix = prefix
}

func (env *Env) SetPath(path string) error {
	if path == "" {
		dir, err := os.Getwd()
		if err != nil {
			return nil
		}

		idx := strings.LastIndex(dir, string(os.PathSeparator))
		if idx < 1 {
			log.Fatal("can't find \\")
		}

		env.Path = dir[0:idx] + "/libraries/configs/env.ini"
		return nil
	}
	env.Path = path
	return nil
}

func (env *Env) GetPath() string {
	return env.Path
}
