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
