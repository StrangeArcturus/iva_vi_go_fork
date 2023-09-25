package ivavigofork

import (
	"encoding/json"
	//"fmt"
	//"io"
	"os"
	"strings"
	//"io/ioutil"
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


func (self *config) init() config {
	file, err := os.Open(configPath)
	if err != nil {
		file, err := os.Create(configPath)
		if err != nil {
			panic("Something went wrong while file creating")
		}
		file.Write([]byte(defaultConfig))
	}
	defer file.Close()

	rawFileData, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	//fileData := string(rawFileData)

	//var configObject map[string]string = make(map[string]string)

	json.Unmarshal(rawFileData, &self)

	//fmt.Println(`log_file: {log_file}`)

	return *self
}

var Config = func () config {
	con := &config{}
	con.init()
	return *con
}()
