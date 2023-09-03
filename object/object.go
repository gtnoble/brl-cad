package object

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

// Major object types
const (
	RESERVED                   = 0
	BRL_CAD_NONGEOMETRY        = 1
	BRL_CAD_GEOMETRY           = 2
	ATTRIBUTE_ONLY_OBJECT      = 3
	EXPERIMENTAL_BINARY_OBJECT = 8
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

const BIT_WID_FLAG_64 = 0b11

type DbObject struct {
	dli             byte
	majorType       byte
	minorType       byte
	compressionFlag byte
	name            *string
	attributes      map[string]string
	body            []byte
}

func MakeHeader() DbObject {
	return DbObject{dli: HEADER_OBJECT_FLAG}
}

func MakeGeometry(name string, primativeType byte, body []byte) DbObject {
	return DbObject{
		dli:       APPLICATION_DATA_OBJECT_FLAG,
		majorType: BRL_CAD_GEOMETRY,
		minorType: primativeType,
		name:      &name,
		body:      body,
	}
}
