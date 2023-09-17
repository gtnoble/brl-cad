package object

import (
	"fmt"

	"github.com/google/uuid"
)

// DLI Flags
const (
	APPLICATION_DATA_OBJECT_FLAG = iota
	HEADER_OBJECT_FLAG
	FREE_DB_STORAGE_FLAG
)

// AZ/BZ Flags
const (
	NO_COMPRESSION_FLAG = iota
	GZIP_COMPRESSION_FLAG
	BZIP_COMPRESSION_FLAG
)

const HIDDEN_OBJECT_FLAG = 0x4

// Major object types
const (
	RESERVED                   = 0
	BRL_CAD_OBJECT             = 1
	ATTRIBUTE_ONLY_OBJECT      = 2
	EXPERIMENTAL_BINARY_OBJECT = 0x18
	UNIFORM_BINARY_OBJECT      = 9
	MIME_TYPED_BINARY_OBJECT   = 10
)

// Nongeometry minor types
const (
	_ = iota
	COMBINATION_TYPE
	GRIP_TYPE
	JOINT_TYPE
)

// Geometry minor types
const (
	_ = iota
	TORUS
	TRUNCATED_GENERAL_CONE
	ELLIPSOID
	ARB8
	ARS
	HALF_SPACE
	RIGHT_ELLIPTICAL_CYLINDER
	POLYSOLID
	B_SPLINE_SOLID
	SPHERE
	N_MANIFOLD_GEOMETRY
	EXTRUDED_BIT_MAP
	VOLUME
	ARBN
	PIPE
	PARTICLE
	RIGHT_PARABOLIC_CYLINDER
	RIGHT_HYPERBOLIC_CYLINDER
	ELLIPTICAL_PARABOLOID
	ELLIPTICAL_HYPERBOLOID
	ELLIPTICAL_TORUS
	GRIP_NONGEOMETRIC
	JOINT_NONGEOMETRIC
	HEIGHT_FIELD
	DISPLACEMENT_MAP
	SKETCH
	EXTRUDE
	SUBMODEL
	CLINE
	BAG_O_TRIANGLES
	COMBINATION_RECORD
	EXPERIMENTAL_BINARY
	UNIFORM_ARRAY_BINARY
	MIME_TYPED_BINARY
	SUPERQUADRATIC_ELLIPSOID
	METABALL
	BREP
	HYPERBOLOID
	CONSTRAINT
	SOLID_OF_REVOLUTION
	COLLECTION_OF_POINTS
)

const GLOBAL_OBJECT_NAME = "_GLOBAL"

const BIT_WID_FLAG_64 = 0b11

type DbObject struct {
	dli             byte
	majorType       byte
	minorType       byte
	isHidden        bool
	compressionFlag byte
	attributes      map[string]string
	Name            *string
	body            func() []byte
}

func (db *DbObject) node() *combinationNode {
	return &combinationNode{
		OPERAND,
		nil,
		nil,
		db,
	}
}

func generateName(prefix string) string {
	uniqueId := uuid.New().String()
	return fmt.Sprintf("%s-%s", prefix, uniqueId)
}

func BrlCadObject(baseName string, minorType byte, attributes map[string]string, body func() []byte) DbObject {
	name := generateName(baseName)
	return DbObject{
		majorType:  BRL_CAD_OBJECT,
		minorType:  minorType,
		attributes: attributes,
		Name:       &name,
		body:       body,
	}
}

func GlobalObject(title string, unitConversion float64) DbObject {
	globalObjectName := GLOBAL_OBJECT_NAME
	return DbObject{
		dli:       APPLICATION_DATA_OBJECT_FLAG,
		isHidden:  true,
		majorType: ATTRIBUTE_ONLY_OBJECT,
		minorType: RESERVED,
		Name:      &globalObjectName,
		attributes: map[string]string{
			"title": title,
			"units": fmt.Sprintf("%E", unitConversion),
		},
	}
}
