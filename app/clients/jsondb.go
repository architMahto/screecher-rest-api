package clients

import (
	"encoding/json"
	"io/fs"
	"os"
)

type JsonClient interface {
	GetFilename() string
	ReadJsonContents(dest interface{}, reader Reader) error
	UpdateJsonContents(dest interface{}, writer Writer) error
}

type JsonDbClient struct {
	PathToDataDir string
}

func (jsonDbClient JsonDbClient) GetFilename() string {
	filename := jsonDbClient.PathToDataDir + "/tokens.json"

	return filename
}

func (jsonDbClient JsonDbClient) ReadJsonContents(
	dest interface{},
	reader Reader,
) error {
	filePath := jsonDbClient.GetFilename()

	data, readFileErr := reader.ReadFile(filePath)

	if readFileErr != nil {
		return readFileErr
	}

	unmarshalErr := json.Unmarshal(data, dest)

	if unmarshalErr != nil {
		return unmarshalErr
	}

	return nil
}

func (jsonDbClient JsonDbClient) UpdateJsonContents(
	dest interface{},
	writer Writer,
) error {
	filePath := jsonDbClient.GetFilename()
	dataStr, marshalErr := json.Marshal(dest)

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
