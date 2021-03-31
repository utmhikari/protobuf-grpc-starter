package db

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/utmhikari/protobuf-grpc-starter/internal/shared/models"
	"github.com/utmhikari/protobuf-grpc-starter/internal/shared/util"
	"github.com/utmhikari/protobuf-grpc-starter/internal/shared/util/fs"
	"log"
	"os"
)


const Sep = "\n"


func GetDBFileName() string {
	return fmt.Sprintf("%s/%s", util.GetTmpDir(), "docs.dat")
}


// touchDBFile return touched, syserr
func touchDBFile() (bool, error) {
	if err := util.MakeTmpDir(); err != nil {
		return false, err
	}

	fp := GetDBFileName()
	if fs.IsDirectory(fp) {
		return false, errors.New("db data file is a directory")
	}

	if !fs.ExistsPath(fp) {
		// touch it
		f, err := os.Create(fp)
		if err != nil {
			return false, err
		} else {
			defer func() {
				_ = f.Close()
			}()
			return true, nil
		}
	} else {
		fi, err := os.Stat(fp)
		if err != nil {
			return false, err
		}
		if fi.Size() == 0 {
			// empty file is same as first touched
			return true, nil
		} else {
			return false, nil
		}
	}
}


func GetAllDocs() (*[]models.Document, int, error) {
	_, err := touchDBFile()
	if err != nil {
		return nil, 0, err
	}

	b, err := os.ReadFile(GetDBFileName())
	if err != nil {
		return nil, 0, err
	}

	items := bytes.Split(b, []byte(Sep))
	data := make([]models.Document, len(items))
	cnt := 0
	for i := 0; i < len(items); i++ {  // last one is always empty
		var doc models.Document
		err = json.Unmarshal(items[i], &doc)
		if err != nil {
			log.Printf("unmarshal doc err: %s\n", err.Error())
		} else {
			data[cnt] = doc
			cnt++
		}
	}

	log.Printf("current documents <%d>: %+v\n", cnt, data)

	return &data, cnt, nil
}


func Save(doc *models.Document) error {
	if nil == doc {
		return errors.New("nil document")
	}

	b, err := json.Marshal(doc)
	if err != nil {
		return err
	}

	touched, err := touchDBFile()
	if err != nil {
		return err
	}
	if !touched {
		b = append([]byte(Sep), b...)
	}

	// append file
	f, err := os.OpenFile(GetDBFileName(), os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	_, err = f.Write(b)
	return err
}

