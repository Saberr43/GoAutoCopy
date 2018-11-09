package configs

import (
	"encoding/xml"
	"os"
)

//Actions model of actions node in cofig
type Actions struct {
	XMLName  xml.Name `xml:"actions"`
	Settings []Action `xml:"action"`
}

//Action model of action node in cofig
type Action struct {
	XMLName     xml.Name `xml:"action"`
	Source      string   `xml:"source,attr"`
	Destination string   `xml:"destination,attr"`
	FileTypes   string   `xml:"filetypes,attr"`
}

//GetActions gets actions from config.xml
func GetActions(absFilePath string) (Actions, error) {
	var actions Actions

	file, err := os.Open(absFilePath)
	if err != nil {
		return actions, err
	}

	defer file.Close()

	err = xml.NewDecoder(file).Decode(&actions)

	if err != nil {
		return actions, err
	}

	return actions, nil
}
