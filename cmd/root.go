package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "iget",
	Short: "CLI to download images from URL.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage: iget [command] [flags]\n\nFor more information, use help.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(getCmd)
}
