package index

import (
	"archive/zip"
	"bytes"
	"database/sql"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/dunhamsteve/iwork/proto/TSP"

	"github.com/golang/protobuf/proto"
	"github.com/golang/snappy"

	// register sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

// Index holds the content of an iwork file
type Index struct {
	Type    string                 `json:"type"`
	Records map[uint64]interface{} `json:"records"`
}

// Open loads a document into an Index structure
func Open(doc string) (*Index, error) {
	indexType := strings.TrimSuffix(filepath.Ext(doc)[1:], "-tef")
	fn := path.Join(doc, "Index.zip")
	zf, err := zip.OpenReader(fn)
	if err != nil {
		// iWork 5.5
		zf, err = zip.OpenReader(doc)
	}
	if err == nil {
		defer zf.Close()
		ix := &Index{indexType, nil}
		err = ix.loadZip(zf)

		return ix, err
	}

	// .pages-tef files, sqlite
	fn = path.Join(doc, "index.db")
	_, err = os.Stat(fn)
	if err == nil {
		db, err := sql.Open("sqlite3", fn)
		if err == nil {
			defer db.Close()
			ix := &Index{indexType, nil}
			err = ix.loadSQL(db)
			return ix, err
		}
	}

	return nil, err
}

func (ix *Index) loadSQL(db *sql.DB) error {
	ix.Records = make(map[uint64]interface{})
	stmt := `select o.identifier, o.class, ds.state from objects o join dataStates ds on o.state = ds.identifier`
	rows, err := db.Query(stmt)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var id uint64
		var class uint32
		var data []byte
		err = rows.Scan(&id, &class, &data)
		if err != nil {
			return err
		}
		ix.decodePayload(id, class, data)
	}
	return nil
}

func (ix *Index) loadZip(zf *zip.ReadCloser) error {
	ix.Records = make(map[uint64]interface{})
	for _, f := range zf.File {
		if strings.HasSuffix(f.Name, ".iwa") {
			rc, err := f.Open()
			if err != nil {
				return err
			}
			defer rc.Close()

			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return err
			}
			err = ix.loadIWA(data)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Deref returns the object pointed to by a TSP.Reference
func (ix *Index) Deref(ref *TSP.Reference) interface{} {
	if ref == nil {
		return nil
	}
	return ix.Records[*ref.Identifier]
}

func (ix *Index) loadIWA(data []byte) error {
	data, err := unsnap(data)
	if err != nil {
		return err
	}

	r := bytes.NewBuffer(data)
	for {
		l, err := binary.ReadUvarint(r)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		chunk := make([]byte, l)
		_, err = r.Read(chunk)
		if err != nil {
			return err
		}
		var ai TSP.ArchiveInfo
		err = proto.Unmarshal(chunk, &ai)
		if err != nil {
			return err
		}

		for _, info := range ai.MessageInfos {
			payload := make([]byte, *info.Length)
			_, err := r.Read(payload)
			if err != nil {
				return err
			}

			id, typ := *ai.Identifier, *info.Type

			ix.decodePayload(id, typ, payload)
		}
	}
	return nil
}

func (ix *Index) decodePayload(id uint64, typ uint32, payload []byte) {
	var value interface{}
	var err error
	if ix.Type == "pages" {
		value, err = decodePages(typ, payload)
	} else if ix.Type == "numbers" {
		value, err = decodeNumbers(typ, payload)
	} else if ix.Type == "key" {
		value, err = decodeKeynote(typ, payload)
	} else {
		fmt.Println("Cannot decode files of type", ix.Type)
	}

	if err != nil {
		// These we don't care as much about
		fmt.Println("ERR", id, typ, err)
		return
	}

	ix.Records[id] = value
}

func unsnap(data []byte) ([]byte, error) {
	rval := bytes.NewBuffer(nil)
	for len(data) > 0 {
		typ := int(data[0])
		if typ != 0 {
			return nil, errors.New("snap header type not 0")
		}
		l := int(data[1]) | int(data[2])<<8 | int(data[3])<<16
		tmp, err := snappy.Decode(nil, data[4:4+l])
		if err != nil {
			return nil, err
		}
		rval.Write(tmp)
		data = data[4+l:]
	}
	return rval.Bytes(), nil
}
