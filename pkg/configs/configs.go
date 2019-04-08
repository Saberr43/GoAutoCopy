package configs

import (
	"encoding/xml"
	"io"
	"path/filepath"
	"strings"
)

type Config struct {
	NodeName xml.Name `xml:"config"`
	Actions  []Action `xml:"action"`
}

//Action model of action node in cofig
type Action struct {
	NodeName    xml.Name `xml:"action"`
	Source      string   `xml:"source,attr"`
	Destination string   `xml:"destination,attr"`
	FileTypes   string   `xml:"filetypes,attr"`
}

func removeFilename(input string) string {
	b := `\` + filepath.Base(input)
	s := strings.Replace(input, b, "", -1)
	return s
}

//IsValidFileType checks to see if file path passed in is a file type defined in 'FileTypes' field
//If 'FileTypes' is empty return true
func (action *Action) IsValidFileType(src string) bool {
	if action.FileTypes == "" {
		return true
	}

	for _, extension := range strings.Split(action.FileTypes, ",") {
		if filepath.Ext(src) == "."+extension {
			return true
		}
	}

	return false
}

//GetActionBySource returns the action of destination given its source file path
func (cnf *Config) GetActionBySource(src string) Action {
	for _, act := range cnf.Actions {
		if act.Source == removeFilename(src) {
			return act
		}
	}

	return Action{}
}

//GetConfigs gets actions from Reader interface
func GetConfigs(rdr io.Reader) (Config, error) {
	var actions Config

	err := xml.NewDecoder(rdr).Decode(&actions)
	if err != nil {
		return actions, err
	}

	return actions, nil
}
