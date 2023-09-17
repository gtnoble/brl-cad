package object

import (
	"testing"
)

func TestSerializeScalar(t *testing.T) {
	testScalar := Scalar(1.0)
	serializedScalar := testScalar.serialize()
	if numSerialized := len(serializedScalar); numSerialized != SCALAR_SIZE {
		t.Fatalf("Number of serialized bytes is %d, expected %d", numSerialized, SCALAR_SIZE)
	}
	expectedLowOrderByte := 0
	if lowOrderByte := serializedScalar[SCALAR_SIZE-1]; lowOrderByte != byte(expectedLowOrderByte) {
		t.Fatalf("Low order byte of serialized scalar is %x, expected %x. serialized bytes: % x", lowOrderByte, expectedLowOrderByte, serializedScalar)
	}

	expectedHighOrderByte := 0b00111111
	if highOrderByte := serializedScalar[0]; highOrderByte != byte(expectedHighOrderByte) {
		t.Fatalf("High order byte of serialized scalar is %x, expected %x. serialized bytes: % x", highOrderByte, expectedHighOrderByte, serializedScalar)
	}
}

func TestSerializeVector(t *testing.T) {
	testVector := Vector3D{1.0, 0.0, 2.0}
	serializedVector := testVector.serialize()

	expectedFirstHighOrderByte := 0b00111111
	if firstHighOrderByte := serializedVector[0]; firstHighOrderByte != byte(expectedFirstHighOrderByte) {
		t.Fatalf("High order byte of serialized scalar is %x, expected %x. serialized bytes: % x", firstHighOrderByte, expectedFirstHighOrderByte, serializedVector)
	}

	expectedThirdHighOrderByte := 0b01000000
	if ThirdHighOrderByte := serializedVector[16]; ThirdHighOrderByte != byte(expectedThirdHighOrderByte) {
		t.Fatalf("High order byte of serialized scalar is %x, expected %x. serialized bytes: % x", ThirdHighOrderByte, expectedThirdHighOrderByte, serializedVector)
	}

}

func TestSerializeSphere(t *testing.T) {
	testSphere := Sphere(Vector3D{0, 0, 0}, 0)
	serializedSphere := testSphere.body()

	actualLength := len(serializedSphere)
	expectedLength := 96
	if actualLength != expectedLength {
		t.Fatalf("Serialized sphere has unexpected length %v. expected %v. serialized bytes: % x", actualLength, expectedLength, serializedSphere)
	}
}
