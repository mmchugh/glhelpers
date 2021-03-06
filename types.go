package glhelpers

import "math"

type Vec2 [2]float32

type Vec3 [3]float32

type Mat4 [16]float32

func (a Vec2) Add(b Vec2) Vec2 {
	return Vec2{a[0] + b[0], a[1] + b[1]}
}

func (a Vec2) Sub(b Vec2) Vec2 {
	return Vec2{a[0] - b[0], a[1] - b[1]}
}

func (a Vec2) Mult(scalar float32) Vec2 {
	return Vec2{a[0] * scalar, a[1] * scalar}
}

func (a Vec2) Dot(b Vec2) float32 {
	return a[0]*b[0] + a[1]*b[1]
}

func (a Vec2) Len() float32 {
	return float32(math.Sqrt(float64(a[0]*a[0] + a[1]*a[1])))
}

func (a Vec2) Normalize() Vec2 {
	length := a.Len()
	return Vec2{a[0] / length, a[1] / length}
}

func (a Vec3) Add(b Vec3) Vec3 {
	return Vec3{a[0] + b[0], a[1] + b[1], a[2] + b[2]}
}

func (a Vec3) Sub(b Vec3) Vec3 {
	return Vec3{a[0] - b[0], a[1] - b[1], a[2] - b[2]}
}

func (a Vec3) Mult(scalar float32) Vec3 {
	return Vec3{a[0] * scalar, a[1] * scalar, a[2] * scalar}
}

func (a Vec3) Dot(b Vec3) float32 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

func (a Vec3) Cross(b Vec3) Vec3 {
	return Vec3{
		a[1]*b[2] - a[2]*b[1],
		a[2]*b[0] - a[0]*b[2],
		a[0]*b[1] - a[1]*b[0],
	}
}

func (a Vec3) Len() float32 {
	return float32(math.Sqrt(float64(a[0]*a[0] + a[1]*a[1] + a[2]*a[2])))
}

func (a Vec3) Normalize() Vec3 {
	length := a.Len()
	return Vec3{a[0] / length, a[1] / length, a[2] / length}
}

func Ident4() Mat4 {
	return Mat4{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
}

func (a Mat4) Mult(b Mat4) Mat4 {
	return Mat4{
		a[0]*b[0] + a[4]*b[1] + a[8]*b[2] + a[12]*b[3],
		a[1]*b[0] + a[5]*b[1] + a[9]*b[2] + a[13]*b[3],
		a[2]*b[0] + a[6]*b[1] + a[10]*b[2] + a[14]*b[3],
		a[3]*b[0] + a[7]*b[1] + a[11]*b[2] + a[15]*b[3],
		a[0]*b[4] + a[4]*b[5] + a[8]*b[6] + a[12]*b[7],
		a[1]*b[4] + a[5]*b[5] + a[9]*b[6] + a[13]*b[7],
		a[2]*b[4] + a[6]*b[5] + a[10]*b[6] + a[14]*b[7],
		a[3]*b[4] + a[7]*b[5] + a[11]*b[6] + a[15]*b[7],
		a[0]*b[8] + a[4]*b[9] + a[8]*b[10] + a[12]*b[11],
		a[1]*b[8] + a[5]*b[9] + a[9]*b[10] + a[13]*b[11],
		a[2]*b[8] + a[6]*b[9] + a[10]*b[10] + a[14]*b[11],
		a[3]*b[8] + a[7]*b[9] + a[11]*b[10] + a[15]*b[11],
		a[0]*b[12] + a[4]*b[13] + a[8]*b[14] + a[12]*b[15],
		a[1]*b[12] + a[5]*b[13] + a[9]*b[14] + a[13]*b[15],
		a[2]*b[12] + a[6]*b[13] + a[10]*b[14] + a[14]*b[15],
		a[3]*b[12] + a[7]*b[13] + a[11]*b[14] + a[15]*b[15],
	}
}
