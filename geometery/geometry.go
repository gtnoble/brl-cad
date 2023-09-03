package geometry

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

type Ellipsoid struct {
	Vertex  Vector3D
	RadiusA Vector3D
	RadiusB Vector3D
	RadiusC Vector3D
}

func (s Ellipsoid) serialize() []byte {
	return serialize([]serializable{
		s.Vertex,
		s.RadiusA,
		s.RadiusB,
		s.RadiusC,
	})
}

type Sphere struct {
	Vertex Vector3D
	Radius Scalar
}

func (s Sphere) serialize() []byte {
	return serialize([]serializable{s.Vertex, s.Radius})
}

type ARB8 [8]Vector3D

func (p ARB8) serialize() []byte {
	return serialize([]serializable{p[0], p[1], p[2], p[3], p[4], p[5], p[6], p[7]})
}

type Torus struct {
	Vertex      Vector3D
	Normal      Vector3D
	MajorRadius Scalar
	MinorRadius Scalar
}

func (p Torus) serialize() []byte {
	return serialize([]serializable{p.Vertex, p.Normal, p.MajorRadius, p.MinorRadius})
}

type Halfspace struct {
	Normal         Vector3D
	OriginDistance Scalar
}

func (p Halfspace) serialize() []byte {
	return serialize([]serializable{p.Normal, p.OriginDistance})
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
