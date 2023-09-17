package object

import (
	"encoding/binary"
	"math"
)

const (
	SCALAR_SIZE    = 8
	VECTOR_3D_SIZE = 3 * SCALAR_SIZE
)

type Scalar float64

type Vector3D [3]float64

type serializable interface {
	serialize() []byte
}

func Ellipsoid(vertex Vector3D, radiusA Vector3D, radiusB Vector3D, radiusC Vector3D) DbObject {
	body := func() []byte {
		return serialize([]serializable{
			vertex,
			radiusA,
			radiusB,
			radiusC,
		})
	}

	return BrlCadObject(
		"ell",
		ELLIPSOID,
		nil,
		body,
	)
}

func Sphere(vertex Vector3D, radius Scalar) DbObject {
	directedRadius := Vector3D{0, 0, float64(radius)}
	return Ellipsoid(vertex, directedRadius, directedRadius, directedRadius)
}

func Arb8(quad1 [4]Vector3D, quad2 [4]Vector3D) DbObject {
	body := func() []byte {
		return serialize([]serializable{quad1[0], quad1[1], quad1[2], quad1[3], quad2[0], quad2[1], quad2[2], quad2[3]})
	}

	return BrlCadObject(
		"arb8",
		ARB8,
		nil,
		body,
	)
}

func Torus(vertex Vector3D, normal Vector3D, majorRadius Scalar, minorRadius Scalar) DbObject {
	body := func() []byte {
		return serialize([]serializable{vertex, normal, majorRadius, minorRadius})
	}

	return BrlCadObject(
		"tor",
		TORUS,
		nil,
		body,
	)
}

func Halfspace(normal Vector3D, originDistance Scalar) DbObject {
	body := func() []byte {
		return serialize([]serializable{normal, originDistance})
	}

	return BrlCadObject(
		"half",
		HALF_SPACE,
		nil,
		body,
	)
}

func (n Scalar) serialize() []byte {
	buf := make([]byte, SCALAR_SIZE)
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(float64(n)))
	return buf
}

func (v Vector3D) serialize() []byte {
	buf := make([]byte, VECTOR_3D_SIZE)
	for i, element := range v {
		binary.BigEndian.PutUint64(buf[i*8:], math.Float64bits(element))
	}
	return buf
}

func serialize(p []serializable) []byte {
	var buf []byte
	for _, value := range p {
		vBytes := value.serialize()
		buf = append(buf, vBytes[:]...)
	}
	return buf
}
