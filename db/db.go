package db

import (
	"brl-cad/object"
	"io"
)

const dbFileHeader = "\x76\x01\x00\x00\x00\x00\x01\x35"

func writeDb(w io.Writer, title string, unitConversion float64, dataObjects ...object.DbObject) (int, error) {
	var bytesWritten int

	globalObject := object.MakeGlobal(title, unitConversion)
	dataObjects = append([]object.DbObject{globalObject}, dataObjects...)

	n, err := w.Write([]byte(dbFileHeader))
	bytesWritten += n
	if err != nil {
		return bytesWritten, err
	}

	for _, object := range dataObjects {
		n, err := object.Write(w)
		bytesWritten += n
		if err != nil {
			return bytesWritten, err
		}
	}

	return bytesWritten, nil
}
