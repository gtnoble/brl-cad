package db

import (
	"brl-cad/object"
	"io"
)

const (
	MAGIC_START = 0x76
	MAGIC_END   = 0x35
)

func writeDb(w io.Writer, dataObjects ...object.DbObject) (int, error) {
	objects := []object.DbObject{object.MakeHeader()}
	objects = append(objects, dataObjects...)
	var bytesWritten int

	n, err := w.Write([]byte{MAGIC_START})
	bytesWritten += n
	if err != nil {
		return bytesWritten, err
	}

	for _, object := range objects {
		n, err := object.Write(w)
		bytesWritten += n
		if err != nil {
			return bytesWritten, err
		}
	}

	n, err = w.Write([]byte{MAGIC_END})
	bytesWritten += n
	if err != nil {
		return bytesWritten, err
	}
	return bytesWritten, nil
}
