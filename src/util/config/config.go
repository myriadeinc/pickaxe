package ConfigUtil

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Service__name								string 	`env:"service__name"`
	Service__host								string	`env:"service__host"`
	Service__shared_secret			string	`env:"service__shared_secret"`
}

var Config 	*Configuration

func Init() {
	Config = &Configuration{}

	err := gonfig.GetConf(getFileName(), Config)
	if err != nil {
		fmt.Println(err)
		os.Exit(500)
	}
}

func getFileName() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "local"
	}
	filename := []string{"../../../config/", "config.", env, ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))
	return filePath
}