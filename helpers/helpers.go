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
	G33kTeam   string   `yaml:"geekteam_url"`
	ExtraAdmin []string `yaml:"extra_admin"`
}

var AppConfig Cfg

// flags
var Confpathflag string
var Swaggerpathflag string

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
