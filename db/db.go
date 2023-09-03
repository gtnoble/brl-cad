package db

import (
	"brl-cad/object"
	"io"
)

func writeDb(w io.Writer, dataObjects ...object.DbObject) (int, error) {
	objects := []object.DbObject{object.MakeHeader()}
	objects = append(objects, dataObjects...)
	var bytesWritten int

	for _, object := range objects {
		n, err := object.Write(w)
		bytesWritten += n
		if err != nil {
			return bytesWritten, err
		}
	}

	return bytesWritten, nil
}
