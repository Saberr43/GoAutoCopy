package main

import (
	"path/filepath"

	"github.com/Saberr43/GoAutoCopy/configs"
)

func main() {
	configFilePath, err := filepath.Abs("config.xml")
	check(err)

	xml, err := configs.GetActions(configFilePath)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
