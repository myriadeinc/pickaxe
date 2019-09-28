package Logger

import (
	"os"
	"fmt"
	"github.com/bhoriuchi/go-bunyan/bunyan"
)

var Logger *bunyan.Logger

func Init() () {
	staticFields := make(map[string]interface{})
	var config bunyan.Config = bunyan.Config{
		Name: "PickAxe",
		Streams: []bunyan.Stream{
			{
					Name: "INFO",
					Level: bunyan.LogLevelInfo,
					Stream: os.Stdout,
			},
			{
					Name: "ERROR",
					Level: bunyan.LogLevelError,
					Stream: os.Stderr,
			},
		},
		StaticFields: staticFields,
	}	
	log, err := bunyan.CreateLogger(config) 
	if err != nil {
		fmt.Println(err)
	}
	Logger = &log
}