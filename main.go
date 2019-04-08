package main

import (
	"log"
	"os"

	"github.com/fsnotify/fsnotify"

	"github.com/Saberr43/GoAutoCopy/pkg/action"
	"github.com/Saberr43/GoAutoCopy/pkg/configs"
)

func main() {
	configName := "config.xml"

	watcher, err := fsnotify.NewWatcher()
	check(err)

	defer watcher.Close()

	config, err := os.Open(configName)
	check(err)

	defer config.Close()

	configObj, err := configs.GetConfigs(config)
	check(err)

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}

				act := configObj.GetActionBySource(event.Name)
				if act.IsValidFileType(event.Name) {
					err = action.PerformCopy(event.Name, act.Destination)
					check(err)
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(config.Name())
	check(err)

	for _, action := range configObj.Actions {
		err = watcher.Add(action.Source)
		check(err)
	}

	<-done
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
