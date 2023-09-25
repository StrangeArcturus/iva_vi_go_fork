package ivavigofork


import (
	"os"
	"encoding/json"
)


var defaultOwner = `
{
    "name": "",
    "home_city": "Ufa",
    "native_language": "",
    "target_language": ""
}`


const ownerPath = "./owner.json"


type owner struct {
	Name string
	HomeCity string
	NativeLanguage string
	TargetLanguage string
}


func initOwner(own owner) owner {
	file, err := os.OpenFile(ownerPath, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	data, err := os.ReadFile(ownerPath)
	if err != nil {
		panic(err)
	}

	if len(data) == 0 {
		file.Write([]byte(defaultOwner))
		MyLogger.Warn("Owner file is empty, default wroten")
	}

	defer file.Close()

	rawFileData, err := os.ReadFile(ownerPath)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(rawFileData, &own)

	return own
}
