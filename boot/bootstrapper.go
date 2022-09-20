package boot

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/safziy/go-boot/config"
	"github.com/safziy/go-boot/web"
	"os"
	"path/filepath"
	"strings"
)

var (
	confPath  = flag.String("config", "./configs/application.yaml", "config path")
	conf *config.BootConfig
)

type Bootstrapper struct {

}

func NewBootApplication() *Bootstrapper {
	return &Bootstrapper{}
}

func (boot *Bootstrapper) Run() {
	flag.Parse()

	boot.initConfig(*confPath)

	web.InitWeb(conf.Web)

	fmt.Println("test")
}

func (boot *Bootstrapper) initConfig(path string) {
	var err error
	if len(path) < 1 {
		conf, err = config.InitConfig(boot.getRootPath())
	} else {
		conf, err = config.InitConfigWithFullPath(path)
	}
	if err != nil {
		fmt.Println("shutdown " + err.Error())
		boot.shutdown()
	}
	marshal, err := json.Marshal(conf)
	if err != nil {
		fmt.Println("shutdown " + err.Error())
		boot.shutdown()
	}
	fmt.Println(string(marshal))
}

func (boot *Bootstrapper) getRootPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		//boot.shuttingDownAbnormally("failed to get current directory, shutting down...")
	}
	lastIndexOfSlash := strings.LastIndex(dir, "/")
	rootPath := dir[:lastIndexOfSlash]
	return rootPath
}

func (boot *Bootstrapper) shutdown() {

}




