package cocas

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/ESNFranceG33kTeam/coCAS/helpers"
)

func TestMain(m *testing.M) {
	setUp()
	//log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	//log.Println("Do stuff AFTER the tests!")
	//tearDown()
	os.Exit(exitVal)
}

func setUp() {
	helpers.TheAppConfig().ConfPath = "../test/conf.yaml"
	helpers.ReadConfig()
	InitCas()
}

func TestInitCas(t *testing.T) {
	InitCas()
}

func TestExternalLoadConf(t *testing.T) {
	ExternalLoadConf()
}

func TestGetCasServer(t *testing.T) {
	var cas_server interface{} = GetCasServer()

	_, ok := cas_server.(*http.Server)

	if ok == false {
		log.Fatalln("cas_server is not an http.server !")
	}
}
