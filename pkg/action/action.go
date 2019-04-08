package action

import (
	"errors"
	"os"
	"path/filepath"
)

//PerformCopy calls copy from a reader to a writer
func PerformCopy(source string, dest string) error {
	if filepath.Ext(source) == "" {
		return errors.New("source file path lacks file extension")
	}

	if filepath.Ext(dest) == "" {
		//if dest does not end in `\`
		if dest[len(dest)-1:] != `\` {
			dest += `\` //append a `\`
		}

		dest += filepath.Base(source) //append name of file to dest
	}

	//remove file if it exists
	os.Remove(dest)

	return os.Link(source, dest)
}
