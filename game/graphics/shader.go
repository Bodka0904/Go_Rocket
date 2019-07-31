package graphics

import (
	"fmt"
	"strings"

	"Go-Rocket/src/game/dependencies/mathgl/mathgl/mgl32"

	"Go-Rocket/src/game/dependencies/go-gl/gl/v4.1-core/gl"
	"golang.org/x/image/math/f32"
)

type Shader struct {
	program uint32
	uniform map[string]*int32
}

// InitShader - Creates new shader program
func (shader *Shader) InitShader(vertexSrc string, fragmentSrc string) {
	vertexShader, err := compileShader(vertexSrc, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	fragmentShader, err := compileShader(fragmentSrc, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	shader.program = gl.CreateProgram()
	gl.AttachShader(shader.program, vertexShader)
	gl.AttachShader(shader.program, fragmentShader)

	shader.uniform = make(map[string]*int32)
}

func (shader *Shader) AddGeometryShader(geometrySrc string) {
	geometryShader, err := compileShader(geometrySrc, gl.GEOMETRY_SHADER)
	if err != nil {
		panic(err)
	}

	gl.AttachShader(shader.program, geometryShader)
}

// Bind - bind shader for use
func (shader *Shader) Bind() {
	gl.UseProgram(shader.program)
}

func (shader *Shader) AddAttrib(location uint32, name string) {
	gl.BindAttribLocation(shader.program, location, gl.Str(name+"\x00"))
}

func (shader *Shader) GetUniform(name string) int32 {
	return *shader.uniform[name]
}

func (shader *Shader) AddUniform(name string) {
	shader.uniform[name] = new(int32)
	*shader.uniform[name] = gl.GetUniformLocation(shader.program, gl.Str(name+"\x00"))
}

func (shader *Shader) UploadMat4(name string, value mgl32.Mat4) {
	gl.UniformMatrix4fv(*shader.uniform[name], 1, false, &value[0])
}
func (shader *Shader) UploadVec3(name string, value f32.Vec3) {
	gl.Uniform3fv(*shader.uniform[name], 1, &value[0])
}
func (shader *Shader) UploadVec2(name string, value f32.Vec2) {
	gl.Uniform2fv(*shader.uniform[name], 1, &value[0])
}

func (shader *Shader) LinkShader() {
	gl.LinkProgram(shader.program)
}

// compileShader - function that takes as input source of shader( string ) and shaderType(Vertex/Fragment shader)
func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
