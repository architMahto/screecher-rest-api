package clients

import (
	"io/fs"
	"os"
	"reflect"
	"strings"

	"github.com/gertd/go-pluralize"
	"github.com/gocarina/gocsv"
)

type CsvClient interface {
	GetFilename(dest interface{}) string
	ReadCsvContents(dest interface{}, reader Reader) error
	UpdateCsvContents(dest interface{}, writer Writer) error
}

type CsvDbClient struct {
	PathToDataDir string
}

func (csvDbClient CsvDbClient) GetFilename(dest interface{}) string {
	tableType := strings.Split(reflect.TypeOf(dest).String(), ".")[1]
	pluralize := pluralize.NewClient()
	tableName := pluralize.Plural(strings.ToLower(tableType))
	filename := csvDbClient.PathToDataDir + "/" + tableName + ".csv"

	return filename
}

func (csvDbClient CsvDbClient) ReadCsvContents(
	dest interface{},
	reader Reader,
) error {
	filePath := csvDbClient.GetFilename(dest)

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

func (csvDbClient CsvDbClient) UpdateCsvContents(
	dest interface{},
	writer Writer,
) error {
	filePath := csvDbClient.GetFilename(dest)
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
