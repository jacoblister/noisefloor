// +build ignore

package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const basePackageName = "github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor"

type Parameter struct {
	name         string
	dataType     string
	min          float32
	max          float32
	valueDefault float32
}

type Processor struct {
	name        string
	packageName string
	inputs      []string
	outputs     []string
	parameters  []Parameter
	methods     map[string]bool
}

func readProcessor(packageName string, filename string) Processor {
	processor := Processor{packageName: packageName, methods: map[string]bool{}}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		panic("Could not parse file: " + filename)
	}

	ast.Inspect(node, func(n ast.Node) bool {
		ts, ok := n.(*ast.TypeSpec)
		if ok {
			processor.name = ts.Name.String()
		}

		fd, ok := n.(*ast.FuncDecl)
		if ok {
			processor.methods[fd.Name.String()] = true

			if fd.Name.String() == "Process" {
				for i := 0; i < fd.Type.Params.NumFields(); i++ {
					processor.inputs = append(processor.inputs, fd.Type.Params.List[i].Names[0].String())
				}
				for i := 0; i < fd.Type.Results.NumFields(); i++ {
					processor.outputs = append(processor.outputs, fd.Type.Results.List[i].Names[0].String())
				}
			}
		}

		st, ok := n.(*ast.StructType)
		if ok {
			for i := 0; i < st.Fields.NumFields(); i++ {
				if st.Fields.List[i].Tag != nil {
					parameter := Parameter{}
					parameter.name = st.Fields.List[i].Names[0].String()
					parameter.dataType = st.Fields.List[i].Type.(*ast.Ident).Name

					var tag reflect.StructTag
					tagValue := st.Fields.List[i].Tag.Value
					tag = reflect.StructTag(st.Fields.List[i].Tag.Value[1:len(tagValue)])

					fval, _ := strconv.ParseFloat(tag.Get("default"), 32)
					parameter.valueDefault = float32(fval)
					fval, _ = strconv.ParseFloat(tag.Get("min"), 32)
					parameter.min = float32(fval)
					fval, _ = strconv.ParseFloat(tag.Get("max"), 32)
					parameter.max = float32(fval)

					processor.parameters = append(processor.parameters, parameter)
				}
			}
		}

		return true
	})

	return processor
}

func readProcessors(directory string) []Processor {
	processors := []Processor{}

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		panic("Could not read directory: " + directory)
	}

	for _, file := range files {
		if file.Name() == "z_auto.go" || file.Name() == "z_factory.go" {
			continue
		}
		processor := readProcessor(directory, directory+"/"+file.Name())
		if len(processor.name) > 0 {
			processors = append(processors, processor)
		}
	}

	return processors
}

func writeMethodStart(f *os.File, processor Processor) {
	f.WriteString("// Start - init module\n")
	f.WriteString("func (r *" + processor.name + ") Start(sampleRate int, connectedMask int) {}\n\n")
}

func writeMethodStop(f *os.File, processor Processor) {
	f.WriteString("// Stop - release module\n")
	f.WriteString("func (r *" + processor.name + ") Stop() {}\n\n")
}

func writeMethodDefinition(f *os.File, processor Processor) {
	f.WriteString("// Definition exports the constant definition\n")
	f.WriteString("func (r *" + processor.name + ") Definition() (name string, inputs []string, outputs []string, parameters []processor.Parameter) {\n")
	inputs := []string{}
	outputs := []string{}
	for _, input := range processor.inputs {
		inputs = append(inputs, "\""+input+"\"")
	}
	for _, output := range processor.outputs {
		outputs = append(outputs, "\""+output+"\"")
	}
	f.WriteString("\treturn \"" + processor.name + "\", []string{" + strings.Join(inputs, ",") + "}, []string{" + strings.Join(outputs, ",") + "},\n")
	f.WriteString("\t[]processor.Parameter{\n")
	for _, parameter := range processor.parameters {
		min := strconv.FormatFloat(float64(parameter.min), 'f', -1, 32)
		max := strconv.FormatFloat(float64(parameter.max), 'f', -1, 32)
		valueDefault := strconv.FormatFloat(float64(parameter.valueDefault), 'f', -1, 32)
		f.WriteString("\t\tprocessor.Parameter{Name: \"" + parameter.name + "\", Min: " + min + ", Max: " + max + ", Default: " + valueDefault + ", Value: float32(r." + parameter.name + ")},\n")
	}
	f.WriteString("\t}\n")

	f.WriteString("}\n\n")
}

func writeMethodProcessArgs(f *os.File, processor Processor) {
	f.WriteString("//ProcessArgs calls process with an array of input/output samples\n")
	f.WriteString("func (r *" + processor.name + ") ProcessArgs(in []float32) (output []float32) {\n")
	inputs := []string{}
	outputs := []string{}
	for i := 0; i < len(processor.inputs); i++ {
		inputs = append(inputs, "in["+strconv.Itoa(i)+"]")
	}
	for i := 0; i < len(processor.outputs); i++ {
		outputs = append(outputs, "out"+strconv.Itoa(i))
	}

	f.WriteString("\t")
	if len(outputs) > 0 {
		f.WriteString(strings.Join(outputs, ",") + " := ")
	}
	f.WriteString("r.Process(" + strings.Join(inputs, ",") + ")\n")
	f.WriteString("\treturn []float32{" + strings.Join(outputs, ",") + "}\n")

	f.WriteString("}\n\n")
}

func writeMethodProcessSamples(f *os.File, processor Processor) {
	f.WriteString("//ProcessSamples calls process with an array of input/output samples\n")
	f.WriteString("func (r *" + processor.name + ") ProcessSamples(in [][]float32, out [][]float32, length int) {\n")

	outputs := []string{}
	inputs := []string{}
	for i := 0; i < len(processor.inputs); i++ {
		inputs = append(inputs, "in["+strconv.Itoa(i)+"][i]")
	}
	for i := 0; i < len(processor.outputs); i++ {
		outputs = append(outputs, "out["+strconv.Itoa(i)+"][i]")
	}
	f.WriteString("\tfor i := 0; i < length; i++ {\n")
	f.WriteString("\t\t" + strings.Join(outputs, ", ") + " = r.Process(" + strings.Join(inputs, ", ") + ")\n")
	f.WriteString("\t}\n")
	f.WriteString("}\n\n")
}

func writeMethodSetParameter(f *os.File, processor Processor) {
	f.WriteString("//SetParameter set a single processor parameter\n")
	f.WriteString("func (r *" + processor.name + ") SetParameter(index int, value float32) {\n")
	f.WriteString("\tswitch index {\n")
	for i, parameter := range processor.parameters {
		f.WriteString("\tcase " + strconv.Itoa(i) + ":\n")
		f.WriteString("\t\tr." + parameter.name + " = ")
		if parameter.dataType == "float32" {
			f.WriteString("value\n")
		} else {
			f.WriteString(parameter.dataType + "(value + 0.5)\n")
		}
	}

	f.WriteString("\t} \n")
	f.WriteString("}\n\n")
}

func writeMethods(packageName string, processors []Processor) {
	f, err := os.Create(packageName + "/z_factory.go")
	if err != nil {
		panic("Could not open factory file: " + packageName + "/z_factory.go")
	}
	f.WriteString("package " + packageName + "\n\n")

	needImport := false
	for _, processor := range processors {
		if _, ok := processor.methods["Definition"]; !ok {
			needImport = true
		}
	}
	if needImport {
		f.WriteString("import \"" + basePackageName + "\"\n\n")
	}

	for _, processor := range processors {
		if _, ok := processor.methods["Start"]; !ok {
			writeMethodStart(f, processor)
		}
		if _, ok := processor.methods["Stop"]; !ok {
			writeMethodStop(f, processor)
		}
		if _, ok := processor.methods["Definition"]; !ok {
			writeMethodDefinition(f, processor)
		}
		if _, ok := processor.methods["SetParameter"]; !ok {
			writeMethodSetParameter(f, processor)
		}
		if _, ok := processor.methods["ProcessArgs"]; !ok {
			writeMethodProcessArgs(f, processor)
		}
		if _, ok := processor.methods["ProcessSamples"]; !ok {
			writeMethodProcessSamples(f, processor)
		}
	}
	f.Close()
}

func writeFactory(dirs []string, processors []Processor) {
	processorNames := []string{}
	for _, processor := range processors {
		processorNames = append(processorNames, "\""+processor.name+"\"")
	}

	f, err := os.Create("processorfactory/factory.go")
	if err != nil {
		panic("Could not open factory file")
	}
	f.WriteString("package processorfactory\n\n")
	f.WriteString("import (\n")
	f.WriteString("\t\"" + basePackageName + "\"\n")
	for _, dir := range dirs {
		f.WriteString("\t\"" + basePackageName + "/" + dir + "\"\n")
	}
	f.WriteString(")\n\n")

	f.WriteString("// ListProcessors returns a list of available processors\n")
	f.WriteString("func ListProcessors() []string {\n")
	f.WriteString("\treturn []string {" + strings.Join(processorNames, ",") + "}\n")
	f.WriteString("}\n\n")

	f.WriteString("// MakeProcessor generates a new processor by the given processor name\n")
	f.WriteString("func MakeProcessor(name string) processor.Processor {\n")
	f.WriteString("\tvar proc processor.Processor\n\n")
	f.WriteString("\tswitch name {\n")
	for _, processor := range processors {
		f.WriteString("\tcase \"" + processor.name + "\":\n")
		f.WriteString("\t\tproc = &" + processor.packageName + "." + processor.name + "{}\n")
	}

	f.WriteString("\t}\n\n")
	f.WriteString("\tprocessor.SetProcessorDefaults(proc)\n\n")
	f.WriteString("\treturn proc\n")
	f.WriteString("}\n")

	f.Close()
}

func main() {
	allProcessors := []Processor{}
	dirs := []string{"processorbuiltin", "processorbasic"}

	for _, dir := range dirs {
		processors := readProcessors(dir)
		writeMethods(dir, processors)

		allProcessors = append(allProcessors, processors...)
	}

	writeFactory(dirs, allProcessors)
}
