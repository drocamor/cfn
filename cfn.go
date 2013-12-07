package main

import (

	"os"
	"github.com/codegangsta/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "cfn"
	app.Usage = "tools for CloudFormation"

	app.Commands = []cli.Command{
		{
			Name:      "dot",
			ShortName: "d",
			Usage:     "generates Graphviz dot files from templates",
			Action: dot,
		},
	}

	app.Run(os.Args)

}
