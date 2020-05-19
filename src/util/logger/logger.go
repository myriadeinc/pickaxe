package LoggerUtil

import (
	"os"
	"fmt"
	"bytes"
	"net/http"
	"github.com/bhoriuchi/go-bunyan/bunyan"
)

var Logger *bunyan.Logger
var url string = "http://localhost:8080"

// Put on-hold for now since we can login to parse logs
type NetWorkLogger struct{}
func (n NetWorkLogger) Write(b []byte) (int, error) {
	_ , err:= http.Post(url, "application/json", bytes.NewBuffer(b))
	if err != nil {
		fmt.Println(err)
	}
	return 1, nil
}


func Init() () {
	staticFields := make(map[string]interface{})
	// var netlog NetWorkLogger
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
			// {
			// 	Name: "INFO",
			// 	Level: bunyan.LogLevelInfo,
			// 	Stream: netlog,
			// },
			// {
			// 	Name: "ERROR",
			// 	Level: bunyan.LogLevelError,
			// 	Stream: netlog,
			// },	
		},
		StaticFields: staticFields,
	}	
	log, err := bunyan.CreateLogger(config) 
	if err != nil {
		fmt.Println(err)
	}
	Logger = &log
}