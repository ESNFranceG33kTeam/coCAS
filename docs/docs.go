package docs

import (
	_ "embed"

	"github.com/ESNFranceG33kTeam/coCAS/logger"
)

//go:embed VERSION.txt
var version_file string

const draw = `
                #####     #     #####
  ####   ####  #     #   # #   #     #
 #    # #    # #        #   #  #
 #      #    # #       #     #  #####
 #      #    # #       #######       #
 #    # #    # #     # #     # #     #
  ####   ####   #####  #     #  #####

by the French ESN G33kTeam.
`

func DrawStart() {
	logger.GetLogger().LogDraw(draw)
	logger.GetLogger().LogDraw("Version : " + version_file + "\n")
}

func GetVersion() string {
	return version_file
}
