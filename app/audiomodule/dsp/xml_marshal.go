package dsp

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor"
)

// MarshalXML marhalls the graph
func (g *Graph) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	println("marsall")
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "graph"}})

	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "processors"}})
	for i := 0; i < len(g.Processors); i++ {
		procType, _, _, paramters := g.Processors[i].Processor.Definition()

		attrs := []xml.Attr{
			xml.Attr{Name: xml.Name{Local: "type"}, Value: procType},
			xml.Attr{Name: xml.Name{Local: "name"}, Value: g.Processors[i].Name},
			xml.Attr{Name: xml.Name{Local: "x"}, Value: strconv.Itoa(g.Processors[i].X)},
			xml.Attr{Name: xml.Name{Local: "y"}, Value: strconv.Itoa(g.Processors[i].Y)},
		}

		for i := 0; i < len(paramters); i++ {
			attrs = append(attrs, xml.Attr{Name: xml.Name{Local: paramters[i].Name}, Value: strconv.FormatFloat(float64(paramters[i].Value), 'f', -1, 32)})
		}

		e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "processor"}, Attr: attrs})
		e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "processor"}})
	}
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "processors"}})

	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "connectors"}})
	for i := 0; i < len(g.Connectors); i++ {
		fromProcessorDefiniton := g.definitonForProcessor(g.Connectors[i].FromProcessor)
		_, _, fromPorts, _ := fromProcessorDefiniton.Processor.Definition()
		toProcessorDefiniton := g.definitonForProcessor(g.Connectors[i].ToProcessor)
		_, toPorts, _, _ := toProcessorDefiniton.Processor.Definition()

		attrs := []xml.Attr{
			xml.Attr{Name: xml.Name{Local: "fromProcessor"}, Value: fromProcessorDefiniton.GetName()},
			xml.Attr{Name: xml.Name{Local: "fromPort"}, Value: fromPorts[g.Connectors[i].FromPort]},
			xml.Attr{Name: xml.Name{Local: "toProcessor"}, Value: toProcessorDefiniton.GetName()},
			xml.Attr{Name: xml.Name{Local: "toPort"}, Value: toPorts[g.Connectors[i].ToPort]},
		}

		e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "connector"}, Attr: attrs})
		e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "connector"}})
	}
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "connectors"}})

	return e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "graph"}})
}

// attrToMap is a helper funtion to produce a map of xml attributes
func attrToMap(attr []xml.Attr) map[string]string {
	result := make(map[string]string)
	for i := 0; i < len(attr); i++ {
		result[attr[i].Name.Local] = attr[i].Value
	}

	return result
}

//UnmarshalXML Unmarshals the graph
func (g *Graph) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	println("unmarsall")

	for {
		token, err := d.Token()
		if err != nil {
			fmt.Println(g)

			// At End of File
			return nil
		}

		switch el := token.(type) {
		case xml.StartElement:
			attr := attrToMap(el.Attr)

			switch el.Name.Local {
			case "processor":
				x, _ := strconv.Atoi(attr["x"])
				y, _ := strconv.Atoi(attr["y"])
				proc := MakeProcessor(attr["type"])
				for name, value := range attr {
					floatValue, _ := strconv.ParseFloat(value, 32)
					i, err := processor.GetProcessorParameterIndex(proc, name)
					if err == nil {
						println(i, float32(floatValue))
						proc.SetParameter(i, float32(floatValue))
					}
				}

				procDef := ProcessorDefinition{Name: attr["name"], X: x, Y: y, Processor: proc}
				g.Processors = append(g.Processors, procDef)
			case "connector":
				fromProcessor := g.getProcessorByName(attr["fromProcessor"])
				fromPort, _ := processor.GetProcessorOutputIndex(fromProcessor, attr["fromPort"])
				toProcessor := g.getProcessorByName(attr["toProcessor"])
				toPort, _ := processor.GetProcessorInputIndex(toProcessor, attr["toPort"])

				connector := Connector{
					FromProcessor: fromProcessor, FromPort: fromPort,
					ToProcessor: toProcessor, ToPort: toPort,
				}
				g.Connectors = append(g.Connectors, connector)
			}
		}
	}
}
