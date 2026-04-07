package materials

type Material struct {
	Mass   float64
	Color  [3]byte
	isGlow bool
}

func NewMaterial(mass float64, color [3]byte) Material {
	return Material{
		Mass:   mass,
		Color:  color,
		isGlow: false,
	}
}

func GetAir() Material {
	return NewMaterial(
		1.00,
		[3]byte{0, 0, 0},
	)
}

func GetGlass() Material {
	return NewMaterial(
		0.67,
		[3]byte{50, 60, 70},
	)
}

func GetDiamond() Material {
	return NewMaterial(
		0.41,
		[3]byte{80, 90, 100},
	)
}

func GetWall() Material {
	return NewMaterial(
		0.00,
		[3]byte{20, 20, 20},
	)
}

func (m Material) MakeGlow() Material {
	m.isGlow = true
	return m
}

func (m Material) IsGlow() bool {
	return m.isGlow
}
