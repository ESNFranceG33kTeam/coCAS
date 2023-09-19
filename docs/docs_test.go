package docs

import (
	"log"
	"testing"
)

func TestDrawStart(t *testing.T) {
	DrawStart()
}

func TestGetVersion(t *testing.T) {
	var version interface{} = GetVersion()

	_, ok := version.(string)

	if ok == false {
		log.Fatalln("version is not a string !")
	}

	if version == "" {
		log.Fatalln("version is empty !")
	}
}
