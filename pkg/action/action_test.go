package action

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestPerformCopy(t *testing.T) {
	srcFile := mkFile(t, "srcFileName")
	destFile := mkFile(t, "destFileName")
	defer cleanUp(srcFile, destFile)

	writeTestDataToFile(t, srcFile)

	PerformCopy(srcFile.Name(), destFile.Name())

	srcBytes, err := ioutil.ReadAll(srcFile)
	if err != nil {
		t.Fatal("src file could not be read")
	}
	destBytes, err := ioutil.ReadAll(destFile)
	if err != nil {
		t.Fatal("dest file could not be read")
	}

	if !bytes.Equal(srcBytes, destBytes) {
		t.Fatalf("src does not equal dest")
	}

}

func cleanUp(srcFile *os.File, destFile *os.File) {
	os.Remove(srcFile.Name())
	os.Remove(destFile.Name())
}

func writeTestDataToFile(t *testing.T, f *os.File) {
	br := []byte{76, 79, 76}
	_, err := f.Write(br)
	if err != nil {
		t.Fatalf("failed to write to src file: %v", err)
	}
}

func mkFile(t *testing.T, filename string) *os.File {
	f, err := ioutil.TempFile("", filename)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	return f
}
