package glhelpers

import "math"

func (mat Mat4) RotateX(angle float32) Mat4 {
	sin, cos := float32(math.Sin(float64(angle))), float32(math.Cos(float64(angle)))
	rotation := Mat4{
		1, 0, 0, 0,
		0, cos, sin, 0,
		0, -sin, cos, 0,
		0, 0, 0, 1,
	}
	return mat.Mult(rotation)
}

func (mat Mat4) RotateY(angle float32) Mat4 {
	sin, cos := float32(math.Sin(float64(angle))), float32(math.Cos(float64(angle)))
	rotation := Mat4{
		cos, 0, -sin, 0,
		0, 1, 0, 0,
		sin, 0, cos, 0,
		0, 0, 0, 1,
	}
	return mat.Mult(rotation)
}

func (mat Mat4) RotateZ(angle float32) Mat4 {
	sin, cos := float32(math.Sin(float64(angle))), float32(math.Cos(float64(angle)))
	rotation := Mat4{
		cos, sin, 0, 0,
		-sin, cos, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
	return mat.Mult(rotation)
}

func (mat Mat4) Scale(x_scale, y_scale, z_scale float32) Mat4 {
	scale := Mat4{
		x_scale, 0, 0, 0,
		0, y_scale, 0, 0,
		0, 0, z_scale, 0,
		0, 0, 0, 1,
	}
	return mat.Mult(scale)
}

func (mat Mat4) Translate(x, y, z float32) Mat4 {
	translate := Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		x, y, z, 1,
	}
	return mat.Mult(translate)
}
