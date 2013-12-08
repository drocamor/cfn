package main

import (
	"code.google.com/p/gographviz"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/drocamor/cfn/cloudformation"
"strings"
)

func dot(c *cli.Context) {

	g := gographviz.NewGraph()
	g.SetName("G")
	g.SetDir(true)
	for _, file := range c.Args() {
		name := strings.Split(file, ".")[0]
		template := cloudformation.LoadTemplate(file)
	
		g.AddNode("G", name, template.GraphvizAttrs())
	}
	

	s := g.String()
	fmt.Println(s)
}
