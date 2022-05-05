package clients

import (
	"io/ioutil"
	"os"
)

type Reader interface {
	ReadFile(filename string) ([]byte, error)
}

type FileReader struct{}

func (reader FileReader) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

type Writer interface {
	WriteFile(filename string, data []byte, perm os.FileMode) error
}

type FileWriter struct{}

func (writer FileWriter) WriteFile(
	filename string,
	data []byte,
	perm os.FileMode,
) error {
	return ioutil.WriteFile(filename, data, perm)
}
