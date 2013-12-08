package cloudformation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type CloudFormationOutput struct {
	Description string
}

type CloudFormationParameter struct {
	Default     string
	Description string
	Type        string
}

type CloudFormationResource struct {
	Type string
}

type CloudFormationTemplate struct {
	AWSTemplateFormatVersion string
	Description              string
	Mappings                 map[string]interface{}
	Outputs                  map[string]CloudFormationOutput
	Parameters               map[string]CloudFormationParameter
	Resources                map[string]CloudFormationResource
}

func LoadTemplate(path string) (t CloudFormationTemplate) {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	json.Unmarshal(file, &t)

	return t
}

func (t *CloudFormationTemplate) GraphvizAttrs(name string) (ret map[string]string) {

	params := "{ Parameters"
	for p, _ := range t.Parameters {
		params = fmt.Sprintf("%s | %s", params, p)
	}
	params = fmt.Sprintf("%s }", params)

	resources := "{ Resources"
	for p, _ := range t.Resources {
		resources = fmt.Sprintf("%s | %s", resources, p)
	}
	resources = fmt.Sprintf("%s }", resources)

	outputs := "{ Outputs"
	for p, _ := range t.Outputs {
		outputs = fmt.Sprintf("%s | %s", outputs, p)
	}
	outputs = fmt.Sprintf("%s }", outputs)

	ret = map[string]string{
		"shape":    "record",
		"fontname": "Courier",
		"label":    fmt.Sprintf("\"{ %s | { %s | %s | %s}}\"", name, params, resources, outputs),
	}

	return ret
}
