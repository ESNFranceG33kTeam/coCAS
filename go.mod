module github.com/ESNFranceG33kTeam/coCAS

go 1.18

require (
	github.com/stretchr/testify v1.8.0
	gopkg.in/cas.v2 v2.2.2
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/kr/text v0.2.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Packages system

replace github.com/ESNFranceG33kTeam/coCAS/cocas => ../cocas

replace github.com/ESNFranceG33kTeam/coCAS/helpers => ../helpers

replace github.com/ESNFranceG33kTeam/coCAS/logger => ../logger

replace github.com/ESNFranceG33kTeam/coCAS/docs => ../docs

replace github.com/ESNFranceG33kTeam/coCAS/modules/health => ../health
