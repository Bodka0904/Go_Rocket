package graphics

const (
	BasicVertexShaderSource = `
		#version 410
		in vec3 position;
		in vec2 texCoord;

		uniform mat4 ortho;
		uniform mat4 model;
	

		out vec2 texCoord0;

		void main() {

			texCoord0 = texCoord;
			gl_Position = ortho * model * vec4(position, 1.0);

			
		}
	` + "\x00"

	BasicFragmentShaderSource = `
		#version 410
		in vec2 texCoord0;

		uniform sampler2D tex;
		out vec4 frag_colour;

		void main() {
			frag_colour = texture(tex,texCoord0);
		}
	` + "\x00"

	InstanceVertexShaderSource = `	
			#version 330 core
			
		
			layout (location = 0) in vec3 position;
			layout (location = 1) in vec2 texCoord;
		    layout (location = 2) in vec3 offset;
 
		   
			uniform mat4 ortho;
			

			void main()
			{
				gl_Position = ortho * vec4(position + offset,1.0);
			}
	
		` + "\x00"

	InstanceFragmentShaderSource = `
			#version 330 core
			
			out vec4 color;
			
			void main()
			{
				color = vec4(0.5,0.7,0.8,1.0);
			}
			` + "\x00"
)
