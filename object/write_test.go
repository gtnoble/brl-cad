package object

import "testing"

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
		name:       &name,
		attributes: map[string]string{"key": "value"},
		body:       []byte{0},
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

func TestMakeAttributes(t *testing.T) {
	name := "test"
	testObj := DbObject{
		name:       &name,
		attributes: map[string]string{"key": "value"},
		body:       []byte{0},
	}

	expectedSerializedAttributes := "key\x00value\x00"
	actualSerializedAttrs := testObj.serializeAttributes()
	if actualSerializedAttrs != expectedSerializedAttributes {
		t.Fatalf("Actual %q, Expected %q", actualSerializedAttrs, expectedSerializedAttributes)
	}
}
