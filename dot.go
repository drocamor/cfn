package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/drocamor/cfn/cloudformation"
)

func dot(c *cli.Context) {

	fmt.Println("Making dot from: ", c.Args()[0])
	template := cloudformation.LoadTemplate(c.Args()[0])
	fmt.Printf("Template description is: %s\n", template.Description)
}
