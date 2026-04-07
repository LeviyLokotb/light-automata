package materials

type ObjectsManager []Object

func NewObjectsManager(objects []Object) *ObjectsManager {
	om := (ObjectsManager)(objects)
	return &om
}

func (om ObjectsManager) Append(obj Object) ObjectsManager {
	newOm := append(([]Object)(om), obj)
	return (ObjectsManager)(newOm)
}

func (om *ObjectsManager) Add(obj Object) {
	newOm := om.Append(obj)
	om = &newOm
}

func (om ObjectsManager) GetMaterialAt(x, y int) Material {
	for i := len(om) - 1; i >= 0; i-- {
		if om[i].Contain(x, y) {
			return om[i].Material
		}
	}
	return GetAir()
}
