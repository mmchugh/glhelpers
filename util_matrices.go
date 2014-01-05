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
		x_axis[0], y_axis[0], -z_axis[0], 0,
		x_axis[1], y_axis[1], -z_axis[1], 0,
		x_axis[2], y_axis[2], -z_axis[2], 0,
		0, 0, 0, 1,
	}.Translate(-eyeX, -eyeY, -eyeZ)
}
