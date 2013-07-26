package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/henrikhodne/tci/travis"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "tci"
	app.Usage = "A command-line tool for interacting with Travis CI"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{"repo", "", "The repository to interact with. Example: green-eggs/ham"},
	}

	app.Commands = []cli.Command{
		{
			Name:      "show",
			ShortName: "s",
			Usage:     "Displays generic information about a build",
			Action: func(c *cli.Context) {
				slug := c.GlobalString("repo")
				client := travis.NewClient()

				repo, err := client.GetRepository(slug)
				if err != nil {
					fmt.Printf("Unable to get repository: %v\n", err)
					return
				}

				build, err := client.GetBuild(repo.LastBuildID)
				if err != nil {
					fmt.Printf("Unable to get build: %v\n", err)
					return
				}

				fmt.Printf("Build #%s:\n", build.Number)
				fmt.Printf("State:\t\t%s\n", build.State)
			},
		},
	}

	app.Run(os.Args)
}
