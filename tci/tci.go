package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/henrikhodne/tci"
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
				repo, _ := client.GetRepository(slug)
				build, _ := client.GetBuild(repo.LastBuildID)
				fmt.Printf("Build #%s: %s\n", build.Number, build.CommitSubject)
				fmt.Printf("State:\t\t%s\n", build.State)
			},
		},
	}

	app.Run(os.Args)
}
