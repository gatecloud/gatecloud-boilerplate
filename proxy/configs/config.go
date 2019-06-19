package configs

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	libConfig "github.com/gatecloud/webservice-library/config"
)

// LocalConfig stores the configuration of the local environment variable
type LocalConfig struct {
	libConfig.Config
	Port          string `json:"port"`
	DbEngine      string `json:"db_engine"`
	DbConn        string `json:"db_conn"`
	Production    bool   `json:"production"`
	AutoMigration bool   `json:"auto_migration"`
	CorsEnabled   bool   `json:"cors_enabled"`
}

// Configuration is proxy's global configuration
var Configuration LocalConfig

func init() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dir = filepath.FromSlash(dir)
	rootDir := dir[:strings.LastIndex(dir, string(os.PathSeparator)+"proxy")]
	if err := Configuration.SetPath(rootDir + "/proxy/configs/local.ini"); err != nil {
		log.Fatal(err)
	}

	if err := Configuration.SetPrefix("proxy"); err != nil {
		log.Fatal(err)
	}

	if err := libConfig.LoadConfig(&Configuration); err != nil {
		log.Fatal(err)
	}

}
