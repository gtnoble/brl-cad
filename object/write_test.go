package object

import (
	"bytes"
	"encoding/binary"
	"testing"
)

const (
	OBJECT_LENGTH_START = 6
	OBJECT_LENGTH_END   = 14
)

func makeTestFlag(expectedFlag byte, actualFlag byte) func(*testing.T) {
	return func(t *testing.T) {
		if actualFlag != expectedFlag {
			t.Fatalf("Generated flag is equal to %b, should be %b", actualFlag, expectedFlag)
		}
	}
}

func TestMakeFlags(t *testing.T) {

	emptyObj := DbObject{}

	name := "test"
	testObj := DbObject{
		Name:       &name,
		attributes: map[string]string{"key": "value"},
		body:       func() []byte { return []byte{0} },
	}

	t.Run("Testing HFlag with name", func(t *testing.T) {
		expectedFlag := byte(0b11111000)
		actualFlag := testObj.makeHFlags()
		makeTestFlag(expectedFlag, actualFlag)(t)
	})

	t.Run("Testing HFlag without name", func(t *testing.T) {
		expectedFlag := byte(0b11011000)
		actualFlag := emptyObj.makeHFlags()
		makeTestFlag(expectedFlag, actualFlag)
	})

	t.Run("Testing ABFlag with content", func(t *testing.T) {
		makeTestFlag(testObj.makeABFlags(true), 0b11100000)
	})

	t.Run("Testing ABFlag without content", func(t *testing.T) {
		makeTestFlag(testObj.makeABFlags(false), 0b11000000)
	})

}

func TestWriteInt(t *testing.T) {
	testVal := int(1)
	buf := bytes.Buffer{}
	expectedBytesWritten, err := writeInt(&buf, testVal)
	if err != nil {
		t.Fatalf("Failed to write to byte buffer: %v", err)
	}
	numBytesWritten := buf.Len()
	if numBytesWritten != 8 {
		t.Fatalf("We are expecting %d bytes to be written for an unit64, %d bytes actually written. written bytes: % x", expectedBytesWritten, numBytesWritten, buf.Bytes())
	}
	lowestOrderByte := buf.Bytes()[7]
	if lowestOrderByte != 1 {
		t.Fatalf("Lowest order byte has value %d, we are expecting value %d", lowestOrderByte, testVal)
	}
}

func TestWriteDbString(t *testing.T) {
	testString := "test string"
	numBytesLengthField := 8
	buf := bytes.Buffer{}

	n, err := writeDbString(&buf, testString)
	if err != nil {
		t.Fatalf("Failed to write to byte buffer: %v", err)
	}
	expectedLength := len(testString) + 1 + numBytesLengthField
	if n != expectedLength {
		t.Fatalf("Written serialized string has length %d, expected length %d. written bytes: % x", n, expectedLength, buf.Bytes())
	}

	nullTermTestString := "test string\x00"
	writtenString := string(buf.Bytes()[numBytesLengthField:])
	if writtenString != string(nullTermTestString) {
		t.Fatalf("Written serialized string has content %s, expected %s", writtenString, nullTermTestString)
	}
}

func TestSerializeAttributes(t *testing.T) {
	testObj := DbObject{
		attributes: map[string]string{"key": "value", "hello": "world"},
	}

	serializedAttrs := testObj.serializeAttributes()
	possibleSerialization1 := "key\x00value\x00hello\x00world\x00"
	possibleSerialization2 := "hello\x00world\x00key\x00value\x00"
	if !(serializedAttrs == possibleSerialization1 || serializedAttrs == possibleSerialization2) {
		t.Fatalf("Serialized attributes have value %s, expected %s or %s", serializedAttrs, possibleSerialization1, possibleSerialization2)
	}
}

func TestWrite(t *testing.T) {
	testObj := DbObject{}
	buf := bytes.Buffer{}
	n, err := testObj.Write(&buf)
	if err != nil {
		t.Fatalf("Failed to write to byte buffer: %v", err)
	}
	objectField := buf.Bytes()[OBJECT_LENGTH_START:OBJECT_LENGTH_END]
	objectLengthFieldValue := int(binary.BigEndian.Uint64(objectField))
	if objectLengthFieldValue != n/8 {
		t.Fatalf("Object length field has value %d, actual object length is %d", objectLengthFieldValue, n)
	}

}
