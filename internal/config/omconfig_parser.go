package config

import (
	"fmt"
	"math"
	"strings"

	"github.com/LeviyLokotb/light-automata/pkg/materials"
)

func parseObjects(omconf OMConfig) ([]materials.Object, error) {
	var objects []materials.Object

	for _, oconf := range omconf.Objects {
		material, err := getMaterialByName(oconf.Material)
		if err != nil {
			return []materials.Object{}, err
		}

		var object materials.Object
		switch oconf.Shape {
		case "sphere":
			object, err = parseSphere(material, oconf.Params)
		case "rect":
			object, err = parseRect(material, oconf.Params)
		case "triangle":
			object, err = parseTriangle(material, oconf.Params)
		default:
			return []materials.Object{}, fmt.Errorf("unknown shape '%s'", oconf.Shape)
		}
		if err != nil {
			return []materials.Object{}, err
		}

		objects = append(objects, object)
	}

	return objects, nil
}

func getMaterialByName(name string) (materials.Material, error) {
	isGlow := false
	material := materials.GetAir()

	words := strings.Split(name, " ")
	for _, word := range words {
		switch word {
		case "glow":
			isGlow = true
		case "air":
			material = materials.GetAir()
		case "glass":
			material = materials.GetGlass()
		case "diamond":
			material = materials.GetDiamond()
		case "wall":
			material = materials.GetWall()
		default:
			return materials.Material{}, fmt.Errorf("unknown material '%s'", name)
		}
	}

	if isGlow {
		material = material.MakeGlow()
	}
	return material, nil
}

func parseSphere(material materials.Material, params Params) (materials.Object, error) {
	radius, err := params.get("radius", "sphere")
	if err != nil {
		return materials.Object{}, err
	}

	cX, err := params.get("center_x", "sphere")
	if err != nil {
		return materials.Object{}, err
	}

	cY, err := params.get("center_y", "sphere")
	if err != nil {
		return materials.Object{}, err
	}

	return materials.NewSphere(material, radius, cX, cY), nil
}

func parseRect(material materials.Material, params Params) (materials.Object, error) {
	minX, err := params.get("min_x", "rect")
	if err != nil {
		return materials.Object{}, err
	}

	maxX, err := params.get("max_x", "rect")
	if err != nil {
		return materials.Object{}, err
	}

	minY, err := params.get("min_y", "rect")
	if err != nil {
		return materials.Object{}, err
	}

	maxY, err := params.get("max_y", "rect")
	if err != nil {
		return materials.Object{}, err
	}

	return materials.NewRect(material, minX, maxX, minY, maxY), nil
}

func parseTriangle(material materials.Material, params Params) (materials.Object, error) {
	ax, err := params.get("ax", "triangle")
	if err != nil {
		return materials.Object{}, err
	}

	bx, err := params.get("bx", "triangle")
	if err != nil {
		return materials.Object{}, err
	}

	cx, err := params.get("cx", "triangle")
	if err != nil {
		return materials.Object{}, err
	}

	ay, err := params.get("ay", "triangle")
	if err != nil {
		return materials.Object{}, err
	}

	by, err := params.get("by", "triangle")
	if err != nil {
		return materials.Object{}, err
	}

	cy, err := params.get("cy", "triangle")
	if err != nil {
		return materials.Object{}, err
	}

	angle, err := params.get("rotated_by", "triangle")
	if err == nil {
		ag := float64(angle%360) * math.Pi / 180
		ax, ay, bx, by, cx, cy = materials.RotateTriangleInt(ax, ay, bx, by, cx, cy, ag)
	}

	return materials.NewTriangle(material, ax, ay, bx, by, cx, cy), nil
}

func (p Params) get(key string, objectName string) (int, error) {
	v, ok := p[key]
	if !ok {
		return 0, fmt.Errorf("%s requires '%s' parameter", objectName, key)
	}
	return v, nil
}
