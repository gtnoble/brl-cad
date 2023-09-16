package geometry

import "brl-cad/object"

func (e Ellipsoid) Object(name string) object.DbObject {
	return object.MakeGeometry(name, object.ELLIPSOID, e.serialize())
}

func (s Sphere) Object(name string) object.DbObject {
	return object.MakeGeometry(name, object.SPHERE, s.serialize())
}

func (a ARB8) Object(name string) object.DbObject {
	return object.MakeGeometry(name, object.ARB8, a.serialize())
}

func (t Torus) Object(name string) object.DbObject {
	return object.MakeGeometry(name, object.TORUS, t.serialize())
}

func (h Halfspace) Object(name string) object.DbObject {
	return object.MakeGeometry(name, object.HALF_SPACE, h.serialize())
}
