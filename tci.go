package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/henrikhodne/tci/travis"
	"os"
	"strings"
	"time"
)

func color(color int, body string) string {
	return fmt.Sprintf("\033[%dm%s\033[0m", color, body)
}

func bold(body string) string {
	return color(1, body)
}

func pad(str string, length int) string {
	return fmt.Sprintf("%-*s", length, str)
}

func printInfo(name, info string) {
	fmt.Printf("%s%s\n", bold(color(33, pad(name, 15))), info)
}

func formatDuration(seconds int) string {
	return (time.Duration(seconds) * time.Second).String()
}

func formatTime(timeStr string) string {
	t, _ := time.Parse(time.RFC3339, timeStr)
	loc, _ := time.LoadLocation("Local")

	return t.In(loc).Format("2006-01-02 15:04:05")
}

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

				repoResp, _ := client.GetRepository(slug)
				if repoResp == (travis.RepositoryResponse{}) {
					println("Couldn't find repository.")
					return
				}

				buildResp, _ := client.GetBuild(repoResp.Repository.LastBuildID)
				if buildResp == (travis.BuildResponse{}) {
					println("Couldn't find build.")
					return
				}

				build := buildResp.Build
				commit := buildResp.Commit

				fmt.Printf(bold("Build #%s: %s\n"), build.Number, strings.Split(commit.Message, "\n")[0])
				printInfo("State", build.State)
				if build.PullRequest {
					printInfo("Type", "pull request")
				} else {
					printInfo("Type", "push")
				}
				printInfo("Branch", commit.Branch)
				printInfo("Compare URL", commit.CompareURL)
				printInfo("Duration", formatDuration(build.Duration))
				printInfo("Started", formatTime(build.StartedAt))
				printInfo("Finished", formatTime(build.FinishedAt))
			},
		},
	}

	app.Run(os.Args)
}
