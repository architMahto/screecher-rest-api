package clients_test

import (
	"os"
	"testing"

	"github.com/architMahto/screecher-rest-api/app/clients"
	"github.com/architMahto/screecher-rest-api/domain"
)

type MockReader struct {
	MockFileReader clients.Reader
}

func NewMockReader() MockReader {
	return MockReader{
		MockFileReader: clients.FileReader{},
	}
}

type FakeFileReader struct{}

func (fakeFileReader FakeFileReader) ReadFile(filename string) (
	[]byte,
	error,
) {
	readFileResultStr := "id,first_name,last_name,username,password,secret_token,profile_image,date_created,date_modified\n1,Candie,Splevings,csplevings0,JPGheabcNvCA,s97sHjkF244pZzZPcI3AloxgMBPzAU00,http://dummyimage.com/100x140.png/cc0000/ffffff,2021-10-31T09:53:39Z,2021-11-14T09:53:39Z\n2,Gertie,Escritt,gescritt1,7Pr8wSkI9f,8VHxxSkcJ5sSsB0P5uKp6x3M7gmtSKiw,http://dummyimage.com/164x187.png/5fa2dd/ffffff,2021-05-02T07:55:35Z,2021-05-13T07:55:35Z"
	return []byte(readFileResultStr), nil
}

func TestReadCsvContents(t *testing.T) {
	testMockReader := MockReader{
		MockFileReader: FakeFileReader{},
	}
	csvDbClient := clients.CsvDbClient{
		PathToDataDir: "../../data",
	}
	users := []domain.User{}
	err := csvDbClient.ReadCsvContents(&users, testMockReader.MockFileReader)

	if err != nil {
		t.Errorf("An error take place when running this")
	}
}

type MockWriter struct {
	MockFileWriter clients.Writer
}

func NewMockWriter() MockWriter {
	return MockWriter{
		MockFileWriter: clients.FileWriter{},
	}
}

type FakeFileWriter struct{}

func (fakeFileWriter FakeFileWriter) WriteFile(
	filename string,
	data []byte,
	perm os.FileMode,
) error {
	return nil
}

func TestUpdateCsvContents(t *testing.T) {
	testMockWriter := MockWriter{
		MockFileWriter: FakeFileWriter{},
	}
	csvDbClient := clients.CsvDbClient{
		PathToDataDir: "../../data",
	}
	users := []domain.User{}
	err := csvDbClient.UpdateCsvContents(&users, testMockWriter.MockFileWriter)

	if err != nil {
		t.Errorf("An error take place when running this")
	}
}
