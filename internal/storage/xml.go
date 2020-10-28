package storage

import (
	"fmt"
	"github.com/clbanning/mxj"
)

// XMLParser is a Parser implementation to handle yaml files.
type XMLParser struct {
}

// FromBytes returns some Data that is represented by the given bytes.
func (p *XMLParser) FromBytes(byteData []byte) (interface{}, error) {
	data, err := mxj.NewMapXml(byteData)
	if err != nil {
		return data, fmt.Errorf("could not unmarshal config data: %w", err)
	}
	return map[string]interface{}(data), nil
}

// ToBytes returns a slice of bytes that represents the given value.
func (p *XMLParser) ToBytes(value interface{}) ([]byte, error) {
	m, ok := value.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("cannot marshal type `%T` to xml", value)
	}
	mv := mxj.New()
	for k, v := range m {
		mv[k] = v
	}
	return mv.XmlIndent("", "  ")
}
