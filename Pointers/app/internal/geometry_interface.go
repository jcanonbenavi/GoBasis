package internal

const (
	RectType   = "RECTANGLE"
	CircleType = "CIRCLE"
)

// return the geometry type
func NewGeometry(geometryType string, values ...float64) Geometry {
	switch geometryType {
	case RectType:
		return Rectangle{Width: values[0], Height: values[1]}
	case CircleType:
		return Circle{Radius: values[0]}
	}
	return nil
}
