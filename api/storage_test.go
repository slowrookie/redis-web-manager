package api

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/suite"
)

// Reacord is the test struct
type Record struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type StorageSuite struct {
	suite.Suite
	storage *Storage
}

func (sut *StorageSuite) SetupTest() {
	const project = "com.github.slowrookie.redis-web-manager"

	var dir string

	switch runtime.GOOS {
	case "windows":
		dir = fmt.Sprintf("%s/%s", os.Getenv("APPDATA"), project)
	case "darwin":
		dir = fmt.Sprintf("%s/Library/Containers/%s", os.Getenv("HOME"), project)
	case "linux":
		dir = fmt.Sprintf("%s/.%s", os.Getenv("HOME"), project)
	}

	sut.storage = &Storage{dir: path.Join(dir, "storage")}
}

func (sut *StorageSuite) TestWrite() {
	err := sut.storage.Write("records", "one record", Record{Id: "123", Name: "slowrookie", Email: "liujiaxingemail@gmail.com"})
	sut.ErrorIs(err, nil)
	sut.storage.Delete("records", "one record")
}

func (sut *StorageSuite) TestRead() {
	err := sut.storage.Write("records", "one record", Record{Id: "123", Name: "slowrookie", Email: "liujiaxingemail@gmail.com"})
	sut.ErrorIs(err, nil)

	var record = Record{}
	err = sut.storage.Read("records", "one record", &record)
	sut.ErrorIs(err, nil)
	sut.Equal("123", record.Id)

	sut.storage.Delete("records", "one record")
}

func (sut *StorageSuite) TestReadAll() {
	nums := []int{1, 2, 3, 4, 5}
	for _, n := range nums {
		err := sut.storage.Write("records", fmt.Sprintf("record%d", n), Record{Id: fmt.Sprintf("%d", n), Name: "slowrookie", Email: "liujiaxingemail@gmail.com"})
		sut.ErrorIs(err, nil)
	}

	var byts [][]byte
	err := sut.storage.ReadAll("records", &byts)

	var records []Record
	// for _, bts := range byts {
	// 	var record = Record{}
	// 	json.Unmarshal(bts, &record)
	// 	records = append(records, record)
	// }

	RecordsToStruct(byts, &records)

	sut.ErrorIs(err, nil)
	sut.GreaterOrEqual(len(nums), len(records))
}

func TestStorageSuite(t *testing.T) {
	suite.Run(t, new(StorageSuite))
}
