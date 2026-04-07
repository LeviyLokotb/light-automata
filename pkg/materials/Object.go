package materials

import "math"

type ShapeFunc func(int, int) bool

type Object struct {
	Material
	contain ShapeFunc
}

func NewObject(material Material, contains ShapeFunc) Object {
	return Object{
		Material: material,
		contain:  contains,
	}
}

func (o Object) Contain(x, y int) bool {
	return o.contain(x, y)
}

func NewSphere(material Material, radius int, centerX, centerY int) Object {
	return NewObject(material, func(x, y int) bool {
		return math.Pow(float64(x-centerX), 2)+
			math.Pow(float64(y-centerY), 2) <=
			float64(radius*radius)
	})
}

func NewRect(material Material, minX, maxX, minY, maxY int) Object {
	return NewObject(material, func(x, y int) bool {
		return x >= minX &&
			x <= maxX &&
			y >= minY &&
			y <= maxY
	})
}

func NewTriangle(material Material, Ax, Ay, Bx, By, Cx, Cy int) Object {
	return NewObject(material, func(x, y int) bool {
		baseS := triangleArea(Ax, Ay, Bx, By, Cx, Cy)

		sumS := triangleArea(x, y, Bx, By, Cx, Cy) +
			triangleArea(Ax, Ay, x, y, Cx, Cy) +
			triangleArea(Ax, Ay, Bx, By, x, y)

		return math.Abs(sumS-baseS) < 1e-6
	})
}

func NewLine(material Material, x1, y1, x2, y2, thickness int) Object {
	return NewObject(material, func(x, y int) bool {
		// Вычисляем расстояние от точки (x,y) до отрезка (x1,y1)-(x2,y2)

		// Длина отрезка в квадрате
		dx := float64(x2 - x1)
		dy := float64(y2 - y1)
		len2 := dx*dx + dy*dy

		if len2 == 0 {
			// Отрезок вырожден в точку
			return math.Hypot(float64(x-x1), float64(y-y1)) <= float64(thickness)
		}

		// Параметр t - проекция точки на отрезок
		t := float64((x-x1)*int(dx)+(y-y1)*int(dy)) / len2

		// Находим ближайшую точку на отрезке
		var nearestX, nearestY float64
		if t < 0 {
			nearestX, nearestY = float64(x1), float64(y1)
		} else if t > 1 {
			nearestX, nearestY = float64(x2), float64(y2)
		} else {
			nearestX = float64(x1) + t*dx
			nearestY = float64(y1) + t*dy
		}

		// Расстояние от точки до ближайшей точки на отрезке
		distance := math.Hypot(float64(x)-nearestX, float64(y)-nearestY)

		return distance <= float64(thickness)
	})
}

func triangleArea(Ax, Ay, Bx, By, Cx, Cy int) float64 {
	a := math.Hypot(float64(Bx-Cx), float64(By-Cy))
	b := math.Hypot(float64(Ax-Cx), float64(Ay-Cy))
	c := math.Hypot(float64(Ax-Bx), float64(Ay-By))
	p := (a + b + c) / 2
	return math.Sqrt(p * (p - a) * (p - b) * (p - c))
}

func RotateTriangle(x1, y1, x2, y2, x3, y3 float64, angle float64) (float64, float64, float64, float64, float64, float64) {
	cx := (x1 + x2 + x3) / 3
	cy := (y1 + y2 + y3) / 3

	cosA := math.Cos(angle)
	sinA := math.Sin(angle)

	rotate := func(x, y float64) (float64, float64) {
		dx := x - cx
		dy := y - cy
		return cx + dx*cosA - dy*sinA, cy + dx*sinA + dy*cosA
	}

	r1x, r1y := rotate(x1, y1)
	r2x, r2y := rotate(x2, y2)
	r3x, r3y := rotate(x3, y3)

	return r1x, r1y, r2x, r2y, r3x, r3y
}
