package lowlevel

import (
	"github.com/go-gl/gl"
	"io/ioutil"
)

//TODO
//Better error handling

type Program struct {
	gl.Program
}

//Create program from shader file paths
//Vertex Shader path first then Fragment Shader
func LoadShaders(vShade, fShade string) Program {
	//create and compile shaders
	vShader := createShader(vShade, gl.VERTEX_SHADER)
	fShader := createShader(fShade, gl.FRAGMENT_SHADER)

	//attach them to a new program
	prog := Program{gl.CreateProgram()}
	prog.AttachShader(vShader)
	prog.AttachShader(fShader)
	prog.Link()

	//enable program
	prog.Use()
	return prog
}

//Maps the sampler in the program to a texture
func (prog Program) MapSampler(smpler2Dname string, num int) {
	texSampler := prog.GetUniformLocation(smpler2Dname)
	texSampler.Uniform1i(num)
}

//for internal use, gets the source from the file specified in the path
//then compiles
func createShader(path string, shtype gl.GLenum) gl.Shader {
	rawdata, _ := ioutil.ReadFile(path)
	src := string(rawdata)
	shader := gl.CreateShader(shtype)
	shader.Source(src)
	shader.Compile()
	return shader
}

func (prog Program) LinkUBO(varName string, ubo gl.Buffer) {
	CheckError("T")
	ubo.BindBufferBase(gl.UNIFORM_BUFFER, uint(1))
	CheckError("T3")
}
