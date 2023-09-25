package ivavigofork

import (
	"encoding/json"
	"os"
	"strings"
)


var defaultConfig = strings.Trim(`
{
    "log_file": "",
    "say_errors": true,
    "say_warnings": false,
    "say_if_hearing": true
}`, "\n")

const configPath = "./config.json"


type config struct {
	LogFile string
	SayErrors bool
	SayWarnings bool
	SayIfHearing bool
}


func init_config() config {
	con := config{}
	file, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
		//panic("Something went wrong while file creating")
	}
	
	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	
	if len(data) == 0 {
		file.Write([]byte(defaultConfig))
		//MyLogger.Warn("Config file is empty, default wrotten")
		// TODO: !!!think about logger there!!!
	}
	// TODO optimize this

	defer file.Close()

	rawFileData, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(rawFileData, &con)

	return con
}

var Config = init_config()
