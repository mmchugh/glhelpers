package glhelpers

import (
	gl "github.com/go-gl/gl"
	"io/ioutil"
)

func compile_shader(source string, shader_type gl.GLenum) gl.Shader {
	shader := gl.CreateShader(shader_type)
	shader.Source(source)
	shader.Compile()

	if shader.Get(gl.COMPILE_STATUS) != gl.TRUE {
		panic("Could not compile shader: " + shader.GetInfoLog())
	}

	return shader

}

func CreateProgram(vertex_source string, fragment_source string) gl.Program {
	vertex_shader := compile_shader(vertex_source, gl.VERTEX_SHADER)
	fragment_shader := compile_shader(fragment_source, gl.FRAGMENT_SHADER)

	program := gl.CreateProgram()
	program.AttachShader(vertex_shader)
	program.AttachShader(fragment_shader)

	program.Link()
	if program.Get(gl.LINK_STATUS) != gl.TRUE {
		panic("Could not link program: " + program.GetInfoLog())
	}

	return program
}

func CreateProgramFromPaths(vertex_path string, fragment_path string) gl.Program {
	vertex_source, err := ioutil.ReadFile(vertex_path)
	if err != nil {
		panic("Unable to load vertex shader from " + vertex_path)
	}

	fragment_source, err := ioutil.ReadFile(fragment_path)
	if err != nil {
		panic("Unable to load fragment shader from " + fragment_path)
	}

	return CreateProgram(string(vertex_source), string(fragment_source))
}
