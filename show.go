package main

import (
	"fmt"
	"github.com/henrikhodne/tci/travis"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

func init() {
	cmds["show"] = cmd{show, "", "display information about the latest build"}
	cmdHelp["show"] = `Shows a summary of the latest build.
`
}

func show() {
	slug := detectSlug()
	client := travis.NewClient()

	repoResp, _ := client.GetRepository(slug)
	if repoResp == (travis.RepositoryResponse{}) {
		println("Couldn't find repository")
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
}

func detectSlug() string {
	gitHead, _ := exec.Command("git", "name-rev", "--name-only", "HEAD").Output()
	gitRemote, _ := exec.Command("git", "config", "--get", "branch."+strings.TrimSpace(string(gitHead))+".remote").Output()
	gitInfo, _ := exec.Command("git", "config", "--get", "remote."+strings.TrimSpace(string(gitRemote))+".url").Output()
	url := strings.TrimSpace(string(gitInfo))
	re := regexp.MustCompile(`^(?:https://|git://|git@)github\.com[:/](.*/.+?)(\.git)?$`)
	return re.FindStringSubmatch(url)[1]
}

func bold(str string) string {
	return color(1, str)
}

func color(color int, str string) string {
	return fmt.Sprintf("\033[%dm%s\033[0m", color, str)
}

func printInfo(name, info string) {
	fmt.Printf("%s%s\n", bold(color(33, pad(name, 15))), info)
}

func pad(str string, length int) string {
	return fmt.Sprintf("%-*s", length, str)
}

func formatDuration(seconds int) string {
	return (time.Duration(seconds) * time.Second).String()
}

func formatTime(timeStr string) string {
	t, _ := time.Parse(time.RFC3339, timeStr)
	loc, _ := time.LoadLocation("Local")
	return t.In(loc).Format("2006-01-02 15:04:05")
}
