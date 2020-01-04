// +build ignore

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const packageName = "github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor"

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
					processor.parameters = append(processor.parameters, parameter)
				}
			}
		}

		return true
	})

	fmt.Println(processor.name, processor.parameters)

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
	f.WriteString("func (r *" + processor.name + ") Start(sampleRate int) {}\n\n")
}

func writeMethodProcessArgs(f *os.File, processor Processor) {
	f.WriteString("//ProcessArgs calls process with an array of input/output samples\n")
	f.WriteString("func (r *" + processor.name + ") ProcessArgs(in []float32) (output []float32) {\n")
	outputs := []string{}
	inputs := []string{}
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
	f.WriteString("func (r *" + processor.name + ") ProcessSamples(in [][]float32, length int) (out [][]float32) {\n")
	f.WriteString("\tout = make([][]float32, " + strconv.Itoa(len(processor.outputs)) + ")\n")

	outputs := []string{}
	inputs := []string{}
	for i := 0; i < len(processor.inputs); i++ {
		inputs = append(inputs, "in["+strconv.Itoa(i)+"][i]")
	}
	for i := 0; i < len(processor.outputs); i++ {
		f.WriteString("\tout[" + strconv.Itoa(i) + "] = make([]float32, length)\n")
		outputs = append(outputs, "out["+strconv.Itoa(i)+"][i]")
	}
	f.WriteString("\tfor i := 0; i < length; i++ {\n")
	f.WriteString("\t\t" + strings.Join(outputs, ", ") + " = r.Process(" + strings.Join(inputs, ", ") + ")\n")
	f.WriteString("\t}\n")
	f.WriteString("\treturn\n")
	f.WriteString("}\n\n")
}

func writeMethods(packageName string, processors []Processor) {
	f, err := os.Create(packageName + "/z_auto.go")
	if err != nil {
		panic("Could not open factory file: " + packageName + "/z_auto.go")
	}
	f.WriteString("package " + packageName + "\n\n")

	for _, processor := range processors {
		if _, ok := processor.methods["Start"]; !ok {
			writeMethodStart(f, processor)
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
	f.WriteString("\t\"" + packageName + "\"\n")
	for _, dir := range dirs {
		f.WriteString("\t\"" + packageName + "/" + dir + "\"\n")
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
