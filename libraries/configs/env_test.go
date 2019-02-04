package configs

import (
	"log"
	"testing"
)

func TestEnv(t *testing.T) {
	var (
		config   *Env
		dbEngine = "postgres"
	)

	config = &Env{}

	if err := config.SetPath("env.ini"); err != nil {
		log.Fatal(err)
	}

	if err := config.SetPrefix("env"); err != nil {
		log.Fatal(err)
	}

	if err := LoadConfig(config); err != nil {
		t.Errorf("expect=%s\nactual=%s", dbEngine, config.DbConn)
	}
}
