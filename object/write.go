package object

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"
)

const NULL_CHARACTER = "\x00"

type DbObject struct {
	dli             byte
	majorType       byte
	minorType       byte
	compressionFlag byte
	name            *string
	attributes      map[string]string
	body            []byte
}

func (db DbObject) makeHFlags() byte {
	var flags uint8 = 0b11011000
	if db.name != nil {
		flags = flags | 0x1<<5
	}
	flags = flags | db.dli
	return flags
}

func (db DbObject) makeAFlags() byte {
	return db.makeABFlags(db.attributes != nil)
}

func (db DbObject) makeBFlags() byte {
	return db.makeABFlags(db.body != nil)
}

func (db DbObject) makeABFlags(hasContent bool) byte {
	var flags byte = 0b11000000
	if hasContent {
		flags = flags | 0x1<<5
	}
	flags = flags | db.compressionFlag
	return flags
}

func makeAttributes(attributes map[string]string) string {
	attributesAlternated := make([]string, len(attributes)*2+1)
	attributesAlternated[len(attributesAlternated)-1] = NULL_CHARACTER
	var i int
	for key, value := range attributes {
		attributesAlternated[i] = key
		attributesAlternated[i+1] = value
		i += 2
	}

	return strings.Join(attributesAlternated, NULL_CHARACTER)
}

func writeLength(w io.Writer, length int) (int, error) {
	err := binary.Write(w, binary.BigEndian, int64(length))
	return 8, err
}

func writeDbString(w io.Writer, str string) (int, error) {

	var writeCount int
	if n, err := writeLength(w, len(str)); err != nil {
		return writeCount, err
	} else {
		writeCount += n
	}

	if n, err := fmt.Fprintf(w, "%s\x00", str); err != nil {
		return writeCount, err
	} else {
		writeCount += n
	}
	return writeCount, nil
}

func (db DbObject) Write(w io.Writer) (int, error) {

	var writeCount int

	for _, makeFlags := range []func() byte{db.makeHFlags, db.makeAFlags, db.makeBFlags} {
		n, err := w.Write([]byte{makeFlags()})
		writeCount += n
		if err != nil {
			return writeCount, err
		}
	}

	if db.name != nil {
		n, err := writeDbString(w, *db.name)
		writeCount += n
		if err != nil {
			return writeCount, err
		}
	}

	if db.attributes != nil {
		n, err := writeDbString(w, makeAttributes(db.attributes))
		writeCount += n
		if err != nil {
			return writeCount, err
		}
	}

	if db.body != nil {
		n, err := writeLength(w, len(db.body))
		writeCount += n
		if err != nil {
			return writeCount, err
		}

		n, err = w.Write(db.body)
		writeCount += n
		if err != nil {
			return writeCount, err
		}
	}

	return writeCount, nil
}
