package cocas

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/ESNFranceG33kTeam/coCAS/helpers"
	"github.com/ESNFranceG33kTeam/coCAS/logger"
	"gopkg.in/cas.v2"
	"gopkg.in/yaml.v2"
)

var cocas *http.Server
var cocasmwd *http.ServeMux

type myHandler struct{}

var MyHandler = &myHandler{}

type GalaxyConf struct {
	GlobalConfUrl string
	LocalConfUrl  string
	CASServer     string   `yaml:"cas_server"`
	Country       string   `yaml:"country"`
	GlobalRoles   []string `yaml:"global_roles"`
	LocalRoles    []string `yaml:"local_roles"`
	G33kTeam      []string `yaml:"geek_username"`
	AssoId        string   `yaml:"asso_id"`
	ExtraUsername []string `yaml:"extra_username"`
}

var GalaxyUsers GalaxyConf

func InitCas() {

	// Conf
	GalaxyUsers.ExtraUsername = helpers.TheAppConfig().ExtraAdmin[:]
	GalaxyUsers.AssoId = helpers.TheAppConfig().AssoId

	// Global conf
	GalaxyUsers.GlobalConfUrl = helpers.TheAppConfig().GlobalUrl
	global_conf, err := http.Get(GalaxyUsers.GlobalConfUrl)
	if err != nil {
		logger.GetLogger().LogError("cocas", "global conf url file unavailable", err)
	}
	defer global_conf.Body.Close()
	decoder := yaml.NewDecoder(global_conf.Body)
	err = decoder.Decode(&GalaxyUsers)
	if err != nil {
		fmt.Println(err)
	}

	// Local conf
	if helpers.TheAppConfig().LocalAsso != "" {
		GalaxyUsers.LocalConfUrl = helpers.TheAppConfig().LocalAsso
		local_conf, err := http.Get(GalaxyUsers.GlobalConfUrl)
		if err != nil {
			logger.GetLogger().LogError("cocas", "local conf url file unavailable", err)
		}
		defer local_conf.Body.Close()
		decoder = yaml.NewDecoder(local_conf.Body)
		err = decoder.Decode(&GalaxyUsers)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Server init
	cocasmwd = http.NewServeMux()
	cocasmwd.Handle("/", MyHandler)

	url, _ := url.Parse(GalaxyUsers.CASServer)
	client := cas.NewClient(&cas.Options{
		URL: url,
	})

	cocas = &http.Server{
		Addr:    ":" + helpers.TheAppConfig().PortCas,
		Handler: client.Handle(cocasmwd),
	}

	logger.GetLogger().LogInfo("cocas", "Conf up !")
}

func GetCasServer() *http.Server {
	return cocas
}
