package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var gitHash, gitVersionTag, buildDate string

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Display package version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("httpst")
			fmt.Printf("version: %s\n", gitVersionTag)
			fmt.Printf("hash: %s\n", gitHash)
			fmt.Printf("build date: %s\n", buildDate)
		},
	}
}
