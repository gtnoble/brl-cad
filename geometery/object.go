package geometry

import "brl-cad/object"

func (e Ellipsoid) object(name string) object.DbObject {
	return object.MakeGeometry(name, object.ELLIPSOID, e.serialize())
}

func (s Sphere) object(name string) object.DbObject {
	return object.MakeGeometry(name, object.SPHERE, s.serialize())
}

func (a ARB8) object(name string) object.DbObject {
	return object.MakeGeometry(name, object.ARB8, a.serialize())
}

func (t Torus) object(name string) object.DbObject {
	return object.MakeGeometry(name, object.TORUS, t.serialize())
}

func (h Halfspace) object(name string) object.DbObject {
	return object.MakeGeometry(name, object.HALF_SPACE, h.serialize())
}
