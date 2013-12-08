package main

import (
	"code.google.com/p/gographviz"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/drocamor/cfn/cloudformation"
	"strings"
	"path"
)

func dot(c *cli.Context) {

	g := gographviz.NewGraph()
	g.SetName("G")
	g.SetDir(true)
	g.AddAttr("G", "rankdir", "LR")
	for _, file := range c.Args() {
		filename := path.Base(file)
		name := strings.Split(filename, ".")[0]
		template := cloudformation.LoadTemplate(file)

		g.AddNode("G", name, template.GraphvizAttrs(name))
	}

	s := g.String()
	fmt.Println(s)
}
