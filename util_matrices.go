package glhelpers

import "math"

func Perspective(fovy, aspect, near, far float32) Mat4 {
	f := float32(1.0/math.Tan(float64(fovy)/2.0))

	return Mat4{
		f/aspect, 0, 0, 0,
		0, f, 0, 0,
		0, 0, (near + far) / (near - far), -1,
		0, 0, (2 * near * far) / (near - far), 0,
	}
}

func LookAt(eyeX, eyeY, eyeZ, atX, atY, atZ, upX, upY, upZ float32) Mat4 {
	z_axis := Vec3{atX - eyeX, atY - eyeY, atZ - eyeZ}.Normalize()
	up := Vec3{upX, upY, upZ}.Normalize()
//	eye := Vec3{eyeX, eyeY, eyeZ}

	x_axis := z_axis.Cross(up).Normalize()
	y_axis := x_axis.Cross(z_axis)

	return Mat4{
		x_axis.X, y_axis.X, -z_axis.X, 0,
		x_axis.Y, y_axis.Y, -z_axis.Y, 0,
		x_axis.Z, y_axis.Z, -z_axis.Z, 0,
		0, 0, 0, 1,
	}.Translate(-eyeX, -eyeY, -eyeZ)
}
