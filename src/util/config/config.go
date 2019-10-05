package ConfigUtil

import (
	"fmt"
	"os"
	"path"
	"strings"
	"path/filepath"
	"runtime"
	"github.com/spf13/viper"
)

func Init() {

	replacer := strings.NewReplacer(".", "__")
	viper.SetEnvKeyReplacer(replacer)

	viper.AutomaticEnv()
	viper.SetConfigName("local")
	viper.AddConfigPath("./config/")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file, %s", err)
	}
}

func Get(key string) (interface{}){
	return viper.Get(key)
}

func getFileName() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "local"
	}
	filename := []string{"../../../config/", "config.", env, ".yml"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))
	return filePath
}