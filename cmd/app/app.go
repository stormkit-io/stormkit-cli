package app

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stormkit-io/stormkit-cli/cmd"
)

// appCmd represents the app command
var appCmd = &cobra.Command{
	Use:   "app",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("app called")
	},
}

func init() {
	cmd.GetRootCmd().AddCommand(appCmd)
}
