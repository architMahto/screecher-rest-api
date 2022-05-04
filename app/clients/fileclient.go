package clients

import (
	"io/fs"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"github.com/gertd/go-pluralize"
	"github.com/gocarina/gocsv"
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

type FileClient interface {
	GetFilename(dest interface{}) string
	ReadFileContents(dest interface{}, reader Reader) error
	UpdateFileContents(dest interface{}, writer Writer) error
}

type FileDBClient struct {
	PathToDataDir string
}

func (fileDBClient FileDBClient) GetFilename(dest interface{}) string {
	tableType := strings.Split(reflect.TypeOf(dest).String(), ".")[1]
	pluralize := pluralize.NewClient()
	tableName := pluralize.Plural(strings.ToLower(tableType))
	filename := fileDBClient.PathToDataDir + "/" + tableName + ".csv"

	return filename
}

func (fileDBClient FileDBClient) ReadFileContents(
	dest interface{},
	reader Reader,
) error {
	filePath := fileDBClient.GetFilename(dest)

	data, readFileErr := reader.ReadFile(filePath)

	if readFileErr != nil {
		return readFileErr
	}

	unmarshalErr := gocsv.UnmarshalString(string(data), dest)

	if unmarshalErr != nil {
		return unmarshalErr
	}

	return nil
}

func (fileDBClient FileDBClient) UpdateFileContents(
	dest interface{},
	writer Writer,
) error {
	filePath := fileDBClient.GetFilename(dest)
	dataStr, marshalErr := gocsv.MarshalString(dest)

	if marshalErr != nil {
		return marshalErr
	}

	writeFileErr := writer.WriteFile(
		filePath,
		[]byte(dataStr),
		fs.FileMode(os.O_WRONLY|os.O_CREATE),
	)

	if writeFileErr != nil {
		return writeFileErr
	}

	return nil
}
