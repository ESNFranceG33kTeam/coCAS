package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/ESNFranceG33kTeam/coCAS/cocas"
	"github.com/ESNFranceG33kTeam/coCAS/helpers"
	"github.com/ESNFranceG33kTeam/coCAS/logger"
)

func startoptions() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln("Abs path doesn't exist !")
	}

	flag.StringVar(
		&helpers.Confpathflag,
		"conf",
		filepath.Join(path, "test/conf_local.yaml"),
		"path for the configuration file.",
	)

	flag.Parse()
}

func InitConf() {
	helpers.InitFile()
	helpers.ReadConfig()
	logger.GetLogger().LogInit()
}

func main() {
	startoptions()
	InitConf()
	cocas.InitCas()

	logger.GetLogger().LogInfo("main", "coCAS ready.")

	// create a WaitGroup
	wg := new(sync.WaitGroup)
	// add two goroutines to `wg` WaitGroup
	wg.Add(2)

	go func() {
		logger.GetLogger().
			LogCritical("main", "listen cas error", cocas.GetCasServer().ListenAndServe())
		wg.Done()
	}()

	// wait until WaitGroup is done
	wg.Wait()
}
