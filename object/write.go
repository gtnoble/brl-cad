package object

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"
)

const NULL_CHARACTER = "\x00"

const (
	MAGIC_START = 0x76
	MAGIC_END   = 0x35
)

func (db DbObject) makeHFlags() byte {
	var flags uint8 = 0b11011000
	if db.name != nil {
		flags = flags | 0x1<<5
	}
	flags |= db.dli
	if db.isHidden {
		flags |= HIDDEN_OBJECT_FLAG
	}
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

func (db DbObject) serializeAttributes() string {
	attributesAlternated := make([]string, len(db.attributes)*2+1)
	attributesAlternated[len(attributesAlternated)-1] = ""
	var i int
	for key, value := range db.attributes {
		attributesAlternated[i] = key
		attributesAlternated[i+1] = value
		i += 2
	}

	return strings.Join(attributesAlternated, NULL_CHARACTER)
}

func (db DbObject) Write(w io.Writer) (int, error) {
	objectLength, _ := db.tryWrite(io.Discard, 0)
	return db.tryWrite(w, objectLength)
}

func (db DbObject) tryWrite(w io.Writer, objectLength int) (int, error) {

	type writeOp struct {
		condition bool
		operation func() (int, error)
	}

	var writeCount int
	writeOperations := []writeOp{
		{
			true,
			func() (int, error) {
				return w.Write([]byte{MAGIC_START, db.makeHFlags(), db.makeAFlags(), db.makeBFlags(), db.majorType, db.minorType})
			},
		},
		{
			true,
			func() (int, error) { return writeInt(w, objectLength/8) },
		},
		{
			db.name != nil,
			func() (int, error) { return writeDbString(w, *db.name) },
		},
		{
			db.attributes != nil,
			func() (int, error) { return writeDbString(w, db.serializeAttributes()) },
		},
		{
			db.body != nil,
			func() (int, error) { return writeInt(w, len(db.body)) },
		},
		{
			db.body != nil,
			func() (int, error) { return w.Write(db.body) },
		},
		{
			true,
			func() (int, error) {
				unpaddedFinalLength := writeCount + 1
				paddingNeeded := 8 - (unpaddedFinalLength % 8)
				return w.Write(make([]byte, paddingNeeded))
			},
		},
		{
			true,
			func() (int, error) { return w.Write([]byte{MAGIC_END}) },
		},
	}

	for _, operation := range writeOperations {
		if operation.condition {
			n, err := operation.operation()
			writeCount += n
			if err != nil {
				return writeCount, err
			}
		}
	}

	return writeCount, nil
}

func writeInt(w io.Writer, length int) (int, error) {
	err := binary.Write(w, binary.BigEndian, uint64(length))
	return 8, err
}

func writeDbString(w io.Writer, str string) (int, error) {

	nullTermStr := fmt.Sprintf("%s\x00", str)
	var writeCount int
	if n, err := writeInt(w, len(nullTermStr)); err != nil {
		return writeCount, err
	} else {
		writeCount += n
	}

	if n, err := w.Write([]byte(nullTermStr)); err != nil {
		return writeCount, err
	} else {
		writeCount += n
	}
	return writeCount, nil
}
