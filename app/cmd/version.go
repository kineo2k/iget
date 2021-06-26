package cmd

import (
	"fmt"
	"iget/constants"
)
import "github.com/spf13/cobra"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Output the version number.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("%s version %s", constants.AppName, constants.Version))
	},
}
