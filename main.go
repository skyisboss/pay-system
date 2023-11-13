package main

import "github.com/skyisboss/pay-system/cmd"

var (
	gitCommit  string
	gitVersion string
	// embedFrontend string
)

func main() {
	cmd.Version = gitVersion
	cmd.Commit = gitCommit
	// cmd.EmbedFrontend = lo.Must(strconv.ParseBool(embedFrontend))
	cmd.Execute()
}
