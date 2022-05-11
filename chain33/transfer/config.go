package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/kardianos/osext"
)

type Config struct {
	Txnum          int `json:",default=250000"`
	GrpcTxNum      int `json:",default=400"`
	Paras          Paras
	UserAddress    string `json:",default=1MowztiYG22fzqZEmf9PnwCwpEcLqmmjMz"`
	UserPrivateKey string `json:",default=0x1813e88e2ec3ae44dea06227a118d05ca2b7ba1d90e267d844a60538a6d48fbc"`
	ToAddr         string `json:",default=133AfuMYQXRxc45JGUb1jLk1M1W4ka39L1"`
	Amount         int64  `json:",default=10"`
}

type Paras []*Para

var Path string
var cfgPath = flag.String("f", "", "configfile")

// Setup initialize the configuration instance
func Setup() {
	Path = initPath()
	Cfg = initConfig()
}

func initConfig() *Config {
	var conf = Config{}
	flag.Parse()
	if *cfgPath == "" {
		*cfgPath = Path + "/config.toml"
	}
	if _, err := toml.DecodeFile(*cfgPath, &conf); err != nil {
		panic(err)
	}

	return &conf
}

func initPath() string {
	path, err := osext.ExecutableFolder()
	if err != nil {
		panic(err)
	}
	return path
}
