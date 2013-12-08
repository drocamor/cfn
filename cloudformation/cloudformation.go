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
	Default     interface{}
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

func max(nums ...int) (n int) {
	n = nums[0]
	for _, i := range nums {
		if i > n {
			n = i
		}
	}
	return n
}

func (t *CloudFormationTemplate) GraphvizAttrs(name string) (ret map[string]string) {

	maxRows := max(len(t.Parameters), len(t.Resources), len(t.Outputs))

	// Parameters

	params := "{ Parameters"
	for p, _ := range t.Parameters {
		params = fmt.Sprintf("%s | <P_%s> %s", params, p, p)
	}

	for i := 0; i < maxRows-len(t.Parameters); i++ {
		params = fmt.Sprintf("%s | ", params)
	}

	params = fmt.Sprintf("%s }", params)

	// Resources

	resources := "{ Resources"
	for r, _ := range t.Resources {
		resources = fmt.Sprintf("%s | %s", resources, r)
	}

	for i := 0; i < maxRows-len(t.Resources); i++ {
		resources = fmt.Sprintf("%s | ", resources)
	}

	resources = fmt.Sprintf("%s }", resources)

	// Outputs

	outputs := "{ Outputs"
	for o, _ := range t.Outputs {
		outputs = fmt.Sprintf("%s | <O_%s> %s", outputs, o, o)
	}

	for i := 0; i < maxRows-len(t.Outputs); i++ {
		outputs = fmt.Sprintf("%s | ", outputs)
	}

	outputs = fmt.Sprintf("%s }", outputs)

	ret = map[string]string{
		"shape":    "record",
		"fontname": "Courier",
		"label": fmt.Sprintf("\" %s |{ %s | %s | %s}\"", name, params, resources, outputs),
	}

	return ret
}
