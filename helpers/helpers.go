package helpers

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Cfg struct {
	ConfPath   string
	PortCas    string   `yaml:"portcas"`
	Loglevel   int      `yaml:"loglevel"`
	LocalAsso  string   `yaml:"local_url"`
	GlobalUrl  string   `yaml:"global_url"`
	AssoId     string   `yaml:"asso_id"`
	ExtraAdmin []string `yaml:"extra_admin"`
}

var AppConfig Cfg

// flags
var Confpathflag string

func InitFile() {
	AppConfig.ConfPath = Confpathflag
}

func ReadConfig() {
	f, err := os.Open(AppConfig.ConfPath)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		fmt.Println(err)
	}
}

func TheAppConfig() *Cfg {
	return &AppConfig
}
